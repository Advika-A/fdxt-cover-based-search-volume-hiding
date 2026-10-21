package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/grpcwire"
	"fdxt-poc/internal/queryparser"
)

func main() {
	query := flag.String("query", "", "conjunctive query, e.g. \"computer AND security\"")
	addr := flag.String("addr", "127.0.0.1:8443", "server address")
	state := flag.String("state", "state/client.json", "client state")
	cert := flag.String("cert", "state/server.crt", "server certificate")
	dataset := flag.String("dataset", "50", "local dataset label (currently 50; state must be built from the same dataset)")
	c1Budget := flag.Int("c1-cover-budget", fdxt.DefaultC1Config.MaxCover, "maximum adaptive C1 cover entries for n=2")
	flag.Parse()
	if *dataset != "50" {
		fmt.Fprintln(os.Stderr, "unsupported dataset:", *dataset, "(this build uses the deterministic 50-document dataset)")
		os.Exit(2)
	}
	if strings.TrimSpace(*query) == "" {
		fmt.Fprintln(os.Stderr, "usage: go run ./cmd/client --query \"computer AND security\" --dataset 50")
		os.Exit(2)
	}

	q, err := queryparser.Parse(*query)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	p := crypto.Init()
	client, err := fdxt.LoadClientState(*state, p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "load client state:", err)
		os.Exit(1)
	}
	if *c1Budget < 0 {
		fmt.Fprintln(os.Stderr, "c1-cover-budget must be >= 0")
		os.Exit(2)
	}
	parts := client.PrepareSearchWithC1Config(q, fdxt.C1Config{MaxCover: *c1Budget})

	rpc, err := grpcwire.NewClient("https://"+*addr+"/fdxt.FDXT/Search", *cert)
	if err != nil {
		fmt.Fprintln(os.Stderr, "gRPC client:", err)
		os.Exit(1)
	}
	start := time.Now()
	res, err := rpc.Search(parts)
	rpcTime := time.Since(start)
	if err != nil {
		fmt.Fprintln(os.Stderr, "gRPC search:", err)
		os.Exit(1)
	}
	recovered := client.RecoverSearch(q, res)
	if err := client.SaveState(*state); err != nil {
		fmt.Fprintln(os.Stderr, "save client state:", err)
		os.Exit(1)
	}

	sort.Strings(recovered)
	ground := demodata.GroundTruth(q)
	result := compare.Compare(recovered, ground)
	fmt.Println("[CLIENT] Query:", *query)
	fmt.Println("[CLIENT] n:", len(q))
	fmt.Println("[CLIENT] Sending actual FDXT trapdoor over gRPC")
	if parts.C1 != nil {
		fmt.Printf("[CLIENT] C1: automatic n=2 protocol-shaped cover entries enabled (budget=%d)\n", parts.C1.Config.MaxCover)
	} else {
		fmt.Println("[CLIENT] C1: not applied (n != 2)")
	}
	fmt.Println("[CLIENT] Recovered IDs:", result.FDXTIDs)
	fmt.Println("[CLIENT] Ground-truth IDs:", result.GroundTruthIDs)
	fmt.Println("[CLIENT] Returned count:", result.ReturnedCount)
	fmt.Println("[CLIENT] Ground-truth count:", result.GroundTruth)
	fmt.Println("[CLIENT] Exact match:", result.ExactMatch)
	fmt.Println("[CLIENT] False positives:", result.FalsePositives)
	fmt.Println("[CLIENT] False negatives:", result.FalseNegatives)
	fmt.Println("[CLIENT] gRPC round-trip:", rpcTime)
	if parts.C1 != nil {
		cover := len(parts.Rows) - c1BaselineRows(client, q)
		if cover < 0 {
			cover = 0
		}
		fmt.Println("[CLIENT] C1 baseline response volume:", len(parts.Rows)-cover)
		fmt.Println("[CLIENT] C1 observed response volume:", len(res))
		fmt.Println("[CLIENT] C1 cover volume:", cover)
	}
}

func c1BaselineRows(client *fdxt.Client, q []string) int {
	return len(client.PrepareSearchBaseline(q).Rows)
}
