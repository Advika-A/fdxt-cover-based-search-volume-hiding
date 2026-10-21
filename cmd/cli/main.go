// Command cli connects to a real FDXT MongoDB dataset (by default,
// Crime_USENIX_REV.id_keywords, documents shaped as {"k": "<keyword>",
// "val_set": ["<id1>", "<id2>", ...]} -- one document per keyword),
// loads it into the real FDXT encrypted index (the unmodified Session 1
// cryptographic core), runs a real conjunctive search through that
// index, and prints an exact ID-set comparison against a live MongoDB
// ground-truth query.
//
// Usage (real dataset, the default):
//
//	go run ./cmd/cli query "computer AND security"
//
// Usage (small built-in demo dataset -- see internal/demodata):
//
//	go run ./cmd/cli demo
//	go run ./cmd/cli demo "security AND network"
//	go run ./cmd/cli -demo query "computer AND security"
//
// Demo mode runs through the exact same real FDXT pipeline (Setup ->
// Update -> trapdoor generation -> Search -> result recovery) against a
// separate, dedicated MongoDB database (FDXT_DEMO.id_keywords, ~18
// documents) instead of the full Crime_USENIX_REV collection, so a
// faculty demonstration doesn't have to wait for a full-dataset index
// build. It auto-seeds that database from internal/demodata's literal
// dataset (idempotent -- safe to re-run) before indexing and querying it.
//
// Flags (all optional, shown with defaults; note flags must be given
// BEFORE the "query"/"demo" subcommand, matching this CLI's existing
// flag.FlagSet-based parsing):
//
//	-uri          mongodb://localhost:27017
//	-db           Crime_USENIX_REV   (FDXT_DEMO in demo mode, unless set explicitly)
//	-collection   id_keywords
//	-limit        0     (0 = index the entire collection)
//	-timeout      5m
//	-demo         false (equivalent to using the "demo" subcommand)
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/db"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/indexer"
	"fdxt-poc/internal/queryparser"
)

// defaultDemoQuery is used when "demo" is invoked with no query string of
// its own -- it is the primary documented sanity check for the demo
// dataset.
const defaultDemoQuery = "computer AND security"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr *os.File) int {
	fs := flag.NewFlagSet("cli", flag.ContinueOnError)
	fs.SetOutput(stderr)

	uri := fs.String("uri", "mongodb://localhost:27017", "MongoDB connection URI")
	dbName := fs.String("db", "Crime_USENIX_REV", "MongoDB database name (ignored in demo mode unless set explicitly)")
	collName := fs.String("collection", "id_keywords", "MongoDB collection name (real FDXT schema: {k, val_set}, one document per keyword)")
	limit := fs.Int64("limit", 0, "max number of keyword documents to index into FDXT (0 = entire collection)")
	timeout := fs.Duration("timeout", 5*time.Minute, "overall operation timeout")
	demoFlag := fs.Bool("demo", false, "use the small built-in FDXT_DEMO.id_keywords demo dataset instead of the real dataset (auto-seeded); same effect as the \"demo\" subcommand")

	fs.Usage = func() {
		fmt.Fprintln(stderr, `usage:`)
		fmt.Fprintln(stderr, `  cli [flags] query "KEYWORD1 AND KEYWORD2"   (real dataset, default)`)
		fmt.Fprintln(stderr, `  cli [flags] demo ["KEYWORD1 AND KEYWORD2"]  (small built-in demo dataset)`)
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return 2
	}

	var explicitDB, explicitColl bool
	fs.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "db":
			explicitDB = true
		case "collection":
			explicitColl = true
		}
	})

	rest := fs.Args()
	if len(rest) < 1 {
		fs.Usage()
		return 2
	}

	var isDemo bool
	var rawQuery string
	switch rest[0] {
	case "demo":
		isDemo = true
		if len(rest) > 1 {
			rawQuery = strings.Join(rest[1:], " ")
		} else {
			rawQuery = defaultDemoQuery
		}
	case "query":
		if len(rest) < 2 {
			fs.Usage()
			return 2
		}
		rawQuery = strings.Join(rest[1:], " ")
		isDemo = *demoFlag
	default:
		fs.Usage()
		return 2
	}

	keywords, err := queryparser.Parse(rawQuery)
	if err != nil {
		fmt.Fprintln(stderr, "invalid query:", err)
		return 2
	}

	effectiveDB, effectiveColl := *dbName, *collName
	if isDemo {
		if !explicitDB {
			effectiveDB = demodata.DatabaseName
		}
		if !explicitColl {
			effectiveColl = demodata.CollectionName
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	cfg := db.Config{URI: *uri, Database: effectiveDB, Collection: effectiveColl}
	store, err := db.Connect(ctx, cfg)
	if err != nil {
		fmt.Fprintln(stderr, "MongoDB connection failed:", err)
		return 1
	}
	defer store.Close(context.Background())

	if isDemo {
		fmt.Fprintln(stdout, "========================================")
		fmt.Fprintln(stdout, "FDXT DEMO")
		fmt.Fprintln(stdout, "========================================")
		fmt.Fprintln(stdout)

		// Auto-initialize the demo database from the literal dataset in
		// internal/demodata. This is an upsert keyed on keyword, so
		// re-running "demo" is always safe and idempotent -- it never
		// touches Crime_USENIX_REV, only FDXT_DEMO.id_keywords.
		seeded, err := store.Seed(ctx, toDBRecords(demodata.Records()))
		if err != nil {
			fmt.Fprintln(stderr, "demo dataset seeding failed:", err)
			return 1
		}
		fmt.Fprintf(stdout, "Seeded %d demo keyword documents into %s.\n\n", seeded, store.CollectionPath())
	} else {
		fmt.Fprintln(stdout, "Connected to MongoDB:", store.CollectionPath())
	}

	// Refuse to proceed against a missing, empty, or wrongly-shaped
	// collection: indexing zero (keyword,id) pairs would make FDXT's
	// search trivially return nothing, which would then trivially
	// "match" an equally-empty ground-truth result -- a vacuous pass
	// that verifies nothing. See also the Vacuous check on the final
	// result below, which catches the same failure mode even if it
	// only shows up per-query rather than for the whole collection.
	if err := store.ValidateSchema(ctx); err != nil {
		fmt.Fprintln(stderr, "dataset validation failed:", err)
		return 1
	}

	docCount, _ := store.EstimatedDocumentCount(ctx)

	if isDemo {
		fmt.Fprintf(stdout, "Database: %s\n", effectiveDB)
		fmt.Fprintf(stdout, "Collection: %s\n", effectiveColl)
		fmt.Fprintf(stdout, "Documents: %d\n", demodata.DocumentCount())
	} else {
		fmt.Fprintf(stdout, "Estimated keyword documents in collection: %d\n", docCount)
		fmt.Fprintln(stdout, "Indexing dataset into the FDXT encrypted index (this may take a while for large collections)...")
	}

	// Progress reporting: print every 10 pairs for the tiny demo dataset,
	// every 1000 for the (potentially much larger) real dataset, so
	// neither mode is silent nor floods the terminal.
	progressEvery := int64(1000)
	if isDemo {
		progressEvery = 10
	}
	progress := func(pairs int64) {
		fmt.Fprintf(stdout, "Indexed %d pairs...\n", pairs)
	}

	p := crypto.Init()
	client, server, stats, err := indexer.BuildFromMongo(ctx, store, p, *limit, progressEvery, progress)
	if err != nil {
		fmt.Fprintln(stderr, "indexing failed:", err)
		return 1
	}
	if stats.Pairs == 0 {
		fmt.Fprintf(stderr, "indexing failed: 0 (keyword,id) pairs indexed from %s -- collection "+
			"reported documents but none had a usable \"k\" string field plus non-empty \"val_set\" "+
			"array (schema mismatch?)\n", store.CollectionPath())
		return 1
	}

	if isDemo {
		fmt.Fprintf(stdout, "Keyword-ID pairs: %d\n\n", stats.Pairs)
		fmt.Fprintf(stdout, "Query: %s\n\n", rawQuery)
	} else {
		fmt.Fprintf(stdout, "Indexed %d keywords / %d (keyword,id) pairs.\n\n", stats.Keywords, stats.Pairs)
	}

	fdxtIDs := client.Search(server, keywords)

	groundTruthIDs, err := store.KeywordGroundTruth(ctx, keywords)
	if err != nil {
		fmt.Fprintln(stderr, "ground-truth query failed:", err)
		return 1
	}

	result := compare.Compare(fdxtIDs, groundTruthIDs)
	if isDemo {
		printDemoResult(stdout, result)
	} else {
		printResult(stdout, rawQuery, result)
	}

	if result.Vacuous {
		// Never report success from an empty-vs-empty comparison: it is
		// indistinguishable from "the keywords were never indexed at
		// all" and does not validate the protocol.
		return 1
	}
	if !result.ExactMatch {
		return 1
	}
	return 0
}

// toDBRecords converts demodata.Record values into db.Record values.
// internal/demodata deliberately has no dependency on internal/db (it's
// pure literal data), so this tiny, obviously-correct field-for-field
// conversion lives here, at the one place that needs both types.
func toDBRecords(recs []demodata.Record) []db.Record {
	out := make([]db.Record, 0, len(recs))
	for _, r := range recs {
		out = append(out, db.Record{Keyword: r.Keyword, DocIDs: r.DocIDs})
	}
	return out
}

func printResult(w *os.File, rawQuery string, r compare.Result) {
	fmt.Fprintln(w, "Query:", rawQuery)
	fmt.Fprintln(w, "FDXT IDs:", r.FDXTIDs)
	fmt.Fprintln(w, "Ground-truth IDs:", r.GroundTruthIDs)
	fmt.Fprintln(w, "Returned count:", r.ReturnedCount)
	fmt.Fprintln(w, "Ground-truth count:", r.GroundTruth)
	if r.Vacuous {
		fmt.Fprintln(w, "Exact match: INCONCLUSIVE (both result sets are empty -- this does not "+
			"validate anything; check that the queried keywords actually exist in the dataset)")
	} else {
		fmt.Fprintln(w, "Exact match:", r.ExactMatch)
	}
	fmt.Fprintln(w, "False positives:", r.FalsePositives)
	fmt.Fprintln(w, "False negatives:", r.FalseNegatives)
}

// printDemoResult matches the output shape from the Session 2 demo-mode
// request: FDXT's returned IDs and the independently-computed
// ground-truth IDs are both printed and compared as exact sets, never
// hardcoded -- see indexer.BuildFromMongo (real FDXT protocol) and
// db.Store.KeywordGroundTruth (independent MongoDB oracle) above.
func printDemoResult(w *os.File, r compare.Result) {
	fmt.Fprintln(w, "FDXT returned IDs:")
	fmt.Fprintln(w, r.FDXTIDs)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Ground-truth IDs:")
	fmt.Fprintln(w, r.GroundTruthIDs)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Returned count:", r.ReturnedCount)
	fmt.Fprintln(w, "Ground-truth count:", r.GroundTruth)
	if r.Vacuous {
		fmt.Fprintln(w, "Exact match: INCONCLUSIVE (both result sets are empty -- this does not "+
			"validate anything; check that the demo dataset seeded correctly)")
	} else {
		fmt.Fprintln(w, "Exact match:", r.ExactMatch)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "False positives:", r.FalsePositives)
	fmt.Fprintln(w, "False negatives:", r.FalseNegatives)
}
