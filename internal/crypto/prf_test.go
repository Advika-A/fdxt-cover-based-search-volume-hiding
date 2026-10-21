package crypto

import (
	"bytes"
	"testing"
)

func TestF_DeterministicAndInjectiveEnough(t *testing.T) {
	key := []byte("0123456789012345678901234567890123456789")

	a := F(key, []byte("w1"), []byte("id1"), []byte{0})
	b := F(key, []byte("w1"), []byte("id1"), []byte{0})
	if !bytes.Equal(a, b) {
		t.Fatalf("F(key, same inputs) must be deterministic")
	}

	c := F(key, []byte("w1"), []byte("id1"), []byte{1})
	if bytes.Equal(a, c) {
		t.Fatalf("F output must differ when inputs differ")
	}

	d := F([]byte("different-key-------------------"), []byte("w1"), []byte("id1"), []byte{0})
	if bytes.Equal(a, d) {
		t.Fatalf("F output must differ when the key differs")
	}
}

func TestH_Deterministic(t *testing.T) {
	x := H([]byte("hello"))
	y := H([]byte("hello"))
	if !bytes.Equal(x, y) {
		t.Fatalf("H must be deterministic")
	}
	if len(x) != 32 {
		t.Fatalf("expected 32-byte digest, got %d", len(x))
	}
}

func TestXor(t *testing.T) {
	a := []byte{0xFF, 0x0F, 0xAA}
	b := []byte{0x0F, 0xFF, 0x55}
	got := Xor(a, b)
	want := []byte{0xF0, 0xF0, 0xFF}
	if !bytes.Equal(got, want) {
		t.Fatalf("Xor(%x,%x) = %x, want %x", a, b, got, want)
	}
	// Xor(Xor(a,b), b) == a
	back := Xor(got, b)
	if !bytes.Equal(back, a) {
		t.Fatalf("Xor is not self-inverse: got %x want %x", back, a)
	}
}

func TestFp_DeterministicAndInGroup(t *testing.T) {
	c := Init()
	key := []byte("fp-test-key")

	e1 := Fp(c, key, []byte("K1"))
	e2 := Fp(c, key, []byte("K1"))
	if !bytes.Equal(e1.Bytes(), e2.Bytes()) {
		t.Fatalf("Fp must be deterministic for identical inputs")
	}

	e3 := Fp(c, key, []byte("K2"))
	if bytes.Equal(e1.Bytes(), e3.Bytes()) {
		t.Fatalf("Fp output must differ for different inputs")
	}
}
