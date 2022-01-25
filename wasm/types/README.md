## Referencing WASM types

Please don't include `github.com/CosmWasm/wasmd` as an sdk-go project dependency, those types needed on the client side should be included in this repo directly.

### Used types files:

```
wasm/types/ante.go
wasm/types/codec.go
wasm/types/errors.go
wasm/types/events.go
wasm/types/genesis.go
wasm/types/genesis.pb.go
wasm/types/ibc.pb.go
wasm/types/params.go
wasm/types/proposal.go
wasm/types/proposal.pb.go
wasm/types/query.pb.go
wasm/types/query.pb.gw.go
wasm/types/tx.go
wasm/types/tx.pb.go
wasm/types/types.go
wasm/types/types.pb.go
```
