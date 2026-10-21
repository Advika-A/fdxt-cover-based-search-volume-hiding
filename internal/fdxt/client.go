package fdxt

import (
	"crypto/rand"

	"fdxt-poc/internal/crypto"
)

// Client holds the secret key sk = (KW, KT, KX, KY, KZ) and the local
// state-set Count[w] = (c_srch_w, c_updt_w, c_max_w) for every keyword w,
// matching Algorithm 1 (Setup, client section).
type Client struct {
	KW, KT, KX, KY, KZ []byte
	P                  *crypto.Ctx
	Count              map[string]*Counter
	C1Pool             []C1PoolEntry
	lastC1CoverIDs     map[string]bool
}

// NewClient runs the client side of Setup(lambda): sample sk, initialize
// the empty state-set.
func NewClient(p *crypto.Ctx) *Client {
	return &Client{
		KW:             randKey(),
		KT:             randKey(),
		KX:             randKey(),
		KY:             randKey(),
		KZ:             randKey(),
		P:              p,
		Count:          make(map[string]*Counter),
		C1Pool:         nil,
		lastC1CoverIDs: nil,
	}
}

func randKey() []byte {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic("fdxt: failed to sample key: " + err.Error())
	}
	return b
}

func (c *Client) counter(w string) *Counter {
	ctr, ok := c.Count[w]
	if !ok {
		ctr = &Counter{}
		c.Count[w] = ctr
	}
	return ctr
}

// Update runs Algorithm 2 (Update), client section, for a single
// keyword-and-document-identifier entry (w, id) with operation op, and
// uploads the resulting ciphertext tuple to the server.
//
// This is the heart of the FDXT construction:
//
//	xtag  = g^( Fp(KX,w) * Fp(KY,id||op) )     -- the cross-tag
//	X     = H(xtag)
//	C     = X XOR T                             -- forward-private ciphertext
//	alpha = Fp(KY,id||op) * Fp(KZ,w||c_max_w)^-1 -- blinding factor
//
// alpha lets the server later "unblind" a candidate cross-token
// xtk = g^(Fp(KX,wk) * Fp(KZ,w||j)) into xtk^alpha = xtag(wk,id,op),
// without ever learning Fp(KY,id||op) itself.
func (c *Client) Update(srv *Server, op Op, w string, id string) {
	ctr := c.counter(w)
	ctr.Updt++
	ctr.Max++

	wb := []byte(w)
	maxB := u64b(ctr.Max)
	srchB := u64b(ctr.Srch)
	updtB := u64b(ctr.Updt)

	// addr, val: single-keyword DSSE ciphertext (CDB_T entry, part 1)
	addr := crypto.F(c.KW, wb, maxB, []byte{0})
	mask := crypto.F(c.KW, wb, maxB, []byte{1})
	plain := padTo([]byte(id+"|"+string(op)), len(mask))
	val := crypto.Xor(plain, mask)

	// L, T: address + trapdoor for the cross-tag ciphertext (CDB_X entry)
	L := crypto.F(c.KT, wb, srchB, updtB, []byte{0})
	T := crypto.F(c.KT, wb, srchB, updtB, []byte{1})

	// xtag = g^(Fp(KX,w) * Fp(KY,id||op))
	idOp := []byte(id + "|" + string(op))
	fx := crypto.Fp(c.P, c.KX, wb)
	fy := crypto.Fp(c.P, c.KY, idOp)
	exp := c.P.Pairing.NewZr().Mul(fx, fy)
	xtag := c.P.Pairing.NewG1().PowZn(c.P.Generator, exp)

	X := crypto.H(xtag.Bytes())
	C := crypto.Xor(X, T)

	// alpha = Fp(KY,id||op) * Fp(KZ,w||c_max_w)^-1
	fz := crypto.Fp(c.P, c.KZ, wb, maxB)
	fzInv := c.P.Pairing.NewZr().Invert(fz)
	alpha := c.P.Pairing.NewZr().Mul(fy, fzInv)

	srv.StoreUpdate(string(L), C, string(addr), val, alpha)
}
