# Maintainer Guide

This document is intended for **SDK maintainers** who need to update the proto definitions, chain types, or generated data files when a new `injective-core` or `injective-indexer` release is published. SDK users only need [README.md](README.md).

---

## 1. Configuring the upstream versions (do this first)

Before running any update command, set the upstream version tags in [Makefile](Makefile). The two targets that clone the upstream repositories each carry a `-b <tag>` flag that determines which version all downstream steps will draw from:

```makefile
clone-injective-indexer:
	git clone https://github.com/InjectiveLabs/injective-indexer.git -b v1.19.0 --depth 1 --single-branch

clone-injective-core:
	git clone https://github.com/InjectiveLabs/injective-core.git -b v1.19.0 --depth 1 --single-branch
```

To update to a new release (e.g. `v1.20.0`), change both tags:

```makefile
clone-injective-indexer:
	git clone https://github.com/InjectiveLabs/injective-indexer.git -b v1.20.0 --depth 1 --single-branch

clone-injective-core:
	git clone https://github.com/InjectiveLabs/injective-core.git -b v1.20.0 --depth 1 --single-branch
```

Both versions should be kept in sync with the same chain release tag. These two lines are the single source of truth that drives every downstream step: exchange API bindings, chain module types, the `proto/` tree, and the auto-generated `injective_data/chain_messages_list.json`.

---

## 2. Updating Exchange API bindings

```bash
make copy-exchange-client
```

This target (which internally calls `clone-injective-indexer`):

1. Removes the entire `exchange/` directory.
2. Recreates per-service `pb/` subdirectories under `exchange/`.
3. Copies the already-generated `*.pb.go` files from `injective-indexer/api/gen/grpc/.../pb/` into the corresponding `exchange/...` paths.
4. Removes the cloned `injective-indexer` directory.

No local proto generation runs — the SDK consumes the indexer's pre-generated Go bindings directly.

The version cloned is determined by the `-b <tag>` flag in the `clone-injective-indexer` target (see [section 1](#1-configuring-the-upstream-versions-do-this-first)).

---

## 3. Updating chain types and proto

```bash
make copy-chain-types
```

This target (which internally calls `clone-injective-core`) performs the following steps:

1. Copies `codec`, `crypto/*`, and per-module `types/` directories from `injective-core/injective-chain/...` into the corresponding `chain/...` paths. For most modules only `*.pb.go` and `codec.go` are copied; for `exchange`, `evm`, `oracle`, `peggy`, and `wasmx`, additional hand-written `.go` source files are also included. Files matching `*test.go` and `*gw.go` are stripped.
2. Copies the `proto/` tree from `injective-core/proto` into `proto/`, replacing the previous contents entirely.
3. Rewrites all import paths in-place across every copied `.go` file so that `github.com/InjectiveLabs/injective-core/injective-chain[/modules]` becomes `github.com/InjectiveLabs/sdk-go/chain`. This rewrite uses BSD `sed -i ""` syntax, which works on macOS. **Linux maintainers must use GNU `sed -i` instead** — either edit the Makefile temporarily or run the step on macOS.
4. Removes the cloned `injective-core` directory.
5. Automatically invokes `make extract-message-names` to regenerate `injective_data/chain_messages_list.json` from the freshly copied types.

The version cloned is determined by the `-b <tag>` flag in the `clone-injective-core` target (see [section 1](#1-configuring-the-upstream-versions-do-this-first)).

### Syncing `go.mod`

`make copy-chain-types` does **not** update [go.mod](go.mod) automatically. After running it, manually align the SDK's `go.mod` dependencies with those declared in `injective-core`'s `go.mod` for the target tag. The reference file is available directly on GitHub:

```
https://github.com/InjectiveLabs/injective-core/blob/<tag>/go.mod
```

After applying the changes, run:

```bash
go mod tidy
```

Failing to do this after a version bump will typically produce compile errors or mismatched dependency versions.

---

## 4. Generated data files

### `injective_data/chain_messages_list.json`

This file is regenerated automatically at the end of `make copy-chain-types` via the `extract-message-names` target. It can also be regenerated independently:

```bash
make extract-message-names
```

The pipeline:
1. Searches every `tx.pb.go` and `msgs.pb.go` file under `chain/` for `proto.RegisterType(...)` calls.
2. Extracts the fully-qualified proto type name from each call (e.g. `injective.exchange.v1beta1.MsgCreateSpotMarketOrder`).
3. Drops any name ending in `Response`.
4. Deduplicates and sorts the result.
5. Writes the final JSON array to `injective_data/chain_messages_list.json` using `jq`.

The file is a reference list of all chain message type URLs. It is **not** loaded by any Go code in this repository — it ships as static data for downstream consumers and tooling that need a canonical list of injectable message types.

### `injective_data/ofac.json`

This file contains a snapshot of OFAC-sanctioned and restricted wallet addresses sourced from the [injective-lists](https://github.com/InjectiveLabs/injective-lists) repository. To refresh it:

```bash
make update-ofac-list
```

This target delegates to `examples/chain/ofac/1_DownloadOfacList/example.go`, which calls `chainclient.DownloadOfacList()` — the same function used at runtime in `NewOfacChecker()`. The download URL and output path are defined once in [`client/chain/ofac.go`](client/chain/ofac.go).

The committed copy in the repository is a snapshot used for offline use and initial setup. At runtime, `NewOfacChecker()` checks whether the file exists on disk and downloads a fresh copy automatically if it is missing — so `make update-ofac-list` is only needed when you want to explicitly update the committed snapshot.

---

## 5. Running tests and lint

After regenerating exchange bindings or chain types, run the test and lint suites to confirm nothing is broken.

### Tests

```bash
make tests
```

Runs `go test -race ./client/... ./ethereum/...` after clearing the test cache.

To collect a coverage profile:

```bash
make coverage
```

Writes `coverage.out` in atomic mode for the same packages.

### Lint

The lint targets all run `golangci-lint` with a 15-minute timeout. They differ only in which baseline revision they compare against, controlling how much of the codebase is checked:

- `make lint` — reports only findings new since the `dev` branch (default for PR work).
- `make lint-last-commit` — reports findings new since `HEAD~` (just the last commit).
- `make lint-master` — reports findings new since the `master` branch.
- `make lint-all` — runs the linter against the entire repository with no baseline (use for full audits or after bulk regenerations).

`golangci-lint` must be installed locally; all four targets export `GOPROXY=direct` before running.

---

## 6. Release update checklist

1. **Edit `Makefile`** — update the `-b <tag>` flag in both `clone-injective-indexer` and `clone-injective-core` to the new release tag (see [section 1](#1-configuring-the-upstream-versions-do-this-first)).
2. **Update Exchange bindings** — `make copy-exchange-client`
3. **Update chain types and proto** — `make copy-chain-types` (also regenerates `chain_messages_list.json`)
4. **Sync `go.mod`** — align dependency versions in [go.mod](go.mod) with `injective-core`'s `go.mod` for the target tag, then run `go mod tidy` (see [section 3](#syncing-gomod)).
5. **Verify** — `make tests && make lint`
6. **Refresh OFAC list** (optional) — `make update-ofac-list`
7. **Commit** all changed files, including `chain/`, `exchange/`, `proto/`, `go.mod`, `go.sum`, `injective_data/chain_messages_list.json`, and `Makefile`.
