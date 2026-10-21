package demodata

import (
	"reflect"
	"strings"
	"testing"
)

func TestDatasetShape(t *testing.T) {
	if got := DocumentCount(); got != 50 {
		t.Fatalf("DocumentCount() = %d, want 50", got)
	}

	docs := GetAllDocuments()
	if len(docs) != 50 {
		t.Fatalf("GetAllDocuments() returned %d documents, want 50", len(docs))
	}
	seen := make(map[string]bool, len(docs))
	vocab := make(map[string]bool)
	for i, d := range docs {
		wantID := "ID" + itoa(i+1)
		if d.ID != wantID {
			t.Fatalf("document %d has ID %q, want %q", i, d.ID, wantID)
		}
		if seen[d.ID] {
			t.Fatalf("duplicate document ID %q", d.ID)
		}
		seen[d.ID] = true
		if len(d.Keywords) == 0 {
			t.Fatalf("%s has no keywords", d.ID)
		}
		for _, kw := range d.Keywords {
			vocab[kw] = true
		}
	}
	if len(vocab) < 20 || len(vocab) > 30 {
		t.Fatalf("distinct keywords = %d, want 20-30", len(vocab))
	}
}

func TestKeywordFrequencyTiers(t *testing.T) {
	freq := KeywordFrequencies()
	high, medium, low := 0, 0, 0
	for kw, n := range freq {
		switch {
		case n >= 20:
			high++
		case n >= 8:
			medium++
		case n >= 3:
			low++
		default:
			t.Fatalf("keyword %q occurs only %d times; minimum is 3", kw, n)
		}
	}
	if high == 0 || medium == 0 || low == 0 {
		t.Fatalf("frequency tiers missing: high=%d medium=%d low=%d", high, medium, low)
	}
}

func TestRecordsDeterministicAndUnique(t *testing.T) {
	a := Records()
	b := Records()
	if !reflect.DeepEqual(a, b) {
		t.Fatal("Records() is not deterministic")
	}
	seen := make(map[string]bool, len(a))
	for _, r := range a {
		if seen[r.Keyword] {
			t.Fatalf("duplicate keyword record %q", r.Keyword)
		}
		seen[r.Keyword] = true
		if len(r.DocIDs) == 0 {
			t.Fatalf("keyword %q has no document IDs", r.Keyword)
		}
	}
}

func TestDocumentLookup(t *testing.T) {
	if got := GetKeywords("ID1"); len(got) == 0 {
		t.Fatal("GetKeywords(ID1) returned empty")
	}
	if _, ok := GetDocumentByID("ID50"); !ok {
		t.Fatal("GetDocumentByID(ID50) returned false")
	}
	if _, ok := GetDocumentByID("ID51"); ok {
		t.Fatal("GetDocumentByID(ID51) returned true")
	}
}

func TestGroundTruthQueries(t *testing.T) {
	checks := map[string]int{
		"security AND network":                  10,
		"access AND data":                       12,
		"computer AND software":                 10,
		"malware AND phishing":                  1,
		"cloud AND database":                    5,
		"malware AND vulnerability":             1,
		"authentication AND backup":             1,
		"forensics AND vulnerability":           1,
		"firewall AND api":                      0,
		"cloud AND application":                 0,
		"computer AND security AND network":     2,
		"computer AND software AND application": 6,
		"cloud AND data AND database":           5,
		"forensics AND malware AND threat":      4,
		"firewall AND api AND web":              0,
	}
	for q, want := range checks {
		got := queryGroundTruth(q)
		if len(got) != want {
			t.Fatalf("%s: count=%d, want %d; IDs=%v", q, len(got), want, got)
		}
	}
}

func TestGroundTruthMissingKeyword(t *testing.T) {
	if got := GroundTruth([]string{"computer", "keyword_that_does_not_exist"}); len(got) != 0 {
		t.Fatalf("missing keyword conjunction = %v, want empty", got)
	}
}

func queryGroundTruth(q string) []string {
	var keywords []string
	for _, part := range splitAND(q) {
		keywords = append(keywords, part)
	}
	return GroundTruth(keywords)
}

func splitAND(q string) []string {
	var out []string
	for _, f := range strings.Fields(q) {
		if f != "AND" {
			out = append(out, f)
		}
	}
	return out
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
