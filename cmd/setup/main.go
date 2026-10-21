package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/fdxt"
)

func main() {
	stateDir := flag.String("state", "state", "directory for client/server state")
	dataset := flag.String("dataset", "50", "local dataset to index (currently only 50 is supported)")
	flag.Parse()
	if *dataset != "50" {
		fmt.Fprintf(os.Stderr, "unsupported dataset %q: only the deterministic 50-document dataset is supported\n", *dataset)
		os.Exit(2)
	}
	if err := os.MkdirAll(*stateDir, 0700); err != nil {
		panic(err)
	}

	p := crypto.Init()
	client := fdxt.NewClient(p)
	server := fdxt.NewServer(p)
	pairs := 0
	pool, err := client.BuildC1Pool(demodata.Keywords(), 64)
	if err != nil {
		panic(err)
	}
	type event struct {
		cover       *fdxt.C1PoolEntry
		keyword, id string
	}
	var events []event
	for _, r := range demodata.Records() {
		for _, id := range r.DocIDs {
			events = append(events, event{keyword: r.Keyword, id: id})
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
			client.Update(server, fdxt.OpAdd, e.keyword, e.id)
			pairs++
		}
	}
	client.C1Pool = pool
	clientPath := filepath.Join(*stateDir, "client.json")
	serverPath := filepath.Join(*stateDir, "server.json")
	if err := client.SaveState(clientPath); err != nil {
		panic(err)
	}
	if err := server.SaveState(serverPath); err != nil {
		panic(err)
	}
	fmt.Printf("FDXT setup complete.\nDataset: deterministic local 50-document dataset\nDocuments: %d\nDistinct keywords: %d\nKeyword-ID pairs: %d\nClient state: %s\nServer state: %s\n", demodata.DocumentCount(), len(demodata.Keywords()), pairs, clientPath, serverPath)
}
