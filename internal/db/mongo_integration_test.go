package db

import (
	"context"
	"os"
	"testing"
	"time"
)

// integrationConfig builds a Config from environment variables and skips
// the test (does not fail it) if FDXT_MONGO_URI is not set. This is the
// standard Go pattern for "only run against infrastructure that's
// actually there" -- CI or a machine without MongoDB simply skips these,
// it does not report a false failure.
//
//	FDXT_MONGO_URI          e.g. mongodb://localhost:27017   (required to run)
//	FDXT_MONGO_DB           default "Crime_USENIX_REV"
//	FDXT_MONGO_COLLECTION   default "id_keywords"
func integrationConfig(t *testing.T) Config {
	t.Helper()
	uri := os.Getenv("FDXT_MONGO_URI")
	if uri == "" {
		t.Skip("FDXT_MONGO_URI not set; skipping MongoDB integration test. " +
			"Set FDXT_MONGO_URI (and optionally FDXT_MONGO_DB / FDXT_MONGO_COLLECTION) " +
			"to run this against a real Crime_USENIX_REV instance.")
	}
	dbName := os.Getenv("FDXT_MONGO_DB")
	if dbName == "" {
		dbName = "Crime_USENIX_REV"
	}
	collName := os.Getenv("FDXT_MONGO_COLLECTION")
	if collName == "" {
		collName = "id_keywords"
	}
	return Config{URI: uri, Database: dbName, Collection: collName}
}

func TestIntegration_ConnectAndPing(t *testing.T) {
	cfg := integrationConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	store, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect(%+v) failed: %v", cfg, err)
	}
	defer store.Close(context.Background())
}

func TestIntegration_StreamRecords_ProducesUsableRecords(t *testing.T) {
	cfg := integrationConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	store, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	const sampleLimit = 50
	var seen int
	var pairs int
	count, err := store.StreamRecords(ctx, sampleLimit, func(r Record) error {
		seen++
		if r.Keyword == "" {
			t.Errorf("record with empty Keyword: %+v", r)
		}
		if len(r.DocIDs) == 0 {
			t.Errorf("record %q has no document IDs", r.Keyword)
		}
		pairs += len(r.DocIDs)
		return nil
	})
	if err != nil {
		t.Fatalf("StreamRecords failed: %v", err)
	}
	if count == 0 {
		t.Fatalf("StreamRecords decoded 0 records from %s -- is the collection empty, "+
			"or do documents not have both a \"k\" string field and a \"val_set\" array?", cfg.Collection)
	}
	if int64(seen) != count {
		t.Fatalf("callback invoked %d times but StreamRecords reported count=%d", seen, count)
	}
	if pairs == 0 {
		t.Fatalf("StreamRecords decoded %d keyword records but 0 total (keyword,id) pairs", count)
	}
}

func TestIntegration_KeywordGroundTruth_SingleKeywordIsNonEmpty(t *testing.T) {
	cfg := integrationConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	store, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	kw1, _, err := store.SampleCoOccurringKeywordPair(ctx)
	if err != nil {
		t.Skipf("could not sample a keyword pair from the dataset (is it loaded?): %v", err)
	}

	ids, err := store.KeywordGroundTruth(ctx, []string{kw1})
	if err != nil {
		t.Fatalf("KeywordGroundTruth(%q) failed: %v", kw1, err)
	}
	if len(ids) == 0 {
		t.Fatalf("KeywordGroundTruth(%q) returned 0 documents, but the keyword was just "+
			"sampled from a real id_keywords entry -- the lookup looks broken", kw1)
	}
}

func TestIntegration_KeywordGroundTruth_ConjunctiveIsSubsetOfEach(t *testing.T) {
	cfg := integrationConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	store, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	kw1, kw2, err := store.SampleCoOccurringKeywordPair(ctx)
	if err != nil {
		t.Skipf("could not sample a co-occurring keyword pair (is the dataset loaded?): %v", err)
	}

	single1, err := store.KeywordGroundTruth(ctx, []string{kw1})
	if err != nil {
		t.Fatalf("KeywordGroundTruth(%q) failed: %v", kw1, err)
	}
	single2, err := store.KeywordGroundTruth(ctx, []string{kw2})
	if err != nil {
		t.Fatalf("KeywordGroundTruth(%q) failed: %v", kw2, err)
	}
	conj, err := store.KeywordGroundTruth(ctx, []string{kw1, kw2})
	if err != nil {
		t.Fatalf("KeywordGroundTruth(%q AND %q) failed: %v", kw1, kw2, err)
	}

	// Basic set-theoretic sanity check on the ground-truth oracle itself
	// (independent of FDXT): the conjunctive result must be a subset of
	// each individual keyword's result set, and must be non-empty since
	// kw1/kw2 were sampled specifically because their sets intersect.
	if len(conj) == 0 {
		t.Fatalf("conjunctive ground truth for %q AND %q is empty, but the pair was sampled "+
			"specifically because their document-ID sets intersect", kw1, kw2)
	}
	set1 := toStringSet(single1)
	set2 := toStringSet(single2)
	for _, id := range conj {
		if !set1[id] {
			t.Errorf("id %q in (%q AND %q) ground truth but not in %q ground truth alone", id, kw1, kw2, kw1)
		}
		if !set2[id] {
			t.Errorf("id %q in (%q AND %q) ground truth but not in %q ground truth alone", id, kw1, kw2, kw2)
		}
	}
}

func TestIntegration_ValidateSchema_RealCollectionIsUsable(t *testing.T) {
	cfg := integrationConfig(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	store, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	if err := store.ValidateSchema(ctx); err != nil {
		t.Fatalf("ValidateSchema(%+v) failed against what should be a real, populated dataset: %v", cfg, err)
	}
}

func TestIntegration_ValidateSchema_RejectsMissingCollection(t *testing.T) {
	cfg := integrationConfig(t)
	cfg.Collection = "this_collection_should_not_exist_in_the_real_dataset"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	store, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer store.Close(context.Background())

	if err := store.ValidateSchema(ctx); err == nil {
		t.Fatalf("ValidateSchema succeeded against a nonexistent collection %q -- it should have "+
			"refused, per the requirement to never treat a missing/empty dataset as usable", cfg.Collection)
	}
}
