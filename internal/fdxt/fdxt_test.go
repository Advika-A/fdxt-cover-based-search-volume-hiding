package fdxt

import (
	"reflect"
	"sort"
	"testing"

	"fdxt-poc/internal/crypto"
)

// toyDataset is the exact dataset requested for the POC:
//
//	ID1 -> [K1, K2]
//	ID2 -> [K1]
//	ID3 -> [K2]
//	ID4 -> [K1, K2]
//	ID5 -> [K3]
var toyDataset = []struct {
	id       string
	keywords []string
}{
	{"ID1", []string{"K1", "K2"}},
	{"ID2", []string{"K1"}},
	{"ID3", []string{"K2"}},
	{"ID4", []string{"K1", "K2"}},
	{"ID5", []string{"K3"}},
}

func newLoadedDemo(t *testing.T) (*Client, *Server) {
	t.Helper()
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)

	for _, d := range toyDataset {
		for _, kw := range d.keywords {
			client.Update(server, OpAdd, kw, d.id)
		}
	}
	return client, server
}

func sorted(s []string) []string {
	out := append([]string(nil), s...)
	sort.Strings(out)
	return out
}

// TestCrossTagAlgebra checks, independently of the Search protocol, that
// the core FDXT relationship holds:
//
//	xtk_{j,k}^alpha == xtag(wk, id, op)
//
// where xtk_{j,k} = g^(Fp(KX,wk) * Fp(KZ,w1||j)) and
//
//	alpha = Fp(KY,id||op) * Fp(KZ,w1||j)^-1
//
// is exactly the blinding factor produced by Update for w1's j-th entry.
func TestCrossTagAlgebra(t *testing.T) {
	p := crypto.Init()
	c := NewClient(p)

	w1, wk, id, op := "K1", "K2", "ID1", OpAdd

	// Simulate one Update(w1, id) manually to get j=1's alpha.
	ctr := c.counter(w1)
	ctr.Updt++
	ctr.Max++
	maxB := u64b(ctr.Max)

	fy := crypto.Fp(p, c.KY, []byte(id+"|"+string(op)))
	fz := crypto.Fp(p, c.KZ, []byte(w1), maxB)
	fzInv := p.Pairing.NewZr().Invert(fz)
	alpha := p.Pairing.NewZr().Mul(fy, fzInv)

	// Expected cross-tag for (wk, id, op):
	fxWk := crypto.Fp(p, c.KX, []byte(wk))
	expExp := p.Pairing.NewZr().Mul(fxWk, fy)
	expected := p.Pairing.NewG1().PowZn(p.Generator, expExp)

	// Candidate cross-token xtk_{1,k} and its blinded form:
	xtk := p.Pairing.NewG1().PowZn(p.Generator, p.Pairing.NewZr().Mul(fxWk, fz))
	got := p.Pairing.NewG1().PowZn(xtk, alpha)

	if string(got.Bytes()) != string(expected.Bytes()) {
		t.Fatalf("cross-tag algebra mismatch:\n got = %x\nwant = %x", got.Bytes(), expected.Bytes())
	}
}

func TestSearchSingleKeyword_N1(t *testing.T) {
	client, server := newLoadedDemo(t)

	got := sorted(client.Search(server, []string{"K1"}))
	want := []string{"ID1", "ID2", "ID4"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(K1) = %v, want %v", got, want)
	}
}

func TestSearchSingleKeyword_K3(t *testing.T) {
	client, server := newLoadedDemo(t)

	got := sorted(client.Search(server, []string{"K3"}))
	want := []string{"ID5"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(K3) = %v, want %v", got, want)
	}
}

// TestSearchConjunctive_N2 is the most important test in this POC:
// K1 AND K2 must return exactly {ID1, ID4} -- the actual ID set, not just
// a count.
func TestSearchConjunctive_N2(t *testing.T) {
	client, server := newLoadedDemo(t)

	got := sorted(client.Search(server, []string{"K1", "K2"}))
	want := []string{"ID1", "ID4"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(K1 AND K2) = %v, want %v", got, want)
	}
}

// TestSearchConjunctive_QueryOrderIndependent checks that swapping the
// query keyword order does not change the (still w1-selection-driven)
// result.
func TestSearchConjunctive_QueryOrderIndependent(t *testing.T) {
	client, server := newLoadedDemo(t)

	got := sorted(client.Search(server, []string{"K2", "K1"}))
	want := []string{"ID1", "ID4"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(K2 AND K1) = %v, want %v", got, want)
	}
}

// TestSearchConjunctive_NoMatch checks a conjunctive query with an empty
// intersection.
func TestSearchConjunctive_NoMatch(t *testing.T) {
	client, server := newLoadedDemo(t)

	got := client.Search(server, []string{"K1", "K3"})
	if len(got) != 0 {
		t.Fatalf("Search(K1 AND K3) = %v, want empty", got)
	}
}

// TestSearchConjunctive_AfterPriorSingleKeywordSearch exercises FDXT's
// deferred cross-tag reveal: K2's cross-tags are only decrypted into XSET
// once K2 is actually searched. Searching K1 alone first, and only later
// issuing K1 AND K2, must still produce the correct intersection.
func TestSearchConjunctive_AfterPriorSingleKeywordSearch(t *testing.T) {
	client, server := newLoadedDemo(t)

	_ = client.Search(server, []string{"K1"}) // K1's cross-tags revealed, K2's are not yet

	got := sorted(client.Search(server, []string{"K1", "K2"}))
	want := []string{"ID1", "ID4"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(K1 AND K2) after prior Search(K1) = %v, want %v", got, want)
	}
}

// TestSearchConjunctive_AfterFurtherUpdates checks that a new Add issued
// after a conjunctive search is correctly picked up by a subsequent search
// (forward privacy should not break correctness).
func TestSearchConjunctive_AfterFurtherUpdates(t *testing.T) {
	client, server := newLoadedDemo(t)

	got1 := sorted(client.Search(server, []string{"K1", "K2"}))
	if !reflect.DeepEqual(got1, []string{"ID1", "ID4"}) {
		t.Fatalf("initial Search(K1 AND K2) = %v", got1)
	}

	client.Update(server, OpAdd, "K1", "ID3") // ID3 already has K2
	client.Update(server, OpAdd, "K2", "ID3")

	got2 := sorted(client.Search(server, []string{"K1", "K2"}))
	want2 := []string{"ID1", "ID3", "ID4"}
	if !reflect.DeepEqual(got2, want2) {
		t.Fatalf("Search(K1 AND K2) after updates = %v, want %v", got2, want2)
	}
}

func TestSearchConjunctive_N3(t *testing.T) {
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)
	dataset := map[string][]string{
		"ID1": {"K1", "K2", "K3"},
		"ID2": {"K1", "K2"},
		"ID3": {"K1", "K3"},
		"ID4": {"K2", "K3"},
	}
	for id, kws := range dataset {
		for _, kw := range kws {
			client.Update(server, OpAdd, kw, id)
		}
	}
	got := sorted(client.Search(server, []string{"K1", "K2", "K3"}))
	want := []string{"ID1"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(K1 AND K2 AND K3) = %v, want %v", got, want)
	}
}
