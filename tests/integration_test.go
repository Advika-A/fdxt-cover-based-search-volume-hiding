package tests

import (
	"reflect"
	"sort"
	"testing"

	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/fdxt"
)

// TestDemoScenario reproduces exactly the scenario requested for the POC:
// load the 5-document toy dataset, then run K1 and K1 AND K2 searches and
// check the returned ID sets (not just counts).
func TestDemoScenario(t *testing.T) {
	p := crypto.Init()
	client := fdxt.NewClient(p)
	server := fdxt.NewServer(p)

	dataset := map[string][]string{
		"ID1": {"K1", "K2"},
		"ID2": {"K1"},
		"ID3": {"K2"},
		"ID4": {"K1", "K2"},
		"ID5": {"K3"},
	}
	ids := []string{"ID1", "ID2", "ID3", "ID4", "ID5"} // deterministic order
	for _, id := range ids {
		for _, kw := range dataset[id] {
			client.Update(server, fdxt.OpAdd, kw, id)
		}
	}

	k1 := client.Search(server, []string{"K1"})
	sort.Strings(k1)
	if want := []string{"ID1", "ID2", "ID4"}; !reflect.DeepEqual(k1, want) {
		t.Fatalf("Search(K1) = %v, want %v", k1, want)
	}

	k1AndK2 := client.Search(server, []string{"K1", "K2"})
	sort.Strings(k1AndK2)
	if want := []string{"ID1", "ID4"}; !reflect.DeepEqual(k1AndK2, want) {
		t.Fatalf("Search(K1 AND K2) = %v, want %v", k1AndK2, want)
	}
}
