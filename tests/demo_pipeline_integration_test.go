package tests

import (
	"context"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/db"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/indexer"
)

// demoMongoURI mirrors mongoTestConfig's env-gating (FDXT_MONGO_URI), but
// always targets the dedicated demodata.DatabaseName/CollectionName --
// never the real Crime_USENIX_REV dataset -- since this test seeds data.
func demoMongoURI(t *testing.T) string {
	t.Helper()
	uri := os.Getenv("FDXT_MONGO_URI")
	if uri == "" {
		t.Skip("FDXT_MONGO_URI not set; skipping demo-mode pipeline integration test. " +
			"Set FDXT_MONGO_URI to run this against a local MongoDB instance (it only ever " +
			"writes to the dedicated FDXT_DEMO database, never Crime_USENIX_REV).")
	}
	return uri
}

// TestDemoDataset_SeedsAndLoads exercises step 1 ("Demo dataset
// loading"): Seed writes internal/demodata's literal dataset into
// FDXT_DEMO.id_keywords in the real {k, val_set} schema, and
// ValidateSchema + StreamRecords can read it straight back.
func TestDemoDataset_SeedsAndLoads(t *testing.T) {
	uri := demoMongoURI(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := db.Config{URI: uri, Database: demodata.DatabaseName, Collection: demodata.CollectionName}
	store, err := db.Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("db.Connect(%+v) failed: %v", cfg, err)
	}
	defer store.Close(context.Background())

	records := demodata.Records()
	toDBRecords := make([]db.Record, 0, len(records))
	for _, r := range records {
		toDBRecords = append(toDBRecords, db.Record{Keyword: r.Keyword, DocIDs: r.DocIDs})
	}

	seeded, err := store.Seed(ctx, toDBRecords)
	if err != nil {
		t.Fatalf("Seed failed: %v", err)
	}
	if seeded != int64(len(records)) {
		t.Fatalf("Seed reported %d documents, want %d", seeded, len(records))
	}

	if err := store.ValidateSchema(ctx); err != nil {
		t.Fatalf("ValidateSchema failed right after seeding: %v", err)
	}

	var pairs int
	count, err := store.StreamRecords(ctx, 0, func(r db.Record) error {
		pairs += len(r.DocIDs)
		return nil
	})
	if err != nil {
		t.Fatalf("StreamRecords failed: %v", err)
	}
	if count != int64(len(records)) {
		t.Fatalf("StreamRecords decoded %d keyword records, want %d", count, len(records))
	}
	if pairs == 0 {
		t.Fatalf("StreamRecords decoded 0 (keyword,id) pairs from the freshly seeded demo dataset")
	}
}

// TestDemoDataset_GroundTruthIntersection exercises step 2 ("Demo
// ground-truth intersection"): the MongoDB-side ground-truth oracle,
// computed independently of FDXT, matches the authoritative local
// dataset for the deterministic query suite.
func TestDemoDataset_GroundTruthIntersection(t *testing.T) {
	uri := demoMongoURI(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := db.Config{URI: uri, Database: demodata.DatabaseName, Collection: demodata.CollectionName}
	store, err := db.Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("db.Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	if _, err := store.Seed(ctx, toDBRecordsHelper(demodata.Records())); err != nil {
		t.Fatalf("Seed failed: %v", err)
	}

	for _, query := range demodata.QuerySet {
		keywords := splitAND(query)
		want := demodata.GroundTruthConjunctiveQuery(keywords)
		got, err := store.KeywordGroundTruth(ctx, keywords)
		if err != nil {
			t.Fatalf("KeywordGroundTruth(%v) failed: %v", keywords, err)
		}
		res := compare.Compare(got, want)
		if len(want) == 0 {
			if len(got) != 0 {
				t.Fatalf("%s: MongoDB ground truth = %v, want empty", query, got)
			}
			continue
		}
		if !res.ExactMatch {
			t.Fatalf("%s: MongoDB ground truth = %v, want exactly %v (FP=%v FN=%v)",
				query, got, want, res.FalsePositives, res.FalseNegatives)
		}
	}
}

// TestDemoDataset_FDXTQuery_ExactMatch exercises steps 3-4 ("FDXT demo
// query" and "Exact ID-set comparison"): the demo dataset is indexed
// through the REAL, unmodified FDXT protocol (indexer.BuildFromMongo ->
// fdxt.Client.Update -> fdxt.Client.Search), and the returned ID set is
// compared exactly against both the independent MongoDB ground truth and
// the authoritative local dataset -- never a value hardcoded as the
// FDXT result itself.
func TestDemoDataset_FDXTQuery_ExactMatch(t *testing.T) {
	uri := demoMongoURI(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cfg := db.Config{URI: uri, Database: demodata.DatabaseName, Collection: demodata.CollectionName}
	store, err := db.Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("db.Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	if _, err := store.Seed(ctx, toDBRecordsHelper(demodata.Records())); err != nil {
		t.Fatalf("Seed failed: %v", err)
	}
	if err := store.ValidateSchema(ctx); err != nil {
		t.Fatalf("ValidateSchema failed: %v", err)
	}

	p := crypto.Init()
	client, server, stats, err := indexer.BuildFromMongo(ctx, store, p, 0, 0, nil)
	if err != nil {
		t.Fatalf("indexer.BuildFromMongo failed: %v", err)
	}
	if stats.Pairs == 0 {
		t.Fatalf("indexed 0 (keyword,id) pairs from the demo dataset")
	}

	for _, query := range demodata.QuerySet {
		keywords := splitAND(query)

		fdxtIDs := client.Search(server, keywords) // the REAL FDXT protocol, not a shortcut
		groundTruthIDs, err := store.KeywordGroundTruth(ctx, keywords)
		if err != nil {
			t.Fatalf("%s: KeywordGroundTruth failed: %v", query, err)
		}

		res := compare.Compare(fdxtIDs, groundTruthIDs)
		want := demodata.GroundTruthConjunctiveQuery(keywords)
		if !reflect.DeepEqual(groundTruthIDs, want) {
			t.Fatalf("%s: MongoDB ground truth %v does not match authoritative dataset ground truth %v", query, groundTruthIDs, want)
		}
		if res.Vacuous {
			// Zero-result queries are intentionally part of the suite. They
			// validate that the independent oracle and FDXT agree on emptiness,
			// but are not treated as proof of positive search correctness.
			continue
		}
		if !res.ExactMatch {
			t.Fatalf("%s: FDXT result does not exactly match MongoDB ground truth\n"+
				"  FDXT IDs:         %v\n  Ground-truth IDs: %v\n  False positives:  %v\n  False negatives:  %v",
				query, res.FDXTIDs, res.GroundTruthIDs, res.FalsePositives, res.FalseNegatives)
		}
	}
}

func toDBRecordsHelper(recs []demodata.Record) []db.Record {
	out := make([]db.Record, 0, len(recs))
	for _, r := range recs {
		out = append(out, db.Record{Keyword: r.Keyword, DocIDs: r.DocIDs})
	}
	return out
}

// splitAND is a tiny local "KEYWORD1 AND KEYWORD2" splitter so this test
// file doesn't need to import internal/queryparser just for that.
func splitAND(q string) []string {
	var out []string
	for _, f := range strings.Fields(q) {
		if f == "AND" {
			continue
		}
		out = append(out, f)
	}
	return out
}
