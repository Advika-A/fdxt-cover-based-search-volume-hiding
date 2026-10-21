package grpcproto

import (
	"bytes"
	"testing"
)

func TestSearchRequestWireRoundTrip(t *testing.T) {
	in := &SearchRequest{
		TKL:          []*TKLEntry{{L: []byte{1, 2}, T: []byte{3, 4}}},
		Rows:         []*XTKRow{{Addr: []byte{9}, XTK: []*XTKEntry{{Keyword: "K2", Element: []byte{5, 6}}}}},
		RestKeywords: []string{"K2", "K3"},
		C1State:      []byte{7, 8, 9, 10, 11},
		C1MaxCover:   8,
	}
	var out SearchRequest
	if err := out.Unmarshal(in.Marshal()); err != nil {
		t.Fatal(err)
	}
	if len(out.TKL) != 1 || !bytes.Equal(out.TKL[0].L, in.TKL[0].L) || !bytes.Equal(out.TKL[0].T, in.TKL[0].T) {
		t.Fatalf("TKL mismatch: %#v", out)
	}
	if len(out.Rows) != 1 || len(out.Rows[0].XTK) != 1 || out.Rows[0].XTK[0].Keyword != "K2" || !bytes.Equal(out.Rows[0].XTK[0].Element, []byte{5, 6}) {
		t.Fatalf("row mismatch: %#v", out)
	}
	if len(out.RestKeywords) != 2 || out.RestKeywords[1] != "K3" {
		t.Fatalf("keywords mismatch: %#v", out.RestKeywords)
	}
	if !bytes.Equal(out.C1State, in.C1State) || out.C1MaxCover != 8 {
		t.Fatalf("C1 fields mismatch: %#v", out)
	}
}

func TestSearchResponseWireRoundTrip(t *testing.T) {
	in := &SearchResponse{Outputs: []*SearchOutput{{Val: []byte{0, 255, 7}, Cnt: 2}, {Val: []byte{8}, Cnt: 1}}}
	var out SearchResponse
	if err := out.Unmarshal(in.Marshal()); err != nil {
		t.Fatal(err)
	}
	if len(out.Outputs) != 2 || !bytes.Equal(out.Outputs[0].Val, in.Outputs[0].Val) || out.Outputs[0].Cnt != 2 {
		t.Fatalf("response mismatch: %#v", out)
	}
}
