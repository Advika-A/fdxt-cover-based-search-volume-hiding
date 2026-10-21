// Command demo runs the corrected FDXT protocol against the deterministic
// 50-document local dataset. It is deliberately database-free and network-
// free; the separate-process gRPC path is exercised by cmd/client/cmd/server.
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/queryparser"
)

func main() {
	statsOnly := flag.Bool("dataset-stats", false, "print dataset statistics and the deterministic experiment query set")
	flag.Parse()

	if *statsOnly {
		printStats()
		return
	}

	p := crypto.Init()
	client := fdxt.NewClient(p)
	server := fdxt.NewServer(p)

	fmt.Println("FDXT 50-document local demo")
	fmt.Println("============================")
	fmt.Printf("Documents: %d\n", demodata.DocumentCount())
	fmt.Printf("Distinct keywords: %d\n", len(demodata.Keywords()))
	fmt.Printf("Keyword-ID pairs: %d\n", pairCount())
	fmt.Println()

	fmt.Println("Indexing 50 documents into the in-memory FDXT baseline...")
	pool, err := client.BuildC1Pool(demodata.Keywords(), 64)
	if err != nil {
		panic(err)
	}
	type event struct {
		cover *fdxt.C1PoolEntry
		doc   *demodata.RecordDocument
		kw    string
	}
	var events []event
	for _, d := range demodata.GetAllDocuments() {
		for _, kw := range d.Keywords {
			dc := d
			events = append(events, event{doc: &dc, kw: kw})
		}
	}
	for i := range pool {
		ce := pool[i]
		events = append(events, event{cover: &ce})
	}
	for i := len(events) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(err)
		}
		j := int(n.Int64())
		events[i], events[j] = events[j], events[i]
	}
	for _, e := range events {
		if e.cover != nil {
			if err := client.InstallC1Entry(server, *e.cover); err != nil {
				panic(err)
			}
		} else {
			client.Update(server, fdxt.OpAdd, e.kw, e.doc.ID)
		}
	}
	client.C1Pool = pool
	fmt.Println("Indexing complete.")
	fmt.Println()

	fmt.Println("Representative conjunctive queries:")
	for _, raw := range demodata.QuerySet {
		q, err := queryparser.Parse(raw)
		if err != nil {
			panic(err)
		}
		// Indexing is deliberately performed once above; each query then
		// follows the real in-memory FDXT Prepare -> Server.ProcessSearch ->
		// Recover pipeline.
		got := client.Search(server, q)
		ground := demodata.GroundTruthConjunctiveQuery(q)
		res := compare.Compare(got, ground)
		fmt.Printf("%-42s count=%-2d exact=%v", raw, res.ReturnedCount, res.ExactMatch)
		if res.Vacuous {
			fmt.Print(" (zero-result query)")
		}
		fmt.Println()
	}
}

func pairCount() int {
	n := 0
	for _, r := range demodata.Records() {
		n += len(r.DocIDs)
	}
	return n
}

func printStats() {
	stats := demodata.Stats()
	fmt.Println("FDXT 50-document dataset statistics")
	fmt.Println("====================================")
	fmt.Printf("Documents: %d\n", stats.Documents)
	fmt.Printf("Distinct keywords: %d\n", stats.DistinctKeywords)
	fmt.Printf("Keyword-ID pairs: %d\n", pairCount())
	fmt.Println()

	keywords := append([]string(nil), demodata.Keywords()...)
	sort.Slice(keywords, func(i, j int) bool {
		fi, fj := stats.KeywordFrequency[keywords[i]], stats.KeywordFrequency[keywords[j]]
		if fi != fj {
			return fi > fj
		}
		return keywords[i] < keywords[j]
	})
	fmt.Println("Keyword -> document frequency")
	for _, kw := range keywords {
		fmt.Printf("  %-16s %d\n", kw, stats.KeywordFrequency[kw])
	}

	fmt.Println()
	fmt.Println("Deterministic experiment query set")
	for i, q := range demodata.QuerySet {
		keywords, err := queryparser.Parse(q)
		if err != nil {
			panic(err)
		}
		ground := demodata.GroundTruthConjunctiveQuery(keywords)
		kind := "n=2"
		if strings.Count(q, " AND ") == 2 {
			kind = "n=3"
		}
		if len(ground) == 0 {
			kind += ", zero-result"
		}
		fmt.Printf("  %2d. %-42s -> %d (%s)\n", i+1, q, len(ground), kind)
	}
}
