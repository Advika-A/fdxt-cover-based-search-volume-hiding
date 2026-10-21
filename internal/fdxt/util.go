package fdxt

import "encoding/binary"

// u64b big-endian encodes a counter value so it can be fed into the PRF F
// as one of the concatenated "||" inputs (e.g. w || c_max_w || 0).
func u64b(x uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, x)
	return b
}

// padTo right-pads (or truncates) b to exactly n bytes with zero bytes.
// Used so that val = (id, op) can be XORed with a fixed-length F-mask.
// This is a POC simplification: it assumes plaintext (id||"|"||op) never
// itself ends in a zero byte, which holds for the ASCII demo identifiers
// used here. A production implementation would length-prefix instead.
func padTo(b []byte, n int) []byte {
	if len(b) >= n {
		return b[:n]
	}
	out := make([]byte, n)
	copy(out, b)
	return out
}

// trimZeros strips trailing zero bytes added by padTo.
func trimZeros(b []byte) []byte {
	i := len(b)
	for i > 0 && b[i-1] == 0 {
		i--
	}
	return b[:i]
}
