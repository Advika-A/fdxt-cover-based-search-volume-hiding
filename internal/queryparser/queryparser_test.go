package queryparser

import (
	"reflect"
	"testing"
)

func TestParse_TwoKeywordUpperAND(t *testing.T) {
	got, err := Parse("F1227 AND F1438")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"F1227", "F1438"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %v, want %v", got, want)
	}
}

func TestParse_CaseInsensitiveAND(t *testing.T) {
	got, err := Parse("F621 and F774")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"F621", "F774"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %v, want %v", got, want)
	}
}

func TestParse_ExtraWhitespace(t *testing.T) {
	got, err := Parse("  F2953   AND    F774  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"F2953", "F774"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %v, want %v", got, want)
	}
}

func TestParse_SingleKeyword(t *testing.T) {
	got, err := Parse("K1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"K1"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %v, want %v", got, want)
	}
}

func TestParse_EmptyQueryErrors(t *testing.T) {
	if _, err := Parse(""); err == nil {
		t.Fatalf("expected error for empty query")
	}
	if _, err := Parse("   "); err == nil {
		t.Fatalf("expected error for whitespace-only query")
	}
}

func TestParse_OnlyANDErrors(t *testing.T) {
	// A bare "AND" with nothing on either side splits into two empty
	// parts, both of which are dropped -- this is correctly a malformed
	// query, not a literal one-word keyword named "AND".
	if _, err := Parse("AND"); err == nil {
		t.Fatalf("expected error for a query that is only the AND separator")
	}
}

func TestParse_TrailingANDLeavesNoSecondKeyword(t *testing.T) {
	// "F1 AND " has a real "AND" token (word-boundary delimited); the
	// empty tail after it is dropped, leaving just the one real keyword
	// rather than erroring.
	got, err := Parse("F1 AND ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"F1"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %v, want %v", got, want)
	}
}

func TestParse_DoesNotSplitANDInsideAWord(t *testing.T) {
	// "BRAND" must not be split into "BR" + "D" -- the separator requires
	// surrounding whitespace.
	got, err := Parse("BRAND")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"BRAND"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %v, want %v", got, want)
	}
}
