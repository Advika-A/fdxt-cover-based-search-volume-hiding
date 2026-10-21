package experiment

import (
	"testing"

	"fdxt-poc/internal/demodata"
)

func TestSelectWorkloadsCoversRequiredCategories(t *testing.T) {
	ws, err := selectWorkloads()
	if err != nil {
		t.Fatal(err)
	}
	seen := make(map[string]bool)
	for _, w := range ws {
		seen[w.Name] = true
	}
	for _, name := range []string{"high_high", "high_medium", "medium_medium", "medium_low", "low_low", "zero_result", "n1_regression", "n3_regression"} {
		if !seen[name] {
			t.Fatalf("missing workload %q", name)
		}
	}
}

func TestSummaryStatistics(t *testing.T) {
	rs := []Record{
		{Variant: "corrected_fdxt", Workload: "x", TotalTimeNS: 1_000_000, CommunicationBytes: 100, StorageBytes: 200, ObservedVolume: 4, CoverVolume: 0, ResidualExposure: 1, ExactMatch: true},
		{Variant: "corrected_fdxt", Workload: "x", TotalTimeNS: 3_000_000, CommunicationBytes: 200, StorageBytes: 300, ObservedVolume: 8, CoverVolume: 2, ResidualExposure: .75, ExactMatch: true},
	}
	ss := Summarize(rs)
	if len(ss) != 1 {
		t.Fatalf("summaries=%d want 1", len(ss))
	}
	if ss[0].MeanLatencyMS != 2 || ss[0].MedianLatencyMS != 2 || !ss[0].AllCorrect {
		t.Fatalf("unexpected summary: %#v", ss[0])
	}
}

func TestSelectAllPairs(t *testing.T) {
	ws, err := selectAllPairs()
	if err != nil {
		t.Fatal(err)
	}
	kw := demodata.Keywords()
	want := len(kw) * (len(kw) - 1) / 2
	if len(ws) != want {
		t.Fatalf("all-pairs workload count = %d, want %d", len(ws), want)
	}
	seen := make(map[string]bool, len(ws))
	for _, w := range ws {
		if w.N != 2 {
			t.Fatalf("workload %s has n=%d, want 2", w.Name, w.N)
		}
		if seen[w.Query] {
			t.Fatalf("duplicate pair query: %s", w.Query)
		}
		seen[w.Query] = true
	}
}
