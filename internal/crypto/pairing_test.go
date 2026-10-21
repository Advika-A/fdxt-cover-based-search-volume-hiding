package crypto

import (
	"bytes"
	"testing"
)

func TestInitReturnsSingleSharedInstance(t *testing.T) {
	c1 := Init()
	c2 := Init()
	if c1 != c2 {
		t.Fatalf("Init() must return the same *Ctx on every call (one pairing, one generator for the process lifetime)")
	}
	if c1.Pairing == nil {
		t.Fatalf("Pairing must not be nil")
	}
	if c1.Generator == nil {
		t.Fatalf("Generator must not be nil")
	}
}

func TestGeneratorIsDeterministic(t *testing.T) {
	c := Init()
	// Recompute the generator independently via the same fixed seed and
	// confirm it matches what Init() produced (deterministic hash-to-G1,
	// not a random sample).
	g2 := c.Pairing.NewG1().SetFromHash([]byte(fixedGeneratorSeed))
	if !bytes.Equal(c.Generator.Bytes(), g2.Bytes()) {
		t.Fatalf("generator is not deterministic across independent derivations")
	}
}

func TestG1SerializationRoundTrip(t *testing.T) {
	c := Init()

	e := c.Pairing.NewG1().Rand()
	b := e.Bytes()
	if len(b) == 0 {
		t.Fatalf("expected non-empty G1 serialization")
	}

	e2 := c.Pairing.NewG1().SetBytes(b)
	if !bytes.Equal(e.Bytes(), e2.Bytes()) {
		t.Fatalf("G1 element did not round-trip through Bytes()/SetBytes(): got %x want %x", e2.Bytes(), e.Bytes())
	}
}

func TestZrSerializationRoundTrip(t *testing.T) {
	c := Init()

	e := c.Pairing.NewZr().Rand()
	b := e.Bytes()
	e2 := c.Pairing.NewZr().SetBytes(b)
	if !bytes.Equal(e.Bytes(), e2.Bytes()) {
		t.Fatalf("Zr element did not round-trip through Bytes()/SetBytes()")
	}
}
