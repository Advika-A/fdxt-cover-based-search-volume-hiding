// Package compare does exact ID-set comparison between an FDXT search
// result and an independent ground-truth ID set. It is pure logic with no
// dependency on MongoDB or the FDXT crypto core, so it is trivially unit
// testable and reusable by both the CLI and future test/experiment code.
package compare

import "sort"

// Result is the outcome of comparing an FDXT-returned ID set against a
// ground-truth ID set for the same query.
type Result struct {
	FDXTIDs        []string // sorted copy of the FDXT result
	GroundTruthIDs []string // sorted copy of the ground-truth result
	ReturnedCount  int
	GroundTruth    int
	FalsePositives []string // in FDXT but NOT in ground truth
	FalseNegatives []string // in ground truth but NOT in FDXT
	ExactMatch     bool     // true iff FDXTIDs and GroundTruthIDs are identical as sets

	// Vacuous is true when both FDXTIDs and GroundTruthIDs are empty, so
	// ExactMatch is trivially true without either side having actually
	// matched anything. This is set-theoretically a correct "exact
	// match" (Compare's own definition doesn't change), but callers that
	// report results to a human -- notably the CLI -- must not present a
	// Vacuous result as a validated "Exact match: true", since it could
	// just as easily mean the dataset/index was empty or the keywords
	// were never indexed at all. Check Vacuous before trusting
	// ExactMatch for anything that claims to have verified something.
	Vacuous bool
}

// Compare does an exact-set comparison, never a mere count comparison:
// ExactMatch is only true when every ID in fdxtIDs is in groundTruthIDs
// and vice versa (i.e. FalsePositives and FalseNegatives are both empty).
func Compare(fdxtIDs, groundTruthIDs []string) Result {
	fdxtSet := toSet(fdxtIDs)
	gtSet := toSet(groundTruthIDs)

	var fp, fn []string
	for id := range fdxtSet {
		if !gtSet[id] {
			fp = append(fp, id)
		}
	}
	for id := range gtSet {
		if !fdxtSet[id] {
			fn = append(fn, id)
		}
	}
	sort.Strings(fp)
	sort.Strings(fn)

	return Result{
		FDXTIDs:        sortedCopy(fdxtIDs),
		GroundTruthIDs: sortedCopy(groundTruthIDs),
		ReturnedCount:  len(fdxtIDs),
		GroundTruth:    len(groundTruthIDs),
		FalsePositives: fp,
		FalseNegatives: fn,
		ExactMatch:     len(fp) == 0 && len(fn) == 0,
		Vacuous:        len(fdxtIDs) == 0 && len(groundTruthIDs) == 0,
	}
}

func toSet(ids []string) map[string]bool {
	m := make(map[string]bool, len(ids))
	for _, id := range ids {
		m[id] = true
	}
	return m
}

func sortedCopy(ids []string) []string {
	out := append([]string(nil), ids...)
	sort.Strings(out)
	return out
}
