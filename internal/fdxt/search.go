package fdxt

import (
	"math/rand"
	"strings"

	"fdxt-poc/internal/crypto"

	"github.com/Nik-U/pbc"
)

// SearchRequestParts contains the client-generated FDXT search material that
// is sent to the server. It contains no ground-truth information.
type SearchRequestParts struct {
	TKL          []TKLEntry
	Rows         []XTKRow
	RestKeywords []string
	C1           *C1Request
}

// PrepareSearch builds the actual FDXT search tokens for q without changing
// the client's counters. It is the client-side half of Algorithm 3 and is
// suitable for transport over a remote protocol.
func (c *Client) prepareSearch(q []string, c1 bool, c1Cfg C1Config) (SearchRequestParts, error) {
	if len(q) == 0 {
		return SearchRequestParts{}, nil
	}

	w1Idx := 0
	minMax := c.counter(q[0]).Max
	for i, w := range q {
		if m := c.counter(w).Max; m < minMax {
			minMax = m
			w1Idx = i
		}
	}
	w1 := q[w1Idx]
	rest := make([]string, 0, len(q)-1)
	for i, w := range q {
		if i != w1Idx {
			rest = append(rest, w)
		}
	}

	var tkl []TKLEntry
	for _, w := range q {
		ctr := c.counter(w)
		wb := []byte(w)
		srchB := u64b(ctr.Srch)
		for l := uint64(1); l <= ctr.Updt; l++ {
			lB := u64b(l)
			L := crypto.F(c.KT, wb, srchB, lB, []byte{0})
			T := crypto.F(c.KT, wb, srchB, lB, []byte{1})
			tkl = append(tkl, TKLEntry{L: L, T: T})
		}
	}
	rand.Shuffle(len(tkl), func(i, j int) { tkl[i], tkl[j] = tkl[j], tkl[i] })

	ctr1 := c.counter(w1)
	w1b := []byte(w1)
	fxByKeyword := make(map[string]*pbc.Element, len(rest))
	for _, wk := range rest {
		fxByKeyword[wk] = crypto.Fp(c.P, c.KX, []byte(wk))
	}

	rows := make([]XTKRow, 0, ctr1.Max)
	for j := uint64(1); j <= ctr1.Max; j++ {
		jB := u64b(j)
		addr := crypto.F(c.KW, w1b, jB, []byte{0})
		row := XTKRow{Addr: addr, XTK: make(map[string]*pbc.Element, len(rest))}
		fz := crypto.Fp(c.P, c.KZ, w1b, jB)
		for _, wk := range rest {
			exp := c.P.Pairing.NewZr().Mul(fxByKeyword[wk], fz)
			xtk := c.P.Pairing.NewG1().PowZn(c.P.Generator, exp)
			row.XTK[wk] = xtk
		}
		rows = append(rows, row)
	}

	parts := SearchRequestParts{TKL: tkl, Rows: rows, RestKeywords: rest}
	if c1 && len(q) == 2 {
		c1req, coverTKL, coverRows, coverIDs, err := c.prepareC1(q, c1Cfg, len(rows))
		if err != nil {
			return SearchRequestParts{}, err
		}
		parts.C1 = &c1req
		parts.TKL = append(parts.TKL, coverTKL...)
		parts.Rows = append(parts.Rows, coverRows...)
		// Shuffle BOTH lists. In particular, cover rows must not simply appear
		// after all genuine rows, otherwise the server could use row position as
		// a cover/genuine classifier.
		rand.Shuffle(len(parts.TKL), func(i, j int) { parts.TKL[i], parts.TKL[j] = parts.TKL[j], parts.TKL[i] })
		rand.Shuffle(len(parts.Rows), func(i, j int) { parts.Rows[i], parts.Rows[j] = parts.Rows[j], parts.Rows[i] })
		c.lastC1CoverIDs = coverIDs
	} else {
		c.lastC1CoverIDs = nil
	}
	return parts, nil
}

// PrepareSearch builds the normal user-facing query request. C1 is selected
// automatically for exactly two keywords; n=1 and n>=3 remain plain FDXT.
func (c *Client) PrepareSearch(q []string) SearchRequestParts {
	return c.PrepareSearchWithC1Config(q, DefaultC1Config)
}

// PrepareSearchWithC1Config is the normal automatic dispatcher with an
// explicit C1 cover-budget configuration. It never exposes a baseline/C1
// choice: query size alone selects the scheme.
func (c *Client) PrepareSearchWithC1Config(q []string, cfg C1Config) SearchRequestParts {
	parts, err := c.prepareSearch(q, len(q) == 2, cfg)
	if err != nil {
		panic(err)
	}
	return parts
}

// PrepareSearchBaseline is the internal experiment-only baseline path. It is
// not exposed by the normal CLI as a user-selectable variant.
func (c *Client) PrepareSearchBaseline(q []string) SearchRequestParts {
	parts, err := c.prepareSearch(q, false, C1Config{})
	if err != nil {
		panic(err)
	}
	return parts
}

// PrepareSearchC1 is the internal experiment path for an explicit C1 run.
func (c *Client) PrepareSearchC1(q []string, cfg C1Config) (SearchRequestParts, error) {
	if len(q) != 2 {
		return c.prepareSearch(q, false, C1Config{})
	}
	return c.prepareSearch(q, true, cfg)
}

// RecoverSearch performs the client-side result recovery for a response
// produced by Server.ProcessSearch, then advances the local search counters.
// This is the client half of Algorithm 3 after the server returns RES.
func (c *Client) RecoverSearch(q []string, res []SearchOutput) []string {
	if len(q) == 0 {
		return nil
	}
	w1Idx := 0
	minMax := c.counter(q[0]).Max
	for i, w := range q {
		if m := c.counter(w).Max; m < minMax {
			minMax = m
			w1Idx = i
		}
	}
	w1 := q[w1Idx]
	w1b := []byte(w1)

	added := make(map[string]bool)
	deleted := make(map[string]bool)
	n := len(q)
	for _, r := range res {
		if r.Val == nil || r.Addr == nil {
			continue
		}
		j, ok := c.lookupRowJ(w1, r.Addr)
		if !ok {
			continue
		}
		mask := crypto.F(c.KW, w1b, u64b(j), []byte{1})
		plain := trimZeros(crypto.Xor(r.Val, mask))
		parts := strings.SplitN(string(plain), "|", 2)
		if len(parts) != 2 {
			continue
		}
		id, op := parts[0], parts[1]
		switch Op(op) {
		case OpAdd:
			if r.Cnt == n && !c.lastC1CoverIDs[id] {
				added[id] = true
			}
		case OpDel:
			if r.Cnt > 0 {
				deleted[id] = true
			}
		}
	}

	out := make([]string, 0, len(added))
	for id := range added {
		if !deleted[id] {
			out = append(out, id)
		}
	}

	for _, w := range q {
		ctr := c.counter(w)
		ctr.Srch++
		ctr.Updt = 0
	}
	c.lastC1CoverIDs = nil
	return out
}

// lookupRowJ resolves the encrypted CDB_T address to the client-known update
// index. Genuine rows use the ordinary 1..c_max range; C1 cover rows use the
// reserved client-only pool indices. The address itself is returned by the
// server but contains no plaintext keyword or cover marker.
func (c *Client) lookupRowJ(w string, addr []byte) (uint64, bool) {
	ctr := c.counter(w)
	for j := uint64(1); j <= ctr.Max; j++ {
		if string(crypto.F(c.KW, []byte(w), u64b(j), []byte{0})) == string(addr) {
			return j, true
		}
	}
	for _, ce := range c.C1Pool {
		if pk, ok := ce.Keywords[w]; ok {
			if string(crypto.F(c.KW, []byte(w), u64b(pk.J), []byte{0})) == string(addr) {
				return pk.J, true
			}
		}
	}
	return 0, false
}

// Search runs Algorithm 3 end-to-end locally. It is retained for existing
// tests and is implemented using the same prepare/server/recovery stages
// used by the remote gRPC path.
func (c *Client) Search(srv *Server, q []string) []string {
	// Local/in-process callers (including the deterministic test and demo
	// pipelines) may construct a fresh Client+Server and immediately call
	// Search without running the setup CLI.  C1 is automatic for n=2, so make
	// sure the opaque cover pool exists before preparing the request.  Remote
	// deployments provision the pool during setup and therefore do not use
	// this path.
	if srv != nil && len(q) == 2 {
		if err := c.ensureC1PoolForQuery(srv, q); err != nil {
			panic(err)
		}
	}
	parts := c.PrepareSearch(q)
	res := srv.ProcessQuery(parts)
	return c.RecoverSearch(q, res)
}
