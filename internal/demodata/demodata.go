// Package demodata contains the deterministic custom dataset used for
// FDXT integration tests, demos, and the upcoming C1 experiments.
//
// This package deliberately contains no cryptography and no MongoDB code.
// It is the single source of truth for the local dataset: indexing,
// statistics, and ground-truth calculations all derive from the same
// document definitions below.
package demodata

import "sort"

// DatabaseName and CollectionName are retained for compatibility with the
// existing optional MongoDB demo/integration tests. This session does not
// modify the MongoDB implementation.
const (
	DatabaseName   = "FDXT_DEMO"
	CollectionName = "id_keywords"
)

// Record mirrors the id_keywords-shaped data consumed by the existing DB
// and indexer layers: one keyword and all document IDs containing it.
type Record struct {
	Keyword string
	DocIDs  []string
}

type document struct {
	id       string
	keywords []string
}

// documents is the authoritative 50-document local dataset. It is fixed,
// human-readable, deterministic, and intentionally non-uniform. Keyword
// frequencies range from 3 documents for low-frequency terms to 24 for
// high-frequency terms, making it suitable for later leakage/C1 experiments.
var documents = []document{
	{"ID1", []string{"data", "firewall", "network", "security"}},
	{"ID2", []string{"computer", "endpoint", "malware", "security", "system"}},
	{"ID3", []string{"access", "authentication", "computer", "data", "phishing", "security"}},
	{"ID4", []string{"attack", "cloud", "monitoring", "network", "system", "traffic"}},
	{"ID5", []string{"authentication", "cloud", "computer", "data", "database", "software"}},
	{"ID6", []string{"access", "monitoring", "network", "security", "server", "system"}},
	{"ID7", []string{"access", "data", "encryption", "network", "security"}},
	{"ID8", []string{"api", "credential", "email", "network", "phishing", "vulnerability"}},
	{"ID9", []string{"api", "application", "computer", "incident", "monitoring", "software", "web"}},
	{"ID10", []string{"forensics", "malware", "network", "threat"}},
	{"ID11", []string{"access", "firewall", "network", "security", "system"}},
	{"ID12", []string{"computer", "endpoint", "malware", "security", "system"}},
	{"ID13", []string{"access", "authentication", "data", "security", "system"}},
	{"ID14", []string{"access", "attack", "data", "network", "traffic", "web"}},
	{"ID15", []string{"cloud", "computer", "credential", "data", "database", "software"}},
	{"ID16", []string{"access", "attack", "monitoring", "server", "system"}},
	{"ID17", []string{"access", "data", "encryption"}},
	{"ID18", []string{"credential", "email", "malware", "network", "phishing", "system"}},
	{"ID19", []string{"api", "application", "computer", "logging", "software", "system", "web"}},
	{"ID20", []string{"access", "forensics", "malware", "threat", "vulnerability"}},
	{"ID21", []string{"access", "firewall", "network", "security", "system"}},
	{"ID22", []string{"computer", "endpoint", "malware", "network", "security"}},
	{"ID23", []string{"access", "authentication", "backup", "data", "security"}},
	{"ID24", []string{"attack", "data", "logging", "network", "system", "traffic"}},
	{"ID25", []string{"cloud", "data", "database", "incident", "network"}},
	{"ID26", []string{"api", "application", "computer", "monitoring", "server", "software", "system"}},
	{"ID27", []string{"access", "data", "encryption"}},
	{"ID28", []string{"backup", "credential", "email", "phishing", "security", "system"}},
	{"ID29", []string{"api", "application", "computer", "data", "security", "software", "web"}},
	{"ID30", []string{"data", "forensics", "malware", "threat"}},
	{"ID31", []string{"access", "firewall", "network", "security", "vulnerability"}},
	{"ID32", []string{"attack", "computer", "endpoint", "malware", "security"}},
	{"ID33", []string{"access", "authentication", "computer", "credential", "database", "security", "software"}},
	{"ID34", []string{"attack", "authentication", "firewall", "network", "phishing", "traffic"}},
	{"ID35", []string{"access", "cloud", "data", "database", "logging", "network"}},
	{"ID36", []string{"access", "monitoring", "network", "server", "system"}},
	{"ID37", []string{"access", "data", "encryption", "network"}},
	{"ID38", []string{"credential", "data", "email", "incident", "phishing"}},
	{"ID39", []string{"api", "application", "computer", "network", "software", "system", "web"}},
	{"ID40", []string{"malware", "network", "security", "threat"}},
	{"ID41", []string{"firewall", "network", "security", "system"}},
	{"ID42", []string{"computer", "endpoint", "malware", "network", "security"}},
	{"ID43", []string{"access", "authentication", "computer", "database", "security", "software", "system"}},
	{"ID44", []string{"api", "attack", "forensics", "network", "system", "traffic"}},
	{"ID45", []string{"backup", "cloud", "data", "database", "software", "system"}},
	{"ID46", []string{"endpoint", "monitoring", "network", "server", "system"}},
	{"ID47", []string{"access", "data", "encryption"}},
	{"ID48", []string{"authentication", "credential", "email", "forensics", "phishing", "system"}},
	{"ID49", []string{"access", "api", "application", "computer", "data", "software", "web"}},
	{"ID50", []string{"access", "data", "forensics", "malware", "software", "threat"}},
}

// QuerySet is the deterministic query suite used by the dataset statistics
// command and integration tests. Expected result IDs are never stored here;
// counts are computed at runtime from the authoritative dataset.
var QuerySet = []string{
	"security AND network",                  // large n=2
	"access AND data",                       // large n=2
	"computer AND software",                 // large/medium n=2
	"malware AND phishing",                  // very small n=2
	"cloud AND database",                    // medium n=2
	"malware AND vulnerability",             // small n=2
	"authentication AND backup",             // small n=2
	"forensics AND vulnerability",           // very small n=2
	"firewall AND api",                      // zero n=2
	"cloud AND application",                 // zero n=2
	"computer AND security AND network",     // n=3
	"computer AND software AND application", // n=3
	"cloud AND data AND database",           // n=3
	"forensics AND malware AND threat",      // n=3
	"firewall AND api AND web",              // zero n=3
}

// Documents returns a defensive copy of the authoritative ID -> keywords
// representation. The returned slices can be modified by callers without
// changing the dataset used by other tests or commands.
func Documents() map[string][]string {
	out := make(map[string][]string, len(documents))
	for _, d := range documents {
		out[d.id] = append([]string(nil), d.keywords...)
	}
	return out
}

// GetAllDocuments returns the authoritative documents in ID order.
func GetAllDocuments() []RecordDocument {
	out := make([]RecordDocument, 0, len(documents))
	for _, d := range documents {
		out = append(out, RecordDocument{ID: d.id, Keywords: append([]string(nil), d.keywords...)})
	}
	return out
}

// RecordDocument is the public representation of one document.
type RecordDocument struct {
	ID       string
	Keywords []string
}

// GetKeywords returns the keywords belonging to a document ID.
func GetKeywords(documentID string) []string {
	for _, d := range documents {
		if d.id == documentID {
			return append([]string(nil), d.keywords...)
		}
	}
	return nil
}

// GetDocumentsForKeyword returns all document IDs containing keyword,
// sorted in numeric document-ID order.
func GetDocumentsForKeyword(keyword string) []string {
	var out []string
	for _, d := range documents {
		for _, kw := range d.keywords {
			if kw == keyword {
				out = append(out, d.id)
				break
			}
		}
	}
	sortDocumentIDs(out)
	return out
}

// GetDocumentByID returns one document and a boolean indicating whether it
// exists.
func GetDocumentByID(documentID string) (RecordDocument, bool) {
	for _, d := range documents {
		if d.id == documentID {
			return RecordDocument{ID: d.id, Keywords: append([]string(nil), d.keywords...)}, true
		}
	}
	return RecordDocument{}, false
}

// Records inverts the authoritative ID -> keywords dataset into the same
// keyword -> document-ID shape used by the existing MongoDB/indexer layer.
func Records() []Record {
	byKeyword := make(map[string][]string)
	for _, doc := range documents {
		for _, kw := range doc.keywords {
			byKeyword[kw] = append(byKeyword[kw], doc.id)
		}
	}

	keywords := make([]string, 0, len(byKeyword))
	for kw := range byKeyword {
		keywords = append(keywords, kw)
	}
	sort.Strings(keywords)

	recs := make([]Record, 0, len(keywords))
	for _, kw := range keywords {
		ids := append([]string(nil), byKeyword[kw]...)
		sortDocumentIDs(ids)
		recs = append(recs, Record{Keyword: kw, DocIDs: ids})
	}
	return recs
}

// GroundTruth returns the exact conjunctive intersection directly from the
// authoritative dataset. It never calls FDXT.
func GroundTruth(keywords []string) []string {
	if len(keywords) == 0 {
		return nil
	}
	sets := make([]map[string]bool, 0, len(keywords))
	for _, kw := range keywords {
		set := make(map[string]bool)
		for _, id := range GetDocumentsForKeyword(kw) {
			set[id] = true
		}
		sets = append(sets, set)
	}

	out := make([]string, 0)
	for id := range sets[0] {
		ok := true
		for i := 1; i < len(sets); i++ {
			if !sets[i][id] {
				ok = false
				break
			}
		}
		if ok {
			out = append(out, id)
		}
	}
	sortDocumentIDs(out)
	return out
}

// GroundTruthConjunctiveQuery is the explicit API name used by the
// dataset specification; it is equivalent to GroundTruth.
func GroundTruthConjunctiveQuery(keywords []string) []string {
	return GroundTruth(keywords)
}

// KeywordFrequencies returns keyword -> number of documents containing it.
func KeywordFrequencies() map[string]int {
	freq := make(map[string]int)
	for _, r := range Records() {
		freq[r.Keyword] = len(r.DocIDs)
	}
	return freq
}

// Keywords returns the complete vocabulary in lexical order.
func Keywords() []string {
	freq := KeywordFrequencies()
	out := make([]string, 0, len(freq))
	for kw := range freq {
		out = append(out, kw)
	}
	sort.Strings(out)
	return out
}

// DatasetStats contains runtime-derived dataset statistics.
type DatasetStats struct {
	Documents        int
	DistinctKeywords int
	KeywordFrequency map[string]int
}

// Stats calculates all dataset statistics from the authoritative dataset.
func Stats() DatasetStats {
	freq := KeywordFrequencies()
	return DatasetStats{
		Documents:        len(documents),
		DistinctKeywords: len(freq),
		KeywordFrequency: freq,
	}
}

// DocumentCount returns the number of dataset documents.
func DocumentCount() int { return len(documents) }

func sortDocumentIDs(ids []string) {
	sort.Slice(ids, func(i, j int) bool {
		return documentNumber(ids[i]) < documentNumber(ids[j])
	})
}

func documentNumber(id string) int {
	if len(id) < 3 {
		return 0
	}
	n := 0
	for _, r := range id[2:] {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}
