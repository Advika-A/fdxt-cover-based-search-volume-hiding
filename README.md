# FDXT Proof of Concept

A **minimal** Go implementation of the core **FDXT** conjunctive
searchable-encryption protocol from:

> Y. Li, P. Xu, B. Zhang, W. Wang, Y. Zheng, K. Liang, "FDXT: Forward and
> Backward Private Conjunctive Searchable Encryption to Suppress Volume
> Leakages Caused by Cross-Tags," IEEE TIFS, vol. 21, 2026.

This is a proof of concept for the corrected FDXT core plus a separated
client/server gRPC execution path. The core has validated `n = 1` and
`n = 2` queries and the implementation is written generically over the
number of conjunctive keywords. C1, experiment automation, and the dashboard
are intentionally not included yet. MongoDB remains in the Session 2 codebase
and is not required by the new local gRPC demo.

It uses **real** pairing-based cryptography via native
[PBC](https://crypto.stanford.edu/pbc/) and the
[`github.com/Nik-U/pbc`](https://github.com/Nik-U/pbc) cgo bindings — no
fake/toy group is substituted anywhere in the protocol code.

---

## 1. What's implemented

- **Setup** — client samples `sk = (KW, KT, KX, KY, KZ)`; one process-wide
  Type-A pairing and one fixed, deterministic G1 generator `g` are
  initialized exactly once (`internal/crypto/pairing.go`).
- **Update / Encrypt** (`internal/fdxt/client.go`) — for `(op, w, id)`:
  - `addr, val` — single-keyword ciphertext entry for `CDB_T`.
  - `L, T` — address/trapdoor for the cross-tag ciphertext in `CDB_X`.
  - `xtag = g^(Fp(KX,w) * Fp(KY,id||op))`
  - `X = H(xtag)`, `C = X XOR T` (stored at `CDB_X[L]`).
  - `alpha = Fp(KY,id||op) * Fp(KZ, w||c_max_w)^-1` (stored alongside `val`
    in `CDB_T[addr]`).
- **Trapdoor generation & Search** (`internal/fdxt/search.go`) — builds
  `TKL` (cross-tag reveal tokens, shuffled), and `STKL`/`XTKL`
  (`xtk_{j,k} = g^(Fp(KX,wk) * Fp(KZ,w1||j))`) for the least-frequently
  updated keyword `w1` versus the remaining keywords.
- **Server-side search** (`internal/fdxt/server.go`) — decrypts revealed
  cross-tags into the permanent `XSET`, then for each historical update to
  `w1` obliviously recomputes `xtk_{j,k}^alpha_j` and checks membership in
  `XSET` to test whether that entry also matches the other keyword(s).
- **Client-side result recovery** — decrypts `val` and returns only entries
  whose op is `add` and whose match count equals `n` (the number of
  queried keywords).

Supported: `n = 1` (plain single-keyword search) and `n = 2` (conjunctive
`AND`) — the code is written generically over `n`, but only these two are
exercised/tested per the request.

---

## 2. Known POC simplifications (documented, not hidden)

- PRFs `F`/`Fp` are instantiated with HMAC-SHA256 (a standard, secure PRF),
  not a bespoke construction.
- `(id, op)` plaintext is packed as `"<id>|<op>"`, zero-padded to 32 bytes
  before XOR-masking, and trailing zero bytes are trimmed on decrypt. This
  is safe for the ASCII demo identifiers used here but is not a general
  length-hiding encoding.
- The generator `g` is derived deterministically via hash-to-G1 from a
  fixed seed (`SetFromHash`) instead of an independently random sample —
  required by the "one fixed generator for the process lifetime" design
  constraint.
- Backward-privacy deletion bookkeeping (`op = del`) is implemented in the
  cross-tag/decrypt path but not separately exercised by the demo/tests.
- No persistence, no network transport — client/server interact via direct
  Go function calls in a single process, which is sufficient to validate
  the protocol's cryptographic logic.

---

## 3. Project layout

```
fdxt-poc/
  cmd/demo/main.go            # tiny in-memory core demo
  cmd/setup/main.go           # creates persistent client/server state from local demo data
  cmd/server/main.go          # standalone FDXT gRPC server
  cmd/client/main.go          # standalone FDXT gRPC client
  internal/crypto/
    pairing.go                 # single pairing + single fixed generator
    prf.go                     # F, Fp, H, Xor
    pairing_test.go
    prf_test.go
  internal/fdxt/
    types.go                   # Op, Counter
    util.go                    # byte-encoding helpers
    client.go                  # Setup + Update
    server.go                  # CDB_X, CDB_T, XSET + server-side Search
    search.go                  # client-side trapdoor generation + Search
    fdxt_test.go                # core correctness tests (the important one!)
  tests/
    integration_test.go        # top-level end-to-end test
  params/a.param                # reference copy of the fixed pairing params
  go.mod
```

---

## 4. Build & run on Kali Linux

### 4.1 Prerequisites

FDXT needs the **native PBC C library** (via cgo) plus GMP and Go.

```bash
sudo apt update
sudo apt install -y build-essential golang-go libgmp-dev flex bison
```

`libpbc` is not packaged on Kali, so build it from source:

```bash
cd /tmp
wget https://crypto.stanford.edu/pbc/files/pbc-0.5.14.tar.gz
tar xzf pbc-0.5.14.tar.gz
cd pbc-0.5.14
./configure
make -j"$(nproc)"
sudo make install
sudo ldconfig
```

(If `wget` is blocked in your environment, `libpbc` and its headers can be
copied in from any machine where you already built it — the POC only needs
`libpbc.so`/`libpbc.a` and `pbc/pbc.h` on the standard library/include
paths.)

### 4.2 Fetch the Go dependency

```bash
cd fdxt-poc
go mod tidy
```

This downloads `github.com/Nik-U/pbc` and generates `go.sum`.

### 4.3 Build

```bash
go build ./...
```

### 4.4 Test

```bash
go test ./... -v
```

The single most important test:

```bash
go test ./internal/fdxt/ -run TestSearchConjunctive_N2 -v
```

It asserts that `K1 AND K2` returns **exactly** the ID set `{ID1, ID4}`
(compared as a set, not just a count).

Other notable tests:
- `TestCrossTagAlgebra` — checks the raw algebraic identity
  `xtk_{j,k}^alpha == xtag(wk, id, op)` in isolation from the rest of the
  protocol.
- `TestSearchConjunctive_AfterPriorSingleKeywordSearch` — checks that the
  deferred cross-tag reveal (forward privacy) still produces the correct
  result when `K1` is searched alone first and `K1 AND K2` is issued later.
- `TestG1SerializationRoundTrip` / `TestZrSerializationRoundTrip` — G1/Zr
  element `Bytes()`/`SetBytes()` round-trip correctly.

### 4.5 Run the demo

```bash
go run ./cmd/demo
```

Expected output:

```
Setup complete. Adding documents...

Search K1:
[ID1 ID2 ID4]

Search K1 AND K2:
[ID1 ID4]
```

---

## 5. Validation performed without a live libpbc install

During development, every Go source file in this repo was:

1. Parsed and formatted with `gofmt` (no syntax errors).
2. Compiled and `go vet`-checked against a byte-for-byte API-compatible
   (but non-cryptographic) stand-in for `github.com/Nik-U/pbc`, built by
   diffing against the real module's exported function signatures
   (`Params.NewPairing`, `Pairing.NewG1/NewZr`, `Element.Mul/Invert/PowZn/
   SetFromHash/Bytes/SetBytes`, etc.) — confirming this code calls the real
   API correctly.
3. Exercised end-to-end against that stand-in: every test that does **not**
   depend on genuine bilinear-group algebra (PRF determinism, XOR/mask
   correctness, address/val encoding, single-keyword search, empty
   intersection) passed. Every test that **does** depend on real pairing
   algebra (the cross-tag matching itself) correctly failed against the
   fake, non-algebraic stand-in — which is expected, since a random hash
   function is not a group homomorphism. This is exactly the failure
   pattern you want to see: it isolates "is the Go code wired correctly"
   from "does the real pairing math work", and only the latter requires
   your local `libpbc` install to confirm.

With the real `github.com/Nik-U/pbc` + real `libpbc` (as installed via
Section 4.1), `TestCrossTagAlgebra` and the conjunctive-search tests are
expected to pass, because the relationship

```
xtk_{j,k}^alpha = (g^(Fp(KX,wk) * Fp(KZ,w1||j))) ^ (Fp(KY,id||op) * Fp(KZ,w1||j)^-1)
                 = g^(Fp(KX,wk) * Fp(KY,id||op))
                 = xtag(wk, id, op)
```

holds identically over the real group, by basic exponent algebra.

**Update (confirmed):** this was subsequently run for real on Kali Linux
with real `libpbc` installed, and produced `[ID1 ID4]` for `K1 AND K2` as
expected -- the cross-tag algebra above checks out against a real pairing,
not just the fake stub described above.

---

# Session 2: real MongoDB dataset (Crime_USENIX_REV)

> **Schema fix applied (v2).** Two earlier guesses at the schema were
> both wrong: first a collection named `documents` with an `id_keywords`
> array field, then a collection `id_keywords` shaped as
> `{"id": "<doc_id>", "val_st": ["<kw1>", "<kw2>", ...]}` (one document
> per dataset record, "id" + a keyword array). Neither matches the real
> BSON.
>
> The **confirmed, actual** schema, inspected directly against the live
> `Crime_USENIX_REV.id_keywords` collection, is the inverse of that
> second guess -- one document **per keyword**, holding that keyword's
> full document-ID set:
>
> ```json
> {"_id": ObjectId("5ec19009d41c80e2c3ee16f7"), "k": "F0", "val_set": ["0", "1", "2", "3", "..."]}
> ```
>
> `k` is the keyword; `val_set` is the list of document IDs that keyword
> appears in. MongoDB's own `_id` is not the FDXT document identifier and
> is not used for anything. There is no separate reverse-index
> collection -- `id_keywords` already *is* the keyword -> document-ID
> index, so the earlier optional `id_keyword_search_list` / `keyword_ids`
> fast-path lookup collection and its fallback logic have been removed
> entirely as unnecessary now that the primary collection serves that
> role directly. See 6.1 and 6.1.1 below for the full schema and the
> pre-flight validation / vacuous-result handling.

Everything in Sections 1-5 above is Session 1 and is **unchanged** --
`internal/crypto/` and `internal/fdxt/{client,server,search,types,util}.go`
were not modified. Session 2 only *adds* a MongoDB layer on top, wired in
through one new package (`internal/indexer`) that calls the existing,
unmodified `fdxt.Client.Update` for every real `(keyword, id)` pair read
from MongoDB -- the same call the Session 1 demo made for the 5-document
toy dataset, just now driven by a real dataset instead of a literal.

## 6. What's new

```
fdxt-poc/
  cmd/cli/main.go                        # query CLI
  internal/db/
    mongo.go                              # MongoDB connect/stream/ground-truth
    decode.go                             # pure id_keywords {k, val_set} decoding logic
    decode_test.go                        # unit tests (no MongoDB needed)
    mongo_integration_test.go             # integration tests (env-gated)
  internal/indexer/indexer.go            # id_keywords records -> fdxt.Client.Update
  internal/compare/
    compare.go                            # exact ID-set comparison
    compare_test.go                       # unit tests
  internal/grpcproto/
    fdxt.proto                 # protobuf service/message schema
    fdxt.pb.go                 # protobuf wire encoder/decoder used by the transport
    fdxt_test.go
  internal/grpcwire/
    client.go                  # gRPC-over-HTTP/2 client
    server.go                  # gRPC-over-HTTP/2 server
  internal/queryparser/
    queryparser.go                        # "F1227 AND F1438" -> ["F1227","F1438"]
    queryparser_test.go                   # unit tests
  tests/
    integration_test.go                   # Session 1, unchanged
    mongo_pipeline_integration_test.go    # full Mongo->FDXT->ground-truth pipeline
```

### 6.1 Dataset schema (the real FDXT schema, `Crime_USENIX_REV`)

This section describes the **real** dataset schema, as directly inspected
against the live BSON -- not a generic/simplified stand-in. There is
exactly **one** relevant collection:

**`id_keywords`** -- one document per keyword, each holding that
keyword's full document-ID set:

```json
{
  "_id": ObjectId("5ec19009d41c80e2c3ee16f7"),
  "k": "F0",
  "val_set": ["0", "1", "2", "3", "..."]
}
```

- **Keyword**: read from the `k` string field. There is no fallback to
  `_id` -- `_id` is MongoDB housekeeping, not the FDXT identifier, and
  `decodeRecord` never touches it.
- **Document IDs**: read from the `val_set` array field. Each element is
  stringified.

This means `id_keywords` already *is* the keyword -> document-ID index.
The earlier design's optional `id_keyword_search_list` / `keyword_ids`
reverse-lookup collection (and the fast-path/fallback logic around it)
has been removed: it was solving a problem (making per-keyword lookups
fast) that the real schema doesn't have, since every `KeywordGroundTruth`
call is already a single `$in` query against `id_keywords` followed by an
in-process set intersection -- no scan, no second collection.

Documents missing a non-empty `k` string field or a usable non-empty
`val_set` array are silently skipped during indexing (not counted, not
fabricated), so a few malformed rows in a real-world dump won't crash the
pipeline.

If your actual collection uses different field names, `internal/db/decode.go`
(the `keywordField` / `docIDsField` constants, plus `extractKeyword` /
`extractDocIDs`) and `internal/db/mongo.go` (`StreamRecords`,
`KeywordGroundTruth`, `SampleCoOccurringKeywordPair`) are the only places
that know the schema -- everything else (FDXT core, CLI, comparison
logic) is schema-agnostic.

### 6.1.1 Pre-flight validation (no more trivial "success" on empty data)

Before indexing anything, the CLI calls `Store.ValidateSchema`, which:

- Confirms the configured `-collection` (default `id_keywords`) actually
  exists in the target database, and errors out immediately if it does
  not (listing the collections that *do* exist, to make a typo obvious).
- Confirms it contains at least one document, and refuses to run
  otherwise -- an empty collection would make FDXT's search and the
  ground-truth query both trivially return nothing, which would then
  trivially "match": that proves nothing about the protocol, so the CLI
  treats it as a hard error rather than a silent pass.
- Samples one document and runs it through the same `decodeRecord` used
  for real indexing; if it doesn't have a usable `k` + `val_set` shape,
  `ValidateSchema` fails immediately with a clear schema-mismatch error
  instead of silently indexing zero pairs and only finding out later.

Separately, even for a single query where the collection *is* populated,
if that specific query happens to return zero results from **both** FDXT
and the ground-truth query, the CLI does not print `Exact match: true`.
It prints an explicit `INCONCLUSIVE` line instead and exits non-zero,
because an empty-vs-empty comparison is indistinguishable from "the
keywords were never indexed at all" and does not validate anything. This
is tracked via the new `compare.Result.Vacuous` field; `compare.Compare`'s
own set-comparison logic is otherwise unchanged (its unit tests, which
predate this fix, still pass unmodified).

### 6.2 The data flow (exactly as specified)

```
query "F1227 AND F1438"
        |
        v
queryparser.Parse            -> ["F1227", "F1438"]
        |
        v
indexer.BuildFromMongo       -> streams every id_keywords document (one per
                                 keyword), calls the UNCHANGED
                                 fdxt.Client.Update(server, OpAdd, keyword, id)
                                 for every id in that keyword's val_set
        |
        v
client.Search(server, ["F1227","F1438"])   <- the real FDXT protocol, n=2
        |
        v
FDXT returned ID set
        |
        v
store.KeywordGroundTruth(["F1227","F1438"])  <- independent MongoDB oracle:
                                                 a single $in query against
                                                 id_keywords for both
                                                 keywords, then an in-process
                                                 set intersection of their
                                                 val_set arrays -- does NOT
                                                 touch the FDXT index
        |
        v
compare.Compare(fdxtIDs, groundTruthIDs)     <- EXACT set comparison:
                                                 false positives, false
                                                 negatives, exact-match flag
```

### 6.3 Build & run the CLI on Kali

Same prerequisites as Section 4.1 (Go + libpbc), plus a running MongoDB
with the real `Crime_USENIX_REV.id_keywords` collection already loaded.

```bash
cd fdxt-poc
go mod tidy      # also fetches the mongo-driver dependency
go build ./...
```

Run a real query (all flags shown here are also the defaults, so this is
identical to just `go run ./cmd/cli query "F1227 AND F1438"`):

```bash
go run ./cmd/cli \
  -uri "mongodb://localhost:27017" \
  -db Crime_USENIX_REV \
  -collection id_keywords \
  query "F1227 AND F1438"
```

If your collection is named differently, pass the real name via
`-collection`. See `go run ./cmd/cli -h` for the full flag list.

Expected output shape (exact IDs/counts depend on your data; per the
sanity check in the task, `F1227 AND F1438` should come out to 109):

```
Connected to MongoDB: Crime_USENIX_REV.id_keywords
Estimated keyword documents in collection: <N>
Indexing dataset into the FDXT encrypted index (this may take a while for large collections)...
Indexed <N> keywords / <M> (keyword,id) pairs.

Query: F1227 AND F1438
FDXT IDs: [... sorted list ...]
Ground-truth IDs: [... sorted list ...]
Returned count: 109
Ground-truth count: 109
Exact match: true
False positives: []
False negatives: []
```

The CLI exits with status `0` only when `Exact match: true` on a
non-vacuous result; it exits `1` on a mismatch, on a vacuous (empty vs.
empty) result, on connection/query errors, or if pre-flight validation
fails (missing/empty/wrongly-shaped `-collection`); a malformed query
string exits `2`. This makes it scriptable for later automation without a
false "success".

`-limit N` caps how many keyword documents get streamed into the FDXT
index (0 = the whole collection, the default). Indexing the full
dataset single-threaded will take some time; this is expected at this
stage of the project (performance work is explicitly deferred, per the
task priorities).

### 6.4 Tests

**Always run, no MongoDB required** (pure logic):

```bash
go test ./internal/compare/... ./internal/queryparser/... ./internal/db/... -run 'TestDecodeRecord|TestCompare|TestParse' -v
```

**Session 1 core, unchanged, still passing** (needs real libpbc):

```bash
go test ./internal/crypto/... ./internal/fdxt/... -v
```

**MongoDB integration tests** -- skip cleanly (not fail) unless you point
them at a real instance:

```bash
export FDXT_MONGO_URI="mongodb://localhost:27017"
export FDXT_MONGO_DB="Crime_USENIX_REV"     # optional, this is the default
export FDXT_MONGO_COLLECTION="id_keywords"  # optional, this is the default

go test ./internal/db/... -run TestIntegration -v
```

This includes `TestIntegration_ValidateSchema_RealCollectionIsUsable` /
`_RejectsMissingCollection` (the pre-flight check from 6.1.1) and the
`KeywordGroundTruth` subset/non-empty sanity checks, run against
keyword pairs sampled live from the collection (nothing dataset-specific
is hardcoded).

**Full pipeline integration test** (MongoDB -> FDXT -> exact ground-truth
comparison, for both n=1 and n=2, using a keyword pair sampled live from
the collection -- nothing dataset-specific is hardcoded):

```bash
export FDXT_MONGO_URI="mongodb://localhost:27017"
export FDXT_MONGO_INDEX_LIMIT=50000   # optional: cap indexing for a faster test run

go test ./tests/ -run TestMongoBackedConjunctiveQuery_ExactMatch -v
```

Run everything at once (integration tests just skip if `FDXT_MONGO_URI`
isn't set):

```bash
go test ./... -v
```

### 6.5 On the "known sanity checks"

The task lists known reference intersection counts (e.g.
`F1227 AND F1438 = 109`). Per the task instructions, **none of these
values are hardcoded anywhere in this codebase** -- not in
`KeywordGroundTruth`, not in the CLI, not in any test. `KeywordGroundTruth`
computes its answer live via a MongoDB `$in` query (over `id_keywords`,
one document per queried keyword) plus an in-process set intersection,
every time it's called, and the integration tests sample their own
keyword pairs from whatever real keyword documents happen to be in the
collection rather than referencing specific keyword codes. If you run
`go run ./cmd/cli query "F1227 AND F1438"` against the real dataset, the
printed `Ground-truth count` and `Returned count` should both come out to
109 on their own -- that's the thing being tested, not something this code
assumes going in.

### 6.6 What I could and couldn't verify myself (Session 2 fix)

This sandbox has neither a running `Crime_USENIX_REV` MongoDB instance
nor real `libpbc` (installing `libpbc-dev` isn't possible -- it isn't in
the Ubuntu package repos, and the upstream PBC source host isn't in this
environment's network allowlist; `github.com/Nik-U/pbc`'s cgo build fails
at `#include <pbc/pbc.h>`), so I could not run the full
MongoDB -> FDXT -> ground-truth pipeline end-to-end myself, and could not
confirm the `F1227 AND F1438 = 109` sanity check directly. That matches
the same limitation already noted for Session 1's crypto core. What I
*did* verify directly, without fabricating results:

- Installed Go 1.22 in the sandbox (`apt-get install golang-go`) and, with
  `GOPROXY=direct` (`proxy.golang.org` isn't reachable here either, so
  modules are fetched straight from `github.com`/`codeload.github.com`,
  which are allowlisted), successfully resolved and downloaded every
  dependency, including `github.com/Nik-U/pbc` and
  `go.mongodb.org/mongo-driver`.
- `go build ./...`, `go vet ./...`, and `go test -v ./...` all run clean
  for `internal/db`, `internal/compare`, and `internal/queryparser` --
  the three packages that carry this session's fix and have no
  dependency on `github.com/Nik-U/pbc`. This is not stubbed: it is the
  real `decodeRecord`, `StreamRecords`, `KeywordGroundTruth`,
  `ValidateSchema`, and `SampleCoOccurringKeywordPair` logic running for
  real against literal and `primitive.A`-shaped BSON documents matching
  the confirmed real schema, including a regression test
  (`TestDecodeRecord_OldSchemaFieldsAreNotUsed`) that asserts the
  previous, incorrect `{id, val_st}` shape is now rejected rather than
  silently accepted.
- `gofmt -l` is clean on every file touched by this fix, and
  `internal/indexer/indexer.go` and `cmd/cli/main.go` (which do depend on
  the unbuildable `internal/fdxt` package) were syntax-checked with
  `gofmt -e` and manually cross-checked line-by-line against the real,
  unmodified `fdxt.Client.Update(srv *Server, op Op, w string, id string)`
  and `fdxt.Client.Search(srv *Server, q []string) []string` signatures
  in `internal/fdxt/client.go` / `search.go` -- the call sites match
  exactly, but this is not the same as an actual `go build` pass on those
  two files, since that requires `libpbc`.
- The MongoDB integration tests and the full pipeline test are written to
  skip cleanly (reporting `SKIP`, not `PASS`) when no MongoDB is
  reachable, so they never fabricate a passing result; no live MongoDB
  was reachable here either, so they did skip.

What remains genuinely unverified until you run it on your own machine
(with real `libpbc` and the real dataset loaded): the actual `go build`
of `internal/fdxt`/`internal/crypto`/`internal/indexer`/`cmd/cli`; the
live MongoDB integration and full-pipeline tests; and the
`F1227 AND F1438 = 109` sanity check itself. If any of those surface an
issue, the error output is designed to point at which layer (Mongo schema
vs. FDXT core vs. CLI wiring) it's in -- in particular, `ValidateSchema`
now fails loudly and specifically on a schema mismatch instead of
quietly indexing zero pairs.

---

# Session 3.5: deterministic 50-document local dataset

The local demo/integration dataset is now a deterministic **50-document**
cybersecurity-oriented dataset. This session changes only the local dataset;
MongoDB code, the corrected FDXT cryptographic core, and the Session 3
client/server gRPC architecture remain unchanged.

## Dataset design

- Exactly 50 documents: `ID1` through `ID50`.
- 30 distinct English/technical keywords.
- Keywords are normal terms such as `computer`, `security`, `network`,
  `malware`, `database`, `cloud`, `encryption`, `phishing`, `firewall`,
  `authentication`, `software`, and `application`.
- The keyword-document distribution is deliberately non-uniform:
  high-frequency keywords occur in about 20-24 documents, medium-frequency
  keywords in about 8-16 documents, and low-frequency keywords in 3-7
  documents.
- Documents have varied keyword counts rather than identical keyword lists.
- The dataset is fixed and deterministic; repeated runs produce the same
  IDs, memberships, frequencies, intersections, and ground truth.

`internal/demodata/demodata.go` is the single source of truth. It provides:

```text
GetAllDocuments()
GetKeywords(documentID)
GetDocumentsForKeyword(keyword)
GetDocumentByID(documentID)
GroundTruthConjunctiveQuery(keywords)
KeywordFrequencies()
Stats()
```

The existing `Records()` function inverts the same data into the keyword ->
ID representation used by the existing optional MongoDB layer. MongoDB code
itself is not modified in this session.

## Dataset statistics

Run:

```bash
go run ./cmd/demo --dataset-stats
```

This prints:

- document count;
- distinct keyword count;
- total keyword-ID pairs;
- every keyword's actual document frequency;
- the deterministic 15-query experiment suite and its runtime-derived
  ground-truth counts.

The current dataset contains 50 documents, 30 distinct keywords, and 269
keyword-ID pairs.

## Deterministic query suite

The query suite covers large, medium, small, and zero-result n=2 queries,
plus n=3 queries. Examples include:

```text
security AND network
access AND data
computer AND software
malware AND phishing
cloud AND database
malware AND vulnerability
firewall AND api
computer AND security AND network
computer AND software AND application
cloud AND data AND database
forensics AND malware AND threat
```

Expected IDs are **not** hardcoded into the search implementation. Ground
truth is calculated directly from the authoritative dataset.

## Local FDXT demo

Run:

```bash
go run ./cmd/demo
```

This indexes all 50 documents into the real in-memory FDXT baseline and runs
the deterministic query suite through the existing:

```text
FDXT Update
    -> PrepareSearch
    -> Server.ProcessSearch
    -> RecoverSearch
    -> independent dataset ground truth
    -> exact-set comparison
```

The small 5-document fixture remains in the FDXT unit tests for focused
cryptographic correctness testing.

## Separate client/server test

The Session 3 gRPC path also uses the 50-document dataset.

Fresh state:

```bash
rm -rf state
go run ./cmd/setup --dataset 50
```

Terminal 1:

```bash
go run ./cmd/server
```

Terminal 2:

```bash
go run ./cmd/client --query "computer AND security" --dataset 50
```

The client generates the real FDXT search material, sends it over the
existing gRPC/HTTP2 path, recovers the result, and compares it against the
independently calculated 50-document ground truth.

Additional representative queries:

```bash
go run ./cmd/client --query "security AND network" --dataset 50
go run ./cmd/client --query "computer AND software AND application" --dataset 50
go run ./cmd/client --query "firewall AND api" --dataset 50
```

The zero-result query is useful for checking that the CLI reports an
inconclusive empty-vs-empty comparison rather than treating it as proof of
correctness.

## Tests added for the 50-document dataset

```bash
go test ./internal/demodata -v
go test ./tests -run TestDataset50 -v
```

The tests validate the dataset shape, deterministic records, frequency tiers,
keyword lookups, n=2/n=3 ground truth, and the corrected FDXT result against
the independent dataset ground truth.

The 50-document baseline is the frozen foundation for the C1 + experiment session described below.

---

# Session 3: separated client/server + gRPC

Session 3 adds a real separated network path around the existing FDXT
implementation. The client generates the actual FDXT search material; the
server executes the actual `Server.ProcessSearch`; the client performs result
recovery and independently computes local ground truth.

The transport uses the standard gRPC wire format over HTTP/2 with protobuf
messages:
- HTTP/2
- `application/grpc`
- 5-byte gRPC message framing
- protobuf wire encoding
- `grpc-status` trailers

The client and server reuse the same fixed PBC parameters and fixed G1
generator. The client state and server encrypted index are persisted
separately under `state/`.

## Setup

From the repository root:

```bash
go run ./cmd/setup
```

This creates:

```text
state/client.json
state/server.json
```

The client state contains the FDXT secret keys and local counters. The server
state contains the encrypted FDXT database and XSET. These files are local POC
state and should not be committed to a public repository.

## Start the server

Terminal 1:

```bash
go run ./cmd/server
```

The server creates a local self-signed certificate on first start:

```text
state/server.crt
state/server.key
```

It listens on:

```text
https://127.0.0.1:8443/fdxt.FDXT/Search
```

## Run the client

Terminal 2:

```bash
go run ./cmd/client --query "computer AND security"
```

The client:
1. parses the query;
2. builds the real FDXT trapdoor/tokens;
3. sends them over the gRPC wire path;
4. receives the server's actual FDXT response;
5. performs client-side recovery;
6. independently calculates the local ground truth;
7. compares exact ID sets.

Example:

```bash
go run ./cmd/client --query "computer AND security"
```

Also test a single keyword:

```bash
go run ./cmd/client --query "computer"
```

If the existing generic FDXT path supports a three-keyword query in the current
dataset, test it through the same client/server path:

```bash
go run ./cmd/client --query "computer AND software AND application"
```

Normal query execution now dispatches automatically by query size: n=2 uses C1; n=1 and n>=3 use corrected FDXT. The user never selects a variant.

## Important state behavior

The client and server must be initialized from the same `cmd/setup` run.
Do not independently generate client/server keys or pairing parameters.

The server persists XSET after each search, and the client persists its search
counters after each successful query. This allows repeated searches to use the
same FDXT state across separate processes.

For a fresh experiment, remove the local `state/` directory and run:

```bash
rm -rf state
go run ./cmd/setup
```

Then start the server again.

## Session 3 verification

Run:

```bash
go test ./...
go build ./...
```

Then verify the two-process path:

```text
Terminal 1: go run ./cmd/server
Terminal 2: go run ./cmd/client --query "computer AND security"
```

The final client output must show:
- recovered FDXT IDs;
- independent ground-truth IDs;
- returned count;
- ground-truth count;
- exact match;
- false positives;
- false negatives;
- gRPC round-trip timing.

The server never computes ground truth and never receives the dataset's expected
answer as part of the search request.


Session 3 state compatibility note: server CDB_X/CDB_T keys and XSET values are arbitrary binary bytes and are base64-encoded in JSON. If upgrading from an earlier Session 3 build, delete `state/` and rerun `go run ./cmd/setup`.


---

# Session 4+5: C1 + experiment framework

This session adds the experimental C1 extension and a reproducible comparison
framework while leaving the MongoDB implementation untouched.

## C1 scope

C1 is deliberately limited to **two-keyword conjunctive queries**. The normal
query path automatically dispatches:

```text
n = 1       -> corrected FDXT
n = 2       -> corrected FDXT + C1
n >= 3      -> corrected FDXT
```

There is no user-facing `--variant=...` switch. The experiment runner can
internally execute both variants so the research comparison remains fair.

## C1 mechanism

C1 is an adaptive **protocol-shaped search-cover layer** for two-keyword
conjunctive queries. It does not append arbitrary values to the final server
response. Instead, the client creates fresh synthetic FDXT-shaped cover
updates and appends their **TKL entries and search rows to the ordinary
FDXT trapdoor**.

For each cover entry, the client constructs the same algebraic objects used
by a real update: fresh `L,T,C`, a synthetic `CDB_T` `(val, alpha)` entry,
and an `XTK` row. The corresponding temporary server-side `CDB_X`/`CDB_T`
state is supplied only for that search. The server then processes genuine and
cover TKL/rows through the ordinary FDXT search path.

The synthetic recovered identifier uses a reserved client-only prefix. After
normal FDXT recovery, the client discards those synthetic records, so the
logical document-ID result is unchanged. The cover entries are not replayed
genuine records and are not stored permanently in the server database.

The adaptive policy targets the next power-of-two response volume, subject to
the configured maximum number of cover entries. Thus, for a genuine candidate
volume `V`, C1 creates `target(V) - V` protocol-shaped cover entries.

This is an experimental volume-cover mechanism, **not a formal security proof**.
The evaluation measures the resulting observable response volume and the
overhead introduced by the additional trapdoor/search material.

### C1 configuration

The normal client accepts:

```bash
go run ./cmd/client --query "computer AND security" --c1-cover-budget 32
```

The budget controls the maximum number of protocol-shaped cover entries
added to the C1 search material. Setting it to `0` disables the additional cover material while retaining
the same automatic n=2 dispatch.

For n=3, C1 is not attached even if a cover budget is supplied.

## Correctness validation

Every query is compared against the independent deterministic 50-document
local ground truth. The comparison reports:

- result count;
- ground-truth count;
- exact ID-set match;
- false positives;
- false negatives.

C1 must not change the logical result. The same query therefore has the same
real IDs under corrected FDXT and corrected FDXT+C1.

## Experiment runner

The exhaustive experiment can be run with one command:

```bash
go run ./cmd/experiment --runs=10 --warmups=1 --c1-cover-budget=32
```

The runner internally compares:

```text
corrected_fdxt
corrected_fdxt_c1
```

for the same deterministic workloads. The normal client never asks the user
to choose between them.

The runner defaults to the exhaustive two-keyword evaluation. It evaluates every
unique unordered pair of keywords in the dataset. The representative workload
mode remains available with `--all-pairs=false` and includes:

- high-frequency AND high-frequency;
- high-frequency AND medium-frequency;
- medium-frequency AND medium-frequency;
- medium-frequency AND low-frequency;
- low-frequency AND low-frequency;
- zero-result n=2;
- n=1 correctness regression;
- n=3 correctness regression.

The exhaustive two-keyword mode runs every unique unordered pair of keywords in the dataset. The current
deterministic dataset has 30 keywords, so this produces C(30,2) = 435 unique
n=2 queries. Each pair is evaluated under both corrected FDXT and corrected
FDXT+C1 for the configured number of measured runs. No pair is randomly
selected or replaced by a representative category in this mode.

Example (explicitly enabling the already-default all-pairs mode):

```bash
go run ./cmd/experiment --all-pairs --runs=10 --warmups=1 --c1-cover-budget=32
```

This produces 435 x 2 x 10 = 8,700 measured records (before any warm-up
records, which are not written to the measured output files). n=1/n=3
regression queries are omitted in all-pairs mode because the mode is
explicitly an exhaustive n=2 pair evaluation.

### Metrics

The runner records actual measurements for:

- setup time;
- update/encryption time;
- trapdoor generation time;
- server search time;
- recovery time;
- total query time;
- request bytes;
- response bytes;
- total communication bytes;
- serialized server-state storage size;
- logical result count;
- independent ground-truth count;
- exact match / FP / FN;
- genuine FDXT response volume;
- observed response volume;
- C1 cover volume;
- response-volume exposure proxy;
- C1 cover overhead.

Communication bytes are calculated from the same protobuf representation and
5-byte gRPC framing used by the real transport, so the experiment does not
invent a communication estimate.


### Response-volume leakage analysis

The exhaustive experiment also performs a separate volume-only leakage evaluation.
For each unique n=2 query, the hidden variable is the genuine FDXT candidate
volume and the observable variable is the encrypted response volume. The simulated
attacker is given only the observable volume and attempts to infer the genuine
volume. Repeated timing runs are collapsed to one logical point per query. A
deterministic 5-fold cross-validation procedure prevents a query from being used
to train its own inference rule.

The leakage report contains:

- **Inference accuracy** — exact equality between predicted and genuine volume; lower is better for privacy.
- **Mean absolute error (MAE)** — average absolute prediction error; higher means poorer inference.
- **RMSE** — square-root mean squared prediction error; higher means poorer inference.
- **Pearson correlation** — linear association between genuine and observed volume; lower means less direct volume tracking.
- **Mutual information (MI)** — information in bits shared by genuine and observed volume; lower means less volume information is exposed.
- **Normalized mutual information (NMI)** — MI normalized by the entropy of the two variables; lower means less normalized dependence.
- **Collision rate** — fraction of unordered query pairs with different genuine volumes that map to the same observed volume; higher means more volume indistinguishability.

These metrics quantify leakage under the explicitly defined response-volume
observation model. They are **experimental evidence, not a formal cryptographic
security proof**. In particular, they do not claim that C1 hides all leakage or
that the mechanism is secure against an observer with additional side channels.

The results are written to:

```text
results/leakage_analysis.csv
```

and are also included as **Table 6 — Response-volume leakage analysis** in
`results/summary_tables.md`.

### Outputs

The command creates:

```text
results/
  raw_runs.csv
  raw_runs.json
  summary.csv
  correctness.csv
  workloads.csv
  summary_tables.md
  search_latency_ms.svg
  communication_bytes.svg
  observed_volume.svg
  cover_volume.svg
  residual_volume_exposure.svg
  c1_overhead.svg
```

`summary_tables.md` contains the generated research tables. The SVG plots are
dependency-free and are generated directly by the Go experiment runner.

The `residual_volume_exposure` value is explicitly documented as a
response-volume cover proxy: genuine FDXT candidate volume divided by the
observed response volume. It should not be presented as a formal leakage or
security proof.

## Reproducibility

The dataset remains deterministic. Each measured run creates a fresh FDXT
client/server index so previous query counters do not contaminate the next
run. The baseline and C1 variants use the same workload, dataset, PBC
parameters, generator, and run count; only the C1 protocol-shaped search-cover mechanism
differs.

Before trusting a result set, run:

```bash
go test ./...
go build ./...
```

Then run the experiment command. The runner exits non-zero if any measured
record fails exact correctness.

MongoDB is intentionally untouched in this session. Dashboard generation is
also intentionally deferred; the dashboard can consume the generated CSV/JSON
files in the later session.

Leakage outputs include `leakage_points.csv`, `leakage_volume_mapping.svg`, and `leakage_metrics.svg`, with bootstrap 95% confidence intervals in Table 6.
