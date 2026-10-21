// Package db is the MongoDB integration layer for the FDXT POC. It is
// intentionally independent of the FDXT cryptographic core in
// internal/fdxt and internal/crypto: this package only knows how to read
// the real Crime_USENIX_REV dataset and compute plaintext ground truth
// directly against MongoDB. internal/indexer is the only place that
// wires this package's records into the (unchanged) FDXT Update protocol.
//
// CONFIRMED REAL SCHEMA (inspected directly against the live BSON in
// Crime_USENIX_REV.id_keywords -- this is not a guess):
//
//	{
//	  "_id": ObjectId("5ec19009d41c80e2c3ee16f7"),
//	  "k": "F0",
//	  "val_set": ["0", "1", "2", "3", ...]
//	}
//
// id_keywords is a KEYWORD -> DOCUMENT-ID-SET index: one document per
// keyword, not one document per record. "k" is the keyword; "val_set" is
// the full list of document IDs that keyword appears in. MongoDB's own
// "_id" is NOT the FDXT document identifier and is never used as one --
// it is opaque ObjectId housekeeping for this collection and is ignored
// entirely by this package.
package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// keywordField and docIDsField are the real id_keywords BSON field names.
// They are the single source of truth for those names; nothing else in
// this package hardcodes the literal strings "k" or "val_set".
const (
	keywordField = "k"
	docIDsField  = "val_set"
)

// Record is one (keyword, document-ID-set) pair read from id_keywords:
// exactly one keyword ("k") and every document ID it appears in
// ("val_set"). This is the inverse of a per-document {id, keywords[]}
// record -- there is exactly one Record per keyword in the dataset, not
// one per document.
type Record struct {
	Keyword string
	DocIDs  []string
}

// Config describes how to reach the dataset. id_keywords is the ONLY
// collection this schema needs: it already maps keyword -> document-ID
// set directly, so there is no separate reverse-lookup collection.
type Config struct {
	URI        string
	Database   string
	Collection string
}

// decodeRecord extracts a Record from a raw, loosely-typed Mongo document
// (decoded as a generic map so this works regardless of whether the
// driver hands back bson.M, primitive.M, or plain map[string]interface{}
// for a given query path). It is pure/deterministic and does not touch
// the network, which is what makes it unit-testable without a live
// MongoDB instance.
//
// ok is false if the document is missing a non-empty "k" string field or
// a usable non-empty "val_set" array, so callers can skip malformed
// documents -- or, in ValidateSchema's case, fail loudly with a clear
// schema-mismatch error -- instead of silently fabricating a record.
func decodeRecord(raw map[string]interface{}) (Record, bool) {
	kw, ok := extractKeyword(raw)
	if !ok {
		return Record{}, false
	}
	ids, ok := extractDocIDs(raw[docIDsField])
	if !ok {
		return Record{}, false
	}
	return Record{Keyword: kw, DocIDs: ids}, true
}

// extractKeyword reads the required "k" field. Unlike the document
// identifier in the old (incorrect) schema, there is deliberately no
// fallback to "_id" here: "_id" is MongoDB housekeeping, not the keyword,
// per the confirmed real schema.
func extractKeyword(raw map[string]interface{}) (string, bool) {
	v, ok := raw[keywordField]
	if !ok || v == nil {
		return "", false
	}
	s := fmt.Sprintf("%v", v)
	if s == "" {
		return "", false
	}
	return s, true
}

// extractDocIDs accepts either a plain []interface{} (used directly by
// the unit tests) or a primitive.A / bson.A (what the real mongo-driver
// produces when it decodes a BSON array into a generic map[string]any
// field, e.g. via bson.M) -- primitive.A is a distinct named type over
// []interface{}, so both cases must be handled explicitly in the type
// switch.
func extractDocIDs(v interface{}) ([]string, bool) {
	elems, ok := toSlice(v)
	if !ok || len(elems) == 0 {
		return nil, false
	}
	out := make([]string, 0, len(elems))
	for _, e := range elems {
		if e == nil {
			continue
		}
		out = append(out, fmt.Sprintf("%v", e))
	}
	if len(out) == 0 {
		return nil, false
	}
	return out, true
}

func toSlice(v interface{}) ([]interface{}, bool) {
	switch arr := v.(type) {
	case []interface{}:
		return arr, true
	case primitive.A:
		return []interface{}(arr), true
	default:
		return nil, false
	}
}
