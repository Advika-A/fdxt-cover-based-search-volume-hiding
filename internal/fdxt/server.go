package fdxt

import (
	"fdxt-poc/internal/crypto"

	"github.com/Nik-U/pbc"
)

// CDBTEntry is a stored tuple (val, alpha) for a single-keyword update
// entry, as written by Algorithm 2 (Update, server section):
//
//	CDB_T[addr] <- (val, alpha)
type CDBTEntry struct {
	Val   []byte
	Alpha *pbc.Element
}

// Server holds the honest-but-curious server's state: CDB_X (unaccessed
// cross-tag ciphertexts), CDB_T (single-keyword index/ciphertext store),
// and XSET (the permanent map of revealed cross-tags), matching Fig. 1 and
// Algorithms 2-3 of the paper.
type Server struct {
	P *crypto.Ctx

	CDBX map[string][]byte    // L -> C
	CDBT map[string]CDBTEntry // addr -> (val, alpha)
	XSET map[string]bool      // H(xtag) -> true
}

// NewServer creates an empty server state (Setup, server section).
func NewServer(p *crypto.Ctx) *Server {
	return &Server{
		P:    p,
		CDBX: make(map[string][]byte),
		CDBT: make(map[string]CDBTEntry),
		XSET: make(map[string]bool),
	}
}

// StoreUpdate is the server side of Algorithm 2 (Update):
//
//	CDB_X[L] <- C ; CDB_T[addr] <- (val, alpha)
func (s *Server) StoreUpdate(L string, C []byte, addr string, val []byte, alpha *pbc.Element) {
	s.CDBX[L] = C
	s.CDBT[addr] = CDBTEntry{Val: val, Alpha: alpha}
}

// TKLEntry is one (L, T) pair from the client's cross-tag token list TKL,
// used by the server to locate + decrypt an unaccessed cross-tag and move
// it permanently into XSET.
type TKLEntry struct {
	L []byte
	T []byte
}

// XTKRow is one row of the client's (STKL, XTKL): the address of the j-th
// update entry of w1 in CDB_T, together with the cross-terms xtk_{j,k} for
// every remaining keyword wk in the conjunctive query.
type XTKRow struct {
	Addr []byte
	XTK  map[string]*pbc.Element // keyword -> xtk_{j,k}
}

// SearchOutput is one entry of RES: the (still-encrypted) val for the j-th
// update of w1, together with the match count across the remaining
// keywords (cnt == n means "matches all n keywords").
type SearchOutput struct {
	Addr []byte
	Val  []byte
	Cnt  int
}

// ProcessSearch is the server side of Algorithm 3 (Search).
//
// Step 1 (lines 1-7, server): for every (L, T) in TKL, decrypt the stored
// cross-tag ciphertext C = CDB_X[L] via X = C XOR T = xtag, and permanently
// record XSET[X] = 1. This is exactly what gives FDXT (and ODXT) forward
// privacy: cross-tags sit encrypted and unlinkable in CDB_X until the
// keyword that owns them is actually searched.
//
// Step 2 (lines 8-17, server): for each row j of (STKL, XTKL) -- i.e. for
// each historical update to w1 -- fetch (val_j, alpha_j) from CDB_T, then
// for every remaining keyword wk obliviously recompute the candidate
// cross-tag X_{j,k} = H(xtk_{j,k}^{alpha_j}) and test membership in XSET.
// cnt starts at 1 (w1 itself always "matches") and is incremented once per
// remaining keyword that also matches, exactly as in the paper.
func (s *Server) ProcessSearch(tkl []TKLEntry, rows []XTKRow, restKeywords []string) []SearchOutput {
	// --- Step 1: reveal cross-tags relevant to this query into XSET ---
	for _, e := range tkl {
		C, ok := s.CDBX[string(e.L)]
		if !ok {
			continue
		}
		X := crypto.Xor(C, e.T)
		s.XSET[string(X)] = true
	}

	// --- Step 2: match w1's entries against the remaining keywords ---
	out := make([]SearchOutput, len(rows))
	for i, row := range rows {
		entry, ok := s.CDBT[string(row.Addr)]
		if !ok {
			continue // no such update entry (should not happen in this POC)
		}
		cnt := 1
		for _, wk := range restKeywords {
			xtk := row.XTK[wk]
			xtkAlpha := s.P.Pairing.NewG1().PowZn(xtk, entry.Alpha)
			Xjk := crypto.H(xtkAlpha.Bytes())
			if s.XSET[string(Xjk)] {
				cnt++
			}
		}
		out[i] = SearchOutput{Addr: append([]byte(nil), row.Addr...), Val: entry.Val, Cnt: cnt}
	}
	return out
}

// ProcessQuery is the server-side entry point. C1 is deliberately not
// represented on the wire: its cover rows/TKL entries are ordinary-looking
// protocol material and are processed by the exact same FDXT Search path.
func (s *Server) ProcessQuery(parts SearchRequestParts) []SearchOutput {
	return s.ProcessSearch(parts.TKL, parts.Rows, parts.RestKeywords)
}
