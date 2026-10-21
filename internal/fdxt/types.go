package fdxt

// Op is the update operation type, op in {add, del} in the paper.
type Op string

const (
	OpAdd Op = "add"
	OpDel Op = "del"
)

// Counter holds (c_srch_w, c_updt_w, c_max_w) for a single keyword w, as
// defined in Table II of the paper:
//
//	c_srch_w: total number of times w has been searched.
//	c_updt_w: number of updates to w since its last search.
//	c_max_w:  total number of times w has ever been updated.
type Counter struct {
	Srch uint64
	Updt uint64
	Max  uint64
}
