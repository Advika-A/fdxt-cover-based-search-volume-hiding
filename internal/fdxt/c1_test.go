package fdxt

import (
	"fdxt-poc/internal/crypto"
	"reflect"
	"testing"
)

func TestC1TargetVolumeAdaptive(t *testing.T) {
	tests := []struct {
		name     string
		policy   C1Policy
		input    int
		expected int
	}{
		// Power-of-2
		{"power2_0", C1PolicyPower2, 0, 1},
		{"power2_1", C1PolicyPower2, 1, 1},
		{"power2_2", C1PolicyPower2, 2, 2},
		{"power2_3", C1PolicyPower2, 3, 4},
		{"power2_5", C1PolicyPower2, 5, 8},
		{"power2_9", C1PolicyPower2, 9, 16},
		{"power2_16", C1PolicyPower2, 16, 16},

		// Square-root
		{"sqrt_0", C1PolicySqrt, 0, 0},
		{"sqrt_1", C1PolicySqrt, 1, 2},
		{"sqrt_4", C1PolicySqrt, 4, 6},
		{"sqrt_10", C1PolicySqrt, 10, 14},
		{"sqrt_16", C1PolicySqrt, 16, 20},

		// Logarithmic
		{"log_0", C1PolicyLog, 0, 0},
		{"log_1", C1PolicyLog, 1, 2},
		{"log_4", C1PolicyLog, 4, 7},
		{"log_10", C1PolicyLog, 10, 14},
		{"log_16", C1PolicyLog, 16, 21},

		// Hybrid
		{"hybrid_0", C1PolicyHybrid, 0, 0},
		{"hybrid_1", C1PolicyHybrid, 1, 3},
		{"hybrid_4", C1PolicyHybrid, 4, 9},
		{"hybrid_10", C1PolicyHybrid, 10, 18},
		{"hybrid_16", C1PolicyHybrid, 16, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := C1Config{
				MaxCover: 100,
				Policy:   tt.policy,
			}

			got := C1TargetVolume(tt.input, cfg)

			if got != tt.expected {
				t.Errorf(
					"policy=%s, input=%d: got %d, want %d",
					tt.policy,
					tt.input,
					got,
					tt.expected,
				)
			}
		})
	}

	// Verify the MaxCover budget still works.
	if got := C1TargetVolume(
		100,
		C1Config{
			MaxCover: 4,
			Policy:   C1PolicyPower2,
		},
	); got != 104 {
		t.Fatalf("budget cap failed: got %d want 104", got)
	}
}

func TestC1N2UsesProtocolShapedTKLAndPreservesLogicalResult(t *testing.T) {
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)
	for _, d := range toyDataset {
		for _, kw := range d.keywords {
			client.Update(server, OpAdd, kw, d.id)
		}
	}
	if err := client.SeedC1Pool(server, []string{"K1", "K2", "K3"}, 8); err != nil {
		t.Fatal(err)
	}
	base := client.PrepareSearchBaseline([]string{"K1", "K2"})
	c1, err := client.PrepareSearchC1(
		[]string{"K1", "K2"},
		C1Config{
			MaxCover: 8,
			Policy:   C1PolicyPower2,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if c1.C1 == nil {
		t.Fatal("C1 request missing")
	}
	if len(c1.TKL) <= len(base.TKL) {
		t.Fatalf("expected C1 to append TKL entries: base=%d c1=%d", len(base.TKL), len(c1.TKL))
	}
	if len(c1.Rows) <= len(base.Rows) {
		t.Fatalf("expected C1 to append rows: base=%d c1=%d", len(base.Rows), len(c1.Rows))
	}
	if c1.C1.CoverCount <= 0 {
		t.Fatal("expected positive C1 cover count")
	}
	res := server.ProcessQuery(c1)
	got := sorted(client.RecoverSearch([]string{"K1", "K2"}, res))
	want := []string{"ID1", "ID4"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("C1 result=%v want %v", got, want)
	}
	if len(res) <= len(base.Rows) {
		t.Fatalf("C1 did not create protocol search outputs: response=%d baseline=%d", len(res), len(base.Rows))
	}
	if len(res)-len(base.Rows) != len(c1.Rows)-len(base.Rows) {
		t.Fatalf("cover response count mismatch")
	}
}

func TestC1CoverRowsRecoverAndFilter(t *testing.T) {
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)
	for _, d := range toyDataset {
		for _, kw := range d.keywords {
			client.Update(server, OpAdd, kw, d.id)
		}
	}
	if err := client.SeedC1Pool(server, []string{"K1", "K2", "K3"}, 8); err != nil {
		t.Fatal(err)
	}
	parts, err := client.PrepareSearchC1(
		[]string{"K1", "K2"},
		C1Config{
			MaxCover: 8,
			Policy:   C1PolicyPower2,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	res := server.ProcessQuery(parts)
	if len(res) != len(parts.Rows) {
		t.Fatalf(
			"response must correspond to all protocol rows: got=%d rows=%d",
			len(res),
			len(parts.Rows),
		)
	}
	got := sorted(client.RecoverSearch([]string{"K1", "K2"}, res))
	if !reflect.DeepEqual(got, []string{"ID1", "ID4"}) {
		t.Fatalf("recovered=%v", got)
	}
}

func TestC1ServerReceivesOnlyOrdinarySearchMaterial(t *testing.T) {
	client, server := newLoadedDemo(t)
	if err := client.SeedC1Pool(server, []string{"K1", "K2", "K3"}, 8); err != nil {
		t.Fatal(err)
	}
	parts, err := client.PrepareSearchC1(
		[]string{"K1", "K2"},
		C1Config{
			MaxCover: 8,
			Policy:   C1PolicyPower2,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if parts.C1 == nil || parts.C1.CoverCount == 0 {
		t.Fatal("expected C1 metadata on client")
	}
	// The server entry point ignores C1 metadata and executes the same search
	// path used by baseline FDXT. No CDBX/CDBT cover state is supplied per query.
	res1 := server.ProcessQuery(parts)
	res2 := server.ProcessSearch(parts.TKL, parts.Rows, parts.RestKeywords)
	if len(res1) != len(res2) {
		t.Fatalf(
			"C1 dispatch changed response processing: %d vs %d",
			len(res1),
			len(res2),
		)
	}
}

func TestC1NotAppliedToN3(t *testing.T) {
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)
	for _, d := range toyDataset {
		for _, kw := range d.keywords {
			client.Update(server, OpAdd, kw, d.id)
		}
	}
	parts := client.PrepareSearch([]string{"K1", "K2", "K3"})
	if parts.C1 != nil {
		t.Fatal("C1 must not be attached to n=3")
	}
}
