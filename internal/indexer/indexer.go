// Package indexer is the only bridge between the MongoDB dataset layer
// (internal/db) and the FDXT cryptographic core (internal/fdxt,
// internal/crypto). It does not implement any cryptography itself: it
// just calls the existing, unmodified fdxt.Client.Update for every real
// (keyword, id) pair read from MongoDB, exactly the way the Session 1
// demo called it for the 5-document toy dataset.
package indexer

import (
	"context"
	crand "crypto/rand"
	"math/big"

	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/db"
	"fdxt-poc/internal/fdxt"
)

// Stats summarizes what was loaded into the in-memory FDXT index.
type Stats struct {
	Keywords int64 // number of id_keywords documents (one per keyword) successfully decoded
	Pairs    int64 // number of (keyword, id) Update calls issued
}

// BuildFromMongo streams every keyword record from store (see
// db.Store.StreamRecords) and, for every document ID in that keyword's
// val_set, calls the real FDXT Update protocol:
// client.Update(server, fdxt.OpAdd, keyword, id).
//
// limit caps the number of MongoDB documents (i.e. keywords) read (0 =
// the whole collection). This is the only place dataset size is bounded;
// the FDXT cryptographic core itself has no notion of a dataset limit.
//
// If progress is non-nil and progressEvery > 0, progress is invoked with
// the running pair count every time it reaches a multiple of
// progressEvery (e.g. progressEvery=10 -> progress(10), progress(20),
// ...), so a caller (the CLI) can print "Indexed N pairs..." for both the
// small demo dataset and the full real dataset. Pass progress=nil (or
// progressEvery<=0) to disable this entirely -- existing callers that
// don't want progress output can do so without changing behavior
// otherwise.
func BuildFromMongo(ctx context.Context, store *db.Store, p *crypto.Ctx, limit int64, progressEvery int64, progress func(pairs int64)) (*fdxt.Client, *fdxt.Server, Stats, error) {
	client := fdxt.NewClient(p)
	server := fdxt.NewServer(p)

	var stats Stats
	var records []db.Record
	kws, err := store.StreamRecords(ctx, limit, func(r db.Record) error {
		records = append(records, r)
		return nil
	})
	if err != nil {
		return nil, nil, stats, err
	}
	stats.Keywords = kws

	keywords := make([]string, 0, len(records))
	for _, r := range records {
		keywords = append(keywords, r.Keyword)
		stats.Pairs += int64(len(r.DocIDs))
	}

	// Build the opaque C1 pool before any server update is installed. The
	// entries are then mixed randomly with genuine update events so the server
	// sees one undifferentiated encrypted update stream rather than a trailing
	// block that would reveal which entries are cover.
	pool, err := client.BuildC1Pool(keywords, 64)
	if err != nil {
		return nil, nil, stats, err
	}
	type event struct {
		cover       *fdxt.C1PoolEntry
		keyword, id string
	}
	events := make([]event, 0, int(stats.Pairs)+len(pool))
	pairsSeen := int64(0)
	for _, r := range records {
		for _, id := range r.DocIDs {
			events = append(events, event{keyword: r.Keyword, id: id})
			pairsSeen++
			if progress != nil && progressEvery > 0 && pairsSeen%progressEvery == 0 {
				progress(pairsSeen)
			}
		}
	}
	for i := range pool {
		ce := pool[i]
		events = append(events, event{cover: &ce})
	}
	for i := len(events) - 1; i > 0; i-- {
		n, err := crand.Int(crand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return nil, nil, stats, err
		}
		j := int(n.Int64())
		events[i], events[j] = events[j], events[i]
	}
	for _, e := range events {
		if e.cover != nil {
			if err := client.InstallC1Entry(server, *e.cover); err != nil {
				return nil, nil, stats, err
			}
		} else {
			client.Update(server, fdxt.OpAdd, e.keyword, e.id)
		}
	}
	client.C1Pool = pool

	return client, server, stats, nil
}
