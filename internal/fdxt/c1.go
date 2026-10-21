package fdxt

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"fdxt-poc/internal/crypto"
	"github.com/Nik-U/pbc"
)

// C1 adds protocol-shaped cover material to the client-generated search.
// Crucially, the server is never told which entries are cover entries and no
// C1 state is sent over the search RPC. Cover records are provisioned in the
// encrypted server index during setup using exactly the same CDB_X/CDB_T
// formats as ordinary records. The client alone keeps the mapping from cover
// identifiers to those encrypted records.
type C1Policy string

const (
	C1PolicyPower2 C1Policy = "power2"
	C1PolicySqrt   C1Policy = "sqrt"
	C1PolicyLog    C1Policy = "log"
	C1PolicyHybrid C1Policy = "hybrid"
)

type C1Config struct {
	MaxCover int
	Policy   C1Policy
}

var DefaultC1Config = C1Config{
	MaxCover: 32,
	Policy:   C1PolicyPower2,
}


// C1PoolKeyword is the client-only metadata needed to reconstruct a
// protocol-shaped TKL token and anchor row for one cover document.
type C1PoolKeyword struct {
	J uint64
	L []byte
	T []byte
}

// C1PoolEntry represents one synthetic encrypted document. The same opaque
// document identifier is provisioned under every keyword so a selected cover
// row can legitimately satisfy a two-keyword conjunction.
type C1PoolEntry struct {
	ID       string
	Keywords map[string]C1PoolKeyword
}

// C1Request is intentionally client-local. It carries only experiment
// metadata; it is NOT serialized into the gRPC request.
type C1Request struct {
	Config     C1Config
	CoverCount int
}

func validateC1Config(cfg C1Config) error {
	if cfg.MaxCover < 0 {
		return fmt.Errorf("c1: MaxCover must be >= 0")
	}
	return nil
}

func nextPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	p := 1
	for p < n {
		if p > int(^uint(0)>>1)/2 {
			return n
		}
		p <<= 1
	}
	return p
}

func C1TargetVolume(candidateCount int, cfg C1Config) int {
	if candidateCount < 0 {
		candidateCount = 0
	}

	if cfg.MaxCover <= 0 {
		return candidateCount
	}

	var target int

	switch cfg.Policy {
	case C1PolicySqrt:
		extra := int(math.Ceil(math.Sqrt(float64(candidateCount))))
		target = candidateCount + extra

	case C1PolicyLog:
		extra := int(math.Ceil(math.Log2(float64(candidateCount + 1))))
		target = candidateCount + extra

	case C1PolicyHybrid:
		sqrtExtra := math.Ceil(math.Sqrt(float64(candidateCount)))
		logExtra := math.Ceil(math.Log2(float64(candidateCount + 1)))
		target = candidateCount + int(sqrtExtra+logExtra)

	case C1PolicyPower2:
		fallthrough
	default:
		target = nextPowerOfTwo(candidateCount)
	}

	// Respect the configured maximum cover budget.
	if target-candidateCount > cfg.MaxCover {
		target = candidateCount + cfg.MaxCover
	}

	return target
}

func randomCoverID() (string, error) {
	b := make([]byte, 16)
	if _, err := crand.Read(b); err != nil {
		return "", err
	}
	// No special C1 prefix is used. The plaintext identifier must look like
	// an ordinary opaque document identifier to anyone without client state.
	return hex.EncodeToString(b), nil
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := crand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SeedC1Pool provisions opaque cover records in the same encrypted server
// structures used by ordinary FDXT updates. The server receives no marker
// identifying these records as cover. Client-side pool metadata is persisted
// separately in client state.
// BuildC1Pool creates client-only cover metadata and the corresponding
// encrypted CDB_X/CDB_T entries without inserting them into the server yet.
// Callers that construct a fresh index can then interleave InstallC1Entry calls
// with genuine updates so the server sees one undifferentiated update stream.
func (c *Client) BuildC1Pool(keywords []string, poolSize int) ([]C1PoolEntry, error) {
	if poolSize <= 0 {
		return nil, fmt.Errorf("c1: poolSize must be > 0")
	}
	const coverJBase = uint64(1) << 63
	pool := make([]C1PoolEntry, 0, poolSize)
	for i := 0; i < poolSize; i++ {
		id, err := randomCoverID()
		if err != nil {
			return nil, err
		}
		idOp := []byte(id + "|" + string(OpAdd))
		entry := C1PoolEntry{ID: id, Keywords: make(map[string]C1PoolKeyword, len(keywords))}
		for _, w := range keywords {
			if w == "" {
				continue
			}
			wb := []byte(w)
			j := coverJBase + uint64(i) + 1
			nonce, err := randomBytes(16)
			if err != nil {
				return nil, err
			}
			l := crypto.F(c.KT, []byte("C1-TKL-v2"), wb, idOp, nonce, []byte{0})
			t := crypto.F(c.KT, []byte("C1-TKL-v2"), wb, idOp, nonce, []byte{1})
			entry.Keywords[w] = C1PoolKeyword{J: j, L: l, T: t}
		}
		pool = append(pool, entry)
	}
	return pool, nil
}

// InstallC1Entry installs one already-generated cover record in the server's
// encrypted state. It uses the same cryptographic construction as an ordinary
// FDXT update and carries no cover marker.
func (c *Client) InstallC1Entry(srv *Server, ce C1PoolEntry) error {
	if srv == nil {
		return fmt.Errorf("c1: nil server")
	}
	idOp := []byte(ce.ID + "|" + string(OpAdd))
	fy := crypto.Fp(c.P, c.KY, idOp)
	for w, pk := range ce.Keywords {
		wb := []byte(w)
		jB := u64b(pk.J)
		addr := crypto.F(c.KW, wb, jB, []byte{0})
		mask := crypto.F(c.KW, wb, jB, []byte{1})
		val := crypto.Xor(padTo(idOp, len(mask)), mask)
		fz := crypto.Fp(c.P, c.KZ, wb, jB)
		alpha := c.P.Pairing.NewZr().Mul(fy, c.P.Pairing.NewZr().Invert(fz))
		srv.CDBT[string(addr)] = CDBTEntry{Val: append([]byte(nil), val...), Alpha: alpha}
		// pk.L/pk.T already encode the per-entry C1 TKL domain separation.
		fx := crypto.Fp(c.P, c.KX, wb)
		exp := c.P.Pairing.NewZr().Mul(fx, fy)
		xtag := c.P.Pairing.NewG1().PowZn(c.P.Generator, exp)
		cv := crypto.Xor(crypto.H(xtag.Bytes()), pk.T)
		srv.CDBX[string(pk.L)] = append([]byte(nil), cv...)
	}
	return nil
}

// InstallC1Pool installs an already-built pool in an arbitrary order supplied
// by the caller. It does not alter FDXT counters.
func (c *Client) InstallC1Pool(srv *Server, pool []C1PoolEntry) error {
	if srv == nil {
		return fmt.Errorf("c1: nil server")
	}
	c.C1Pool = append([]C1PoolEntry(nil), pool...)
	for _, ce := range pool {
		if err := c.InstallC1Entry(srv, ce); err != nil {
			return err
		}
	}
	return nil
}

// SeedC1Pool is the simple fresh-index helper used by the local POC. It
// creates and installs the pool. For a deployment that cares about update
// ordering, callers should instead BuildC1Pool and interleave InstallC1Entry
// with genuine Update calls.
func (c *Client) SeedC1Pool(srv *Server, keywords []string, poolSize int) error {
	pool, err := c.BuildC1Pool(keywords, poolSize)
	if err != nil {
		return err
	}
	c.C1Pool = pool
	for _, ce := range pool {
		if err := c.InstallC1Entry(srv, ce); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) c1PoolEntry(i int) *C1PoolEntry {
	if i < 0 || i >= len(c.C1Pool) {
		return nil
	}
	return &c.C1Pool[i]
}

func randomPoolIndices(n, k int) ([]int, error) {
	if k > n {
		return nil, fmt.Errorf("c1: pool has %d entries, need %d", n, k)
	}
	seen := make(map[int]bool, k)
	out := make([]int, 0, k)
	for len(out) < k {
		r, err := crand.Int(crand.Reader, big.NewInt(int64(n)))
		if err != nil {
			return nil, err
		}
		i := int(r.Int64())
		if seen[i] {
			continue
		}
		seen[i] = true
		out = append(out, i)
	}
	return out, nil
}

// ensureC1PoolForQuery lazily provisions enough opaque cover records for a
// local in-process Search call. Remote deployments provision the pool during
// setup; this helper only exists so the in-process API remains correct when a
// caller constructs a fresh Client+Server and immediately searches.
func (c *Client) ensureC1PoolForQuery(srv *Server, q []string) error {
	if len(q) != 2 {
		return nil
	}
	const poolBatch = 64

	eligible := 0
	for _, ce := range c.C1Pool {
		if _, ok1 := ce.Keywords[q[0]]; !ok1 {
			continue
		}
		if _, ok2 := ce.Keywords[q[1]]; !ok2 {
			continue
		}
		eligible++
	}
	// Default MaxCover is 32, so 32 eligible entries are enough for any
	// adaptive padding decision under the default budget. Keep a 64-entry
	// reserve so repeated queries do not immediately need another setup.
	if eligible >= 32 {
		return nil
	}

	pool, err := c.BuildC1Pool(q, poolBatch)
	if err != nil {
		return err
	}
	for _, ce := range pool {
		if err := c.InstallC1Entry(srv, ce); err != nil {
			return err
		}
	}
	c.C1Pool = append(c.C1Pool, pool...)
	return nil
}

// prepareC1 chooses opaque pre-provisioned cover records. It never sends
// CDB_X/CDB_T state to the server; the server already holds those records as
// ordinary encrypted index entries.
func (c *Client) prepareC1(q []string, cfg C1Config, realRows int) (C1Request, []TKLEntry, []XTKRow, map[string]bool, error) {
	if len(q) != 2 {
		return C1Request{}, nil, nil, nil, nil
	}
	if err := validateC1Config(cfg); err != nil {
		return C1Request{}, nil, nil, nil, err
	}
	need := C1TargetVolume(realRows, cfg) - realRows
	if need <= 0 {
		return C1Request{Config: cfg}, nil, nil, map[string]bool{}, nil
	}
	w1Idx := 0
	min := c.counter(q[0]).Max
	for i, w := range q {
		if m := c.counter(w).Max; m < min {
			min = m
			w1Idx = i
		}
	}
	w1 := q[w1Idx]
	rest := make([]string, 0, 1)
	for i, w := range q {
		if i != w1Idx {
			rest = append(rest, w)
		}
	}
	w1b := []byte(w1)
	fx := make(map[string]*pbc.Element, 2)
	for _, w := range q {
		fx[w] = crypto.Fp(c.P, c.KX, []byte(w))
	}

	// A client state may contain pools provisioned for several keyword pairs.
	// Only entries that contain both query keywords are eligible; choosing from
	// the entire pool would make a mixed pool unsafe.
	eligible := make([]int, 0, len(c.C1Pool))
	for i, ce := range c.C1Pool {
		if _, ok := ce.Keywords[w1]; !ok {
			continue
		}
		if _, ok := ce.Keywords[rest[0]]; !ok {
			continue
		}
		eligible = append(eligible, i)
	}
	if len(eligible) < need {
		return C1Request{}, nil, nil, nil, fmt.Errorf("c1: cover pool has %d eligible entries for query, need %d; rerun setup with a larger pool", len(eligible), need)
	}
	chosen, err := randomPoolIndices(len(eligible), need)
	if err != nil {
		return C1Request{}, nil, nil, nil, err
	}

	coverRows := make([]XTKRow, 0, need)
	coverTKL := make([]TKLEntry, 0, need*len(rest))
	coverIDs := make(map[string]bool, need)

	for _, pos := range chosen {
		idx := eligible[pos]
		ce := c.c1PoolEntry(idx)
		if ce == nil {
			return C1Request{}, nil, nil, nil, fmt.Errorf("c1: invalid pool index %d", idx)
		}
		wk := ce.Keywords[w1]
		fz := crypto.Fp(c.P, c.KZ, w1b, u64b(wk.J))
		row := XTKRow{Addr: crypto.F(c.KW, w1b, u64b(wk.J), []byte{0}), XTK: make(map[string]*pbc.Element, len(rest))}
		for _, rkw := range rest {
			exp := c.P.Pairing.NewZr().Mul(fx[rkw], fz)
			row.XTK[rkw] = c.P.Pairing.NewG1().PowZn(c.P.Generator, exp)
			rt := ce.Keywords[rkw]
			coverTKL = append(coverTKL, TKLEntry{L: append([]byte(nil), rt.L...), T: append([]byte(nil), rt.T...)})
		}
		coverRows = append(coverRows, row)
		coverIDs[ce.ID] = true
	}

	return C1Request{Config: cfg, CoverCount: need}, coverTKL, coverRows, coverIDs, nil
}
