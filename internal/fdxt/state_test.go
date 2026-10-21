package fdxt

import (
	"path/filepath"
	"reflect"
	"testing"

	"fdxt-poc/internal/crypto"
)

// TestServerStateBinaryRoundTrip protects the remote-process persistence
// boundary. FDXT CDB/XSET keys are arbitrary bytes and must survive JSON
// persistence byte-for-byte; raw JSON strings would corrupt invalid UTF-8.
func TestServerStateBinaryRoundTrip(t *testing.T) {
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)

	client.Update(server, OpAdd, "K1", "ID1")
	client.Update(server, OpAdd, "K2", "ID1")

	// Add explicit binary keys as well as the real FDXT state.
	server.CDBX[string([]byte{0xff, 0x00, 0xfe})] = []byte{0x80, 0x01, 0xff}
	server.CDBT[string([]byte{0xc3, 0x28})] = CDBTEntry{
		Val:   []byte{0xff, 0x00, 0x01},
		Alpha: p.Pairing.NewZr().SetFromHash([]byte("state-test-alpha")),
	}
	server.XSET[string([]byte{0xff, 0x00, 0xfe, 0x80})] = true

	path := filepath.Join(t.TempDir(), "server.json")
	if err := server.SaveState(path); err != nil {
		t.Fatal(err)
	}
	loaded, err := LoadServerState(path, p)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(server.CDBX, loaded.CDBX) {
		t.Fatalf("CDBX changed across JSON round-trip")
	}
	if !reflect.DeepEqual(server.XSET, loaded.XSET) {
		t.Fatalf("XSET changed across JSON round-trip")
	}
	if len(loaded.CDBT) != len(server.CDBT) {
		t.Fatalf("CDBT entry count changed: got %d want %d", len(loaded.CDBT), len(server.CDBT))
	}
	for k, want := range server.CDBT {
		got, ok := loaded.CDBT[k]
		if !ok || string(got.Val) != string(want.Val) || string(got.Alpha.Bytes()) != string(want.Alpha.Bytes()) {
			t.Fatalf("CDBT entry %x changed across JSON round-trip", []byte(k))
		}
	}
}

// TestPersistedRemotePathRoundTrip verifies that the same client/server state
// used by the separate-process commands still produces the exact conjunctive
// result after persistence and reload.
func TestPersistedRemotePathRoundTrip(t *testing.T) {
	p := crypto.Init()
	client := NewClient(p)
	server := NewServer(p)
	client.Update(server, OpAdd, "K1", "ID1")
	client.Update(server, OpAdd, "K1", "ID2")
	client.Update(server, OpAdd, "K2", "ID1")
	client.Update(server, OpAdd, "K2", "ID3")
	if err := client.SeedC1Pool(server, []string{"K1", "K2"}, 4); err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	cp := filepath.Join(dir, "client.json")
	sp := filepath.Join(dir, "server.json")
	if err := client.SaveState(cp); err != nil {
		t.Fatal(err)
	}
	if err := server.SaveState(sp); err != nil {
		t.Fatal(err)
	}

	client2, err := LoadClientState(cp, p)
	if err != nil {
		t.Fatal(err)
	}
	server2, err := LoadServerState(sp, p)
	if err != nil {
		t.Fatal(err)
	}

	if len(client2.C1Pool) != len(client.C1Pool) {
		t.Fatalf("C1 pool length changed: got %d want %d", len(client2.C1Pool), len(client.C1Pool))
	}
	got := sorted(client2.Search(server2, []string{"K1", "K2"}))
	want := []string{"ID1"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("persisted Search(K1 AND K2) = %v, want %v", got, want)
	}
}
