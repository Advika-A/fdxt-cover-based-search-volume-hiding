package tests

import (
	"reflect"
	"sort"
	"testing"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/queryparser"
)

func TestDataset50_FDXTExactMatchesGroundTruth(t *testing.T) {
	p := crypto.Init()
	client := fdxt.NewClient(p)
	server := fdxt.NewServer(p)

	for _, doc := range demodata.GetAllDocuments() {
		for _, kw := range doc.Keywords {
			client.Update(server, fdxt.OpAdd, kw, doc.ID)
		}
	}

	// Exercise the three important query sizes first, then the full
	// deterministic query suite. Ground truth is calculated directly from
	// demodata and never through FDXT.
	queries := []string{
		"security",
		"security AND network",
		"computer AND security AND network",
	}
	queries = append(queries, demodata.QuerySet...)

	for _, raw := range queries {
		keywords, err := queryparser.Parse(raw)
		if err != nil {
			t.Fatalf("Parse(%q): %v", raw, err)
		}
		got := client.Search(server, keywords)
		want := demodata.GroundTruthConjunctiveQuery(keywords)
		res := compare.Compare(got, want)
		if !reflect.DeepEqual(res.FDXTIDs, res.GroundTruthIDs) {
			t.Fatalf("%s: FDXT=%v ground-truth=%v FP=%v FN=%v", raw, res.FDXTIDs, res.GroundTruthIDs, res.FalsePositives, res.FalseNegatives)
		}
		if len(got) > 0 && !res.ExactMatch {
			t.Fatalf("%s: expected exact match", raw)
		}
	}
}

func TestDataset50_QuerySetIsDeterministic(t *testing.T) {
	if len(demodata.QuerySet) != 15 {
		t.Fatalf("QuerySet length = %d, want 15", len(demodata.QuerySet))
	}
	for _, q := range demodata.QuerySet {
		keywords, err := queryparser.Parse(q)
		if err != nil {
			t.Fatalf("Parse(%q): %v", q, err)
		}
		if len(keywords) != 2 && len(keywords) != 3 {
			t.Fatalf("%q has n=%d; query suite should contain n=2/n=3 queries only", q, len(keywords))
		}
	}
}

func TestDataset50_KeywordLookupsAreDeterministic(t *testing.T) {
	first := demodata.GetDocumentsForKeyword("computer")
	second := demodata.GetDocumentsForKeyword("computer")
	sort.Strings(first)
	sort.Strings(second)
	if !reflect.DeepEqual(first, second) {
		t.Fatalf("computer lookup changed between calls: %v vs %v", first, second)
	}
	if len(first) == 0 {
		t.Fatal("computer should have a non-empty document set")
	}
}
