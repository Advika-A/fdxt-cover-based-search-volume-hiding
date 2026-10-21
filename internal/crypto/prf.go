package crypto

import (
	"crypto/hmac"
	"crypto/sha256"

	"github.com/Nik-U/pbc"
)

// F is the PRF family F: {0,1}^lambda x {0,1}* -> {0,1}^lambda' from the
// paper (Sec. II-B / Algorithm 1-3). Instantiated with HMAC-SHA256, which is
// a standard, secure PRF construction. All F-calls in FDXT (addr, val-mask,
// L, T) go through this single function; the variadic parts are simply
// concatenated into the MAC input, matching the paper's "||" notation, e.g.
// F(KW, w || c_max_w || 0).
func F(key []byte, parts ...[]byte) []byte {
	mac := hmac.New(sha256.New, key)
	for _, p := range parts {
		mac.Write(p)
	}
	return mac.Sum(nil)
}

// H is the hash function H: {0,1}* -> {0,1}^lambda' used to hash the group
// element xtag (and, symmetrically, xtk^alpha) down to a fixed-length
// bitstring that can be XORed with T / compared against XSET.
func H(input []byte) []byte {
	sum := sha256.Sum256(input)
	return sum[:]
}

// Fp is the PRF family Fp: {0,1}^lambda x {0,1}* -> Zp* from the paper. It
// is instantiated by hashing an F-output into the pairing's Zr field via
// pbc's SetFromHash, i.e. Fp(K, x) := HashToZr(HMAC-SHA256(K, x)).
func Fp(ctx *Ctx, key []byte, parts ...[]byte) *pbc.Element {
	digest := F(key, parts...)
	return ctx.Pairing.NewZr().SetFromHash(digest)
}

// Xor XORs two equal-length (or truncates to the shorter of two) byte
// slices. Used for val = plaintext XOR mask and C = X XOR T.
func Xor(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = a[i] ^ b[i]
	}
	return out
}
