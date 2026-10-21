package compare

import "testing"

func strEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCompare_ExactMatch(t *testing.T) {
	r := Compare([]string{"ID4", "ID1"}, []string{"ID1", "ID4"})
	if !r.ExactMatch {
		t.Fatalf("expected ExactMatch=true, got Result=%+v", r)
	}
	if len(r.FalsePositives) != 0 || len(r.FalseNegatives) != 0 {
		t.Fatalf("expected no FP/FN, got FP=%v FN=%v", r.FalsePositives, r.FalseNegatives)
	}
	if !strEq(r.FDXTIDs, []string{"ID1", "ID4"}) {
		t.Fatalf("FDXTIDs not sorted correctly: %v", r.FDXTIDs)
	}
	if r.ReturnedCount != 2 || r.GroundTruth != 2 {
		t.Fatalf("counts wrong: returned=%d groundTruth=%d", r.ReturnedCount, r.GroundTruth)
	}
}

func TestCompare_FalsePositive(t *testing.T) {
	// FDXT returned an extra ID not in ground truth.
	r := Compare([]string{"ID1", "ID4", "ID9"}, []string{"ID1", "ID4"})
	if r.ExactMatch {
		t.Fatalf("expected ExactMatch=false")
	}
	if !strEq(r.FalsePositives, []string{"ID9"}) {
		t.Fatalf("FalsePositives = %v, want [ID9]", r.FalsePositives)
	}
	if len(r.FalseNegatives) != 0 {
		t.Fatalf("expected no false negatives, got %v", r.FalseNegatives)
	}
}

func TestCompare_FalseNegative(t *testing.T) {
	// FDXT missed an ID that ground truth has.
	r := Compare([]string{"ID1"}, []string{"ID1", "ID4"})
	if r.ExactMatch {
		t.Fatalf("expected ExactMatch=false")
	}
	if !strEq(r.FalseNegatives, []string{"ID4"}) {
		t.Fatalf("FalseNegatives = %v, want [ID4]", r.FalseNegatives)
	}
	if len(r.FalsePositives) != 0 {
		t.Fatalf("expected no false positives, got %v", r.FalsePositives)
	}
}

func TestCompare_BothEmpty(t *testing.T) {
	r := Compare(nil, nil)
	if !r.ExactMatch {
		t.Fatalf("expected ExactMatch=true for two empty sets")
	}
	if r.ReturnedCount != 0 || r.GroundTruth != 0 {
		t.Fatalf("expected zero counts, got returned=%d groundTruth=%d", r.ReturnedCount, r.GroundTruth)
	}
}

func TestCompare_DuplicatesInInputDoNotInflateCounts(t *testing.T) {
	// Compare should not be fooled by duplicate entries in either slice
	// when deciding ExactMatch (it treats both inputs as sets).
	r := Compare([]string{"ID1", "ID1", "ID4"}, []string{"ID1", "ID4", "ID4"})
	if !r.ExactMatch {
		t.Fatalf("expected ExactMatch=true despite duplicates, got %+v", r)
	}
}

func TestCompare_CompletelyDisjoint(t *testing.T) {
	r := Compare([]string{"ID1"}, []string{"ID2"})
	if r.ExactMatch {
		t.Fatalf("expected ExactMatch=false for disjoint sets")
	}
	if !strEq(r.FalsePositives, []string{"ID1"}) {
		t.Fatalf("FalsePositives = %v, want [ID1]", r.FalsePositives)
	}
	if !strEq(r.FalseNegatives, []string{"ID2"}) {
		t.Fatalf("FalseNegatives = %v, want [ID2]", r.FalseNegatives)
	}
}
