package db

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDecodeRecord_RealSchema(t *testing.T) {
	raw := map[string]interface{}{
		"_id":     primitive.NewObjectID(),
		"k":       "F0",
		"val_set": []interface{}{"0", "1", "2", "3"},
	}
	rec, ok := decodeRecord(raw)
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if rec.Keyword != "F0" {
		t.Fatalf("Keyword = %q, want F0", rec.Keyword)
	}
	want := []string{"0", "1", "2", "3"}
	if !reflect.DeepEqual(rec.DocIDs, want) {
		t.Fatalf("DocIDs = %v, want %v", rec.DocIDs, want)
	}
}

func TestDecodeRecord_PrimitiveA(t *testing.T) {
	// This is what the real mongo-driver actually produces for a BSON
	// array decoded into a generic map[string]interface{} value.
	raw := map[string]interface{}{
		"k":       "F1227",
		"val_set": primitive.A{"D000123", "D004821"},
	}
	rec, ok := decodeRecord(raw)
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if rec.Keyword != "F1227" {
		t.Fatalf("Keyword = %q, want F1227", rec.Keyword)
	}
	want := []string{"D000123", "D004821"}
	if !reflect.DeepEqual(rec.DocIDs, want) {
		t.Fatalf("DocIDs = %v, want %v", rec.DocIDs, want)
	}
}

func TestDecodeRecord_MongoObjectIDIsIgnoredNotUsedAsIdentifier(t *testing.T) {
	// "_id" must never be treated as the keyword or as a document ID --
	// only "k" and "val_set" carry meaning under the real schema.
	raw := map[string]interface{}{
		"_id":     primitive.NewObjectID(),
		"k":       "F621",
		"val_set": []interface{}{"7"},
	}
	rec, ok := decodeRecord(raw)
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if rec.Keyword != "F621" {
		t.Fatalf("Keyword = %q, want F621 (must come from \"k\", never \"_id\")", rec.Keyword)
	}
}

func TestDecodeRecord_MissingKeyword(t *testing.T) {
	raw := map[string]interface{}{
		"val_set": []interface{}{"0", "1"},
	}
	_, ok := decodeRecord(raw)
	if ok {
		t.Fatalf("expected ok=false when \"k\" is missing")
	}
}

func TestDecodeRecord_EmptyKeywordString(t *testing.T) {
	raw := map[string]interface{}{
		"k":       "",
		"val_set": []interface{}{"0"},
	}
	_, ok := decodeRecord(raw)
	if ok {
		t.Fatalf("expected ok=false when \"k\" is an empty string")
	}
}

func TestDecodeRecord_MissingDocIDs(t *testing.T) {
	raw := map[string]interface{}{
		"k": "F0",
	}
	_, ok := decodeRecord(raw)
	if ok {
		t.Fatalf("expected ok=false when \"val_set\" is missing")
	}
}

func TestDecodeRecord_EmptyDocIDsArray(t *testing.T) {
	raw := map[string]interface{}{
		"k":       "F0",
		"val_set": []interface{}{},
	}
	_, ok := decodeRecord(raw)
	if ok {
		t.Fatalf("expected ok=false when \"val_set\" is an empty array")
	}
}

// TestDecodeRecord_OldSchemaFieldsAreNotUsed guards against regressing
// back to the incorrect Session 2 schema ({"id": ..., "val_st": [...]}):
// a document that only has the OLD field names must NOT decode.
func TestDecodeRecord_OldSchemaFieldsAreNotUsed(t *testing.T) {
	raw := map[string]interface{}{
		"id":     "D0001",
		"val_st": []interface{}{"F1227", "F1438"},
	}
	_, ok := decodeRecord(raw)
	if ok {
		t.Fatalf("expected ok=false: \"id\"/\"val_st\" are the old, incorrect field names " +
			"and must not be accepted by the real-schema decoder")
	}
}
