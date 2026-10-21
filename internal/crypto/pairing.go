// Package crypto provides the single, process-wide pairing context used by
// the FDXT proof of concept, plus the PRF/hash primitives built on top of it.
//
// Design constraint (per spec): ONE pairing, ONE generator, for the whole
// lifetime of the program. We never call pbc.NewParamsFromString more than
// once, and the generator is derived deterministically (hash-to-G1) instead
// of being independently re-randomized anywhere else in the codebase.
package crypto

import (
	"sync"

	"github.com/Nik-U/pbc"
)

// ParamsA is a fixed, standard Type-A pairing parameter set (symmetric
// pairing, the same family used throughout the FDXT paper's construction:
// xtag = g^(...) in a group G with a bilinear map). It is embedded as a
// constant so the parameters are identical on every run.
const ParamsA = `type a
q 8780710799663312522437781984754049815806883199414208211028653399266475630880222957078625179422662221423155858769582317459277713367317481324925129998224791
h 12016012264891146079388821366740534204802954401251311822919615131047207289359704531102844802183906537786776
r 730750818665451621361119245571504901405976559617
exp2 159
exp1 107
sign1 1
sign0 1
`

// fixedGeneratorSeed is hashed into G1 to obtain a deterministic generator g.
// Using a fixed, published seed (rather than g.Rand()) means every run of
// the POC, and every component within a single run, shares exactly the same
// generator, which is required for the cross-tag algebra to be consistent.
const fixedGeneratorSeed = "FDXT-POC-FIXED-GENERATOR-v1"

// Ctx bundles the one pairing and one generator used for the whole process.
type Ctx struct {
	Pairing   *pbc.Pairing
	Generator *pbc.Element
}

var (
	once   sync.Once
	global *Ctx
)

// Init initializes (on first call) and returns the single, process-wide
// pairing context. Subsequent calls return the same *Ctx. This is the only
// place pairing parameters are ever parsed.
func Init() *Ctx {
	once.Do(func() {
		params, err := pbc.NewParamsFromString(ParamsA)
		if err != nil {
			panic("fdxt: failed to parse pairing parameters: " + err.Error())
		}
		pairing := params.NewPairing()
		g := pairing.NewG1().SetFromHash([]byte(fixedGeneratorSeed))
		global = &Ctx{Pairing: pairing, Generator: g}
	})
	return global
}
