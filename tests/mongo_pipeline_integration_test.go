package tests

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/db"
	"fdxt-poc/internal/indexer"
)

// mongoTestConfig mirrors internal/db's integration test gating: skip
// (don't fail) unless FDXT_MONGO_URI is set. FDXT_MONGO_INDEX_LIMIT caps
// how many keyword documents get streamed into the in-memory FDXT index
// for this test, since the full Crime_USENIX_REV collection can be
// large; the CLI itself defaults to indexing the whole collection.
func mongoTestConfig(t *testing.T) (db.Config, int64) {
	t.Helper()
	uri := os.Getenv("FDXT_MONGO_URI")
	if uri == "" {
		t.Skip("FDXT_MONGO_URI not set; skipping MongoDB pipeline integration test. " +
			"Set FDXT_MONGO_URI (and optionally FDXT_MONGO_DB / FDXT_MONGO_COLLECTION / " +
			"FDXT_MONGO_INDEX_LIMIT) to run this against a real Crime_USENIX_REV instance.")
	}
	dbName := os.Getenv("FDXT_MONGO_DB")
	if dbName == "" {
		dbName = "Crime_USENIX_REV"
	}
	collName := os.Getenv("FDXT_MONGO_COLLECTION")
	if collName == "" {
		collName = "id_keywords"
	}
	limit := int64(0)
	if raw := os.Getenv("FDXT_MONGO_INDEX_LIMIT"); raw != "" {
		if v, err := strconv.ParseInt(raw, 10, 64); err == nil {
			limit = v
		}
	}
	return db.Config{URI: uri, Database: dbName, Collection: collName}, limit
}

// TestMongoBackedConjunctiveQuery_ExactMatch runs the REAL end-to-end
// pipeline: MongoDB -> FDXT encrypted index (unmodified Session 1 core)
// -> FDXT conjunctive search -> exact ID-set comparison against a live
// MongoDB ground-truth query. It does not hardcode any keyword or
// expected result: the two-keyword query is sampled from real,
// co-occurring id_keywords entries at test time, so this test is only
// ever "cheating" in the sense that both the FDXT path and the
// ground-truth path are exercised against whatever real data happens to
// be loaded.
func TestMongoBackedConjunctiveQuery_ExactMatch(t *testing.T) {
	cfg, limit := mongoTestConfig(t)

	connectCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	store, err := db.Connect(connectCtx, cfg)
	if err != nil {
		t.Fatalf("db.Connect(%+v) failed: %v", cfg, err)
	}
	defer store.Close(context.Background())

	if err := store.ValidateSchema(connectCtx); err != nil {
		t.Fatalf("ValidateSchema failed -- dataset not usable: %v", err)
	}

	kw1, kw2, err := store.SampleCoOccurringKeywordPair(connectCtx)
	if err != nil {
		t.Skipf("could not sample a co-occurring keyword pair (is the dataset loaded?): %v", err)
	}

	indexCtx, cancel2 := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel2()
	p := crypto.Init()
	client, server, stats, err := indexer.BuildFromMongo(indexCtx, store, p, limit, 0, nil)
	if err != nil {
		t.Fatalf("indexer.BuildFromMongo failed: %v", err)
	}
	if stats.Pairs == 0 {
		t.Fatalf("indexed 0 (keyword,id) pairs from %s -- collection appears empty or malformed", cfg.Collection)
	}
	t.Logf("indexed %d keywords / %d (keyword,id) pairs", stats.Keywords, stats.Pairs)

	// --- n = 1 ---
	fdxt1 := client.Search(server, []string{kw1})
	gt1, err := store.KeywordGroundTruth(connectCtx, []string{kw1})
	if err != nil {
		t.Fatalf("KeywordGroundTruth(%q) failed: %v", kw1, err)
	}
	res1 := compare.Compare(fdxt1, gt1)
	t.Logf("n=1 query %q: returned=%d groundTruth=%d exactMatch=%v", kw1, res1.ReturnedCount, res1.GroundTruth, res1.ExactMatch)
	if res1.Vacuous {
		t.Fatalf("n=1 query %q: both FDXT and ground truth are empty -- this proves nothing; "+
			"kw1 was just sampled from a real id_keywords entry so it must have documents", kw1)
	}
	if !res1.ExactMatch {
		t.Errorf("n=1 query %q: FDXT result does not exactly match MongoDB ground truth\n"+
			"  FDXT IDs:         %v\n  Ground-truth IDs: %v\n  False positives:  %v\n  False negatives:  %v",
			kw1, res1.FDXTIDs, res1.GroundTruthIDs, res1.FalsePositives, res1.FalseNegatives)
	}

	// --- n = 2 (the case this session is specifically about) ---
	fdxt2 := client.Search(server, []string{kw1, kw2})
	gt2, err := store.KeywordGroundTruth(connectCtx, []string{kw1, kw2})
	if err != nil {
		t.Fatalf("KeywordGroundTruth(%q, %q) failed: %v", kw1, kw2, err)
	}
	res2 := compare.Compare(fdxt2, gt2)
	t.Logf("n=2 query %q AND %q: returned=%d groundTruth=%d exactMatch=%v",
		kw1, kw2, res2.ReturnedCount, res2.GroundTruth, res2.ExactMatch)
	if res2.Vacuous {
		t.Fatalf("n=2 query %q AND %q: both FDXT and ground truth are empty -- this proves "+
			"nothing; kw1/kw2 were sampled specifically because their document-ID sets intersect",
			kw1, kw2)
	}
	if !res2.ExactMatch {
		t.Errorf("n=2 query %q AND %q: FDXT result does not exactly match MongoDB ground truth\n"+
			"  FDXT IDs:         %v\n  Ground-truth IDs: %v\n  False positives:  %v\n  False negatives:  %v",
			kw1, kw2, res2.FDXTIDs, res2.GroundTruthIDs, res2.FalsePositives, res2.FalseNegatives)
	}
}
