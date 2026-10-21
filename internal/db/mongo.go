package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store is a thin wrapper around the real FDXT MongoDB dataset: the
// id_keywords collection, where each document IS one keyword's full
// document-ID set ({"k": "<keyword>", "val_set": ["<id1>", "<id2>", ...]}).
// id_keywords already is a keyword -> document-ID index, so there is no
// separate reverse-lookup collection to manage.
type Store struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	cfg        Config
}

// Connect opens a MongoDB connection and verifies it with Ping before
// returning, so callers get a clear connection error up front rather than
// a confusing failure on the first query. It does not itself validate
// that cfg.Collection exists, is non-empty, or matches the expected
// schema -- call ValidateSchema for that, before doing any real work.
func Connect(ctx context.Context, cfg Config) (*Store, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, fmt.Errorf("db: connect to %q: %w", cfg.URI, err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("db: ping %q: %w", cfg.URI, err)
	}
	database := client.Database(cfg.Database)
	coll := database.Collection(cfg.Collection)

	return &Store{client: client, database: database, collection: coll, cfg: cfg}, nil
}

// ValidateSchema checks that the configured id_keywords collection (see
// Config.Collection) actually exists in the database, contains at least
// one document, and that a sampled document actually matches the real
// {"k": <keyword>, "val_set": [<id>, ...]} shape -- returning a
// descriptive error if any of that fails.
//
// This is a deliberate pre-flight check: the caller (the CLI) uses it to
// refuse to run -- and never to report a trivial/empty "Exact match:
// true", or a silently-empty index from a schema mismatch -- when pointed
// at an empty, misconfigured, or wrongly-shaped database, rather than
// indexing zero (keyword, id) pairs and comparing two empty result sets.
func (s *Store) ValidateSchema(ctx context.Context) error {
	names, err := s.database.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("db: list collections on %q: %w", s.cfg.Database, err)
	}
	found := false
	for _, n := range names {
		if n == s.cfg.Collection {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf(
			"db: collection %q does not exist in database %q (found collections: %v)",
			s.cfg.Collection, s.cfg.Database, names)
	}

	count, err := s.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("db: count documents in %q: %w", s.CollectionPath(), err)
	}
	if count == 0 {
		return fmt.Errorf(
			"db: collection %q exists but contains 0 documents -- refusing to run "+
				"(an empty dataset would make every FDXT-vs-ground-truth comparison "+
				"trivially/vacuously \"match\", which proves nothing)", s.CollectionPath())
	}

	var sample map[string]interface{}
	if err := s.collection.FindOne(ctx, bson.M{}).Decode(&sample); err != nil {
		return fmt.Errorf("db: sample document from %q: %w", s.CollectionPath(), err)
	}
	if _, ok := decodeRecord(sample); !ok {
		return fmt.Errorf(
			"db: schema mismatch on %q: sampled document does not have a non-empty %q "+
				"string field and a non-empty %q array field (the real FDXT id_keywords "+
				"schema is {\"k\": \"<keyword>\", \"val_set\": [\"<id1>\", \"<id2>\", ...]}); "+
				"sampled document: %v", s.CollectionPath(), keywordField, docIDsField, sample)
	}

	return nil
}

// Seed upserts the given records into the collection, one MongoDB
// document per record, in the exact real id_keywords shape
// ({"k": rec.Keyword, "val_set": rec.DocIDs}) -- keyed on the keyword
// field so re-seeding is idempotent (re-running it replaces each
// keyword's document-ID set rather than duplicating it).
//
// This is used only to initialize the small internal/demodata dataset
// into its own dedicated database (see demodata.DatabaseName) for the
// CLI's demo mode; it is never pointed at the real Crime_USENIX_REV
// dataset, which is always populated externally and only ever read by
// this package, never written.
func (s *Store) Seed(ctx context.Context, records []Record) (int64, error) {
	var n int64
	for _, rec := range records {
		filter := bson.M{keywordField: rec.Keyword}
		doc := bson.M{keywordField: rec.Keyword, docIDsField: rec.DocIDs}
		if _, err := s.collection.ReplaceOne(ctx, filter, doc, options.Replace().SetUpsert(true)); err != nil {
			return n, fmt.Errorf("db: seed upsert for keyword %q into %q: %w", rec.Keyword, s.CollectionPath(), err)
		}
		n++
	}
	return n, nil
}

// Close disconnects the underlying MongoDB client.
func (s *Store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

// CollectionPath returns "<database>.<collection>" for logging/CLI output.
func (s *Store) CollectionPath() string {
	return s.cfg.Database + "." + s.cfg.Collection
}

// EstimatedDocumentCount returns a fast, approximate count of documents in
// the collection (used only for CLI progress/diagnostic output). Note
// that under the real schema each "document" here is one keyword, not one
// dataset record.
func (s *Store) EstimatedDocumentCount(ctx context.Context) (int64, error) {
	return s.collection.EstimatedDocumentCount(ctx)
}

// StreamRecords iterates every document in id_keywords (up to limit, or
// unbounded if limit <= 0), decodes it into a Record (one keyword plus
// every document ID it appears in), and invokes fn for each one. It
// returns the number of records successfully decoded and passed to fn.
// Documents that don't decode into a valid Record (missing/empty "k" or
// "val_set") are silently skipped, not counted, and do not abort the scan
// -- this mirrors how a real dataset can contain a few malformed rows
// without that being a bug in the pipeline.
//
// This is a plain, unbounded-memory streaming read (one document in
// flight at a time via the cursor) rather than loading the whole
// collection into a slice, which matters once this points at the full
// multi-thousand-row Crime_USENIX_REV collection.
func (s *Store) StreamRecords(ctx context.Context, limit int64, fn func(Record) error) (int64, error) {
	findOpts := options.Find().SetProjection(bson.M{keywordField: 1, docIDsField: 1})
	if limit > 0 {
		findOpts.SetLimit(limit)
	}

	cur, err := s.collection.Find(ctx, bson.M{}, findOpts)
	if err != nil {
		return 0, fmt.Errorf("db: find: %w", err)
	}
	defer cur.Close(ctx)

	var count int64
	for cur.Next(ctx) {
		var raw map[string]interface{}
		if err := cur.Decode(&raw); err != nil {
			return count, fmt.Errorf("db: decode document: %w", err)
		}
		rec, ok := decodeRecord(raw)
		if !ok {
			continue // malformed document: no usable "k"/"val_set"
		}
		if err := fn(rec); err != nil {
			return count, err
		}
		count++
	}
	if err := cur.Err(); err != nil {
		return count, fmt.Errorf("db: cursor: %w", err)
	}
	return count, nil
}

// KeywordGroundTruth returns the exact set of document IDs that appear
// under EVERY given keyword in id_keywords, computed directly from
// MongoDB. It never goes anywhere near the FDXT encrypted index, so it is
// a trustworthy, independent oracle to compare FDXT's returned results
// against.
//
// Because id_keywords already stores one document per keyword with that
// keyword's full document-ID set, this is a single $in lookup (one round
// trip for all the queried keywords) followed by an in-process set
// intersection -- not a collection scan.
//
// A keyword that does not appear in id_keywords at all has an empty
// document-ID set by definition, so the conjunction with it is empty; it
// is not an error.
//
// len(keywords) == 0 returns an error rather than silently matching every
// document.
func (s *Store) KeywordGroundTruth(ctx context.Context, keywords []string) ([]string, error) {
	if len(keywords) == 0 {
		return nil, fmt.Errorf("db: KeywordGroundTruth: no keywords given")
	}

	inList := make([]interface{}, len(keywords))
	for i, k := range keywords {
		inList[i] = k
	}
	filter := bson.M{keywordField: bson.M{"$in": inList}}

	cur, err := s.collection.Find(ctx, filter,
		options.Find().SetProjection(bson.M{keywordField: 1, docIDsField: 1}))
	if err != nil {
		return nil, fmt.Errorf("db: ground-truth find: %w", err)
	}
	defer cur.Close(ctx)

	found := make(map[string][]string, len(keywords))
	for cur.Next(ctx) {
		var raw map[string]interface{}
		if err := cur.Decode(&raw); err != nil {
			return nil, fmt.Errorf("db: ground-truth decode: %w", err)
		}
		rec, ok := decodeRecord(raw)
		if !ok {
			continue
		}
		found[rec.Keyword] = rec.DocIDs
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("db: ground-truth cursor: %w", err)
	}

	var intersection map[string]bool
	for _, kw := range keywords {
		ids := found[kw] // nil, i.e. empty set, if kw wasn't found at all
		if intersection == nil {
			intersection = make(map[string]bool, len(ids))
			for _, id := range ids {
				intersection[id] = true
			}
			continue
		}
		next := make(map[string]bool, len(intersection))
		for _, id := range ids {
			if intersection[id] {
				next[id] = true
			}
		}
		intersection = next
	}

	out := make([]string, 0, len(intersection))
	for id := range intersection {
		out = append(out, id)
	}
	return out, nil
}

// SampleCoOccurringKeywordPair finds two distinct keywords in id_keywords
// whose document-ID sets actually intersect, without hardcoding any
// specific keyword string from the dataset. It is used only by
// integration tests that need a genuinely co-occurring keyword pair.
//
// It samples up to sampleSize keyword documents (a bounded scan, not the
// whole collection) and checks pairs among them for a non-empty
// intersection.
func (s *Store) SampleCoOccurringKeywordPair(ctx context.Context) (string, string, error) {
	const sampleSize = 200

	findOpts := options.Find().
		SetProjection(bson.M{keywordField: 1, docIDsField: 1}).
		SetLimit(sampleSize)
	cur, err := s.collection.Find(ctx, bson.M{}, findOpts)
	if err != nil {
		return "", "", fmt.Errorf("db: sample co-occurring pair: find: %w", err)
	}
	defer cur.Close(ctx)

	var sampled []Record
	for cur.Next(ctx) {
		var raw map[string]interface{}
		if err := cur.Decode(&raw); err != nil {
			return "", "", fmt.Errorf("db: sample co-occurring pair: decode: %w", err)
		}
		rec, ok := decodeRecord(raw)
		if !ok {
			continue
		}
		sampled = append(sampled, rec)
	}
	if err := cur.Err(); err != nil {
		return "", "", fmt.Errorf("db: sample co-occurring pair: cursor: %w", err)
	}

	for i := 0; i < len(sampled); i++ {
		iSet := toStringSet(sampled[i].DocIDs)
		for j := i + 1; j < len(sampled); j++ {
			for _, id := range sampled[j].DocIDs {
				if iSet[id] {
					return sampled[i].Keyword, sampled[j].Keyword, nil
				}
			}
		}
	}
	return "", "", fmt.Errorf(
		"db: no co-occurring keyword pair found among the first %d sampled keywords", len(sampled))
}

func toStringSet(ids []string) map[string]bool {
	m := make(map[string]bool, len(ids))
	for _, id := range ids {
		m[id] = true
	}
	return m
}
