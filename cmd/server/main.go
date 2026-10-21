package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/grpcwire"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8443", "listen address")
	state := flag.String("state", "state/server.json", "server state")
	cert := flag.String("cert", "state/server.crt", "TLS certificate")
	key := flag.String("key", "state/server.key", "TLS private key")
	flag.Parse()

	if _, err := os.Stat(*state); err != nil {
		fmt.Fprintf(os.Stderr, "server state not found: %s\nRun: go run ./cmd/setup\n", *state)
		os.Exit(2)
	}
	p := crypto.Init()
	srv, err := fdxt.LoadServerState(*state, p)
	if err != nil {
		fmt.Fprintln(os.Stderr, "load server state:", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(filepath.Dir(*cert), 0700); err != nil {
		panic(err)
	}
	g := &grpcwire.Server{FDXT: srv, StatePath: *state}
	if err := g.Serve(*addr, *cert, *key); err != nil {
		fmt.Fprintln(os.Stderr, "server:", err)
		os.Exit(1)
	}
}
