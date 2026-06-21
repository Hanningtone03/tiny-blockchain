![CI](https://github.com/Hanningtone03/tiny-blockchain/actions/workflows/ci.yml/badge.svg)

# Tiny Blockchain

A blockchain built in Go; merkle trees, transaction validation, and peer-to-peer consensus across multiple nodes.

## How it works

Each block contains a set of transactions, a merkle root computed from them, a reference to the previous block's hash, and a nonce. Mining means searching for a nonce that produces a hash starting with enough leading zeros to satisfy the configured difficulty. When a node mines a block, it broadcasts its full chain to its peers, who adopt it if it's longer than their own; a simple longest-chain consensus rule.

Pairs well with [raft-consensus](https://github.com/Hanningtone03/raft-consensus); different mechanism, same underlying problem: how do independent nodes agree on a single shared truth.

## Project structure

```
internal/block/
├── block.go
└── merkle.go
internal/chain/
├── chain.go
└── validation.go
internal/wallet/
└── wallet.go
internal/network/
├── server.go
└── client.go
cmd/node/
└── main.go
```

## Running locally

Open three terminals:

```bash
go run cmd/node/main.go 7001 localhost:7002,localhost:7003
go run cmd/node/main.go 7002 localhost:7001,localhost:7003
go run cmd/node/main.go 7003 localhost:7001,localhost:7002
```

Submit a transaction and mine it:

```bash
curl -X POST http://localhost:7001/transaction -d '{"from":"alice","to":"bob","amount":50}'
curl -X POST http://localhost:7001/mine
```

Check that it propagated:

```bash
curl http://localhost:7002/chain
curl http://localhost:7003/chain
```

## Tech

- Go
- SHA-256 for hashing and merkle trees
- Plain HTTP for peer-to-peer communication
- No external dependencies
