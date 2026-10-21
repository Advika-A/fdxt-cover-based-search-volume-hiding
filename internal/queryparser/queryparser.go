// Package queryparser turns a human-typed query string such as
// "F1227 AND F1438" into the ordered keyword list the FDXT Search
// protocol expects. It is pure string handling -- no crypto, no DB.
package queryparser

import (
	"fmt"
	"regexp"
	"strings"
)

// andSplitter matches a whole "AND" token using word boundaries rather
// than requiring literal surrounding whitespace. This correctly avoids
// splitting inside a keyword that merely contains "AND" (e.g. "BRAND",
// where the "AND" substring is not bounded by a word/non-word transition
// on its left side), while still splitting a trailing/leading "AND" that
// has no whitespace on one side because TrimSpace already removed it
// (e.g. "F1 AND " -> trimmed to "F1 AND", which still needs to split into
// "F1" and an empty tail).
var andSplitter = regexp.MustCompile(`(?i)\bAND\b`)

// Parse splits raw on (case-insensitive, whitespace-delimited) "AND" and
// trims each resulting keyword. It returns an error for an empty query or
// a query that reduces to zero non-empty keywords (e.g. "  AND  ").
//
// Examples:
//
//	Parse("F1227 AND F1438")   -> ["F1227", "F1438"], nil
//	Parse("F1227 and f1438")   -> ["F1227", "f1438"], nil   (case preserved in the keyword itself)
//	Parse("K1")                -> ["K1"], nil
//	Parse("")                  -> nil, error
func Parse(raw string) ([]string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return nil, fmt.Errorf("queryparser: empty query")
	}

	parts := andSplitter.Split(trimmed, -1)
	keywords := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		keywords = append(keywords, p)
	}
	if len(keywords) == 0 {
		return nil, fmt.Errorf("queryparser: query %q contains no keywords", raw)
	}
	return keywords, nil
}
