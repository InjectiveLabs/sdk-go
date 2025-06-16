# Injective Protocol Golang SDK 🌟

[![codecov](https://codecov.io/gh/InjectiveLabs/sdk-go/graph/badge.svg?token=XDGZV265EE)](https://codecov.io/gh/InjectiveLabs/sdk-go)

---

## 📚 Getting Started

Clone the repository locally and install needed dependencies

```bash
$ git clone git@github.com:InjectiveLabs/sdk-go.git
$ cd sdk-go
$ go mod download
```

## Run examples
```bash
# import pk into keyring if you use keyring
injectived keys unsafe-import-eth-key inj-user 5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e

# run chain example
go run examples/chain/bank/1_MsgSend/example.go

# run exchange example
go run examples/exchange/derivatives/4_Orderbook/example.go
```

---

## Choose Exchange V1 or Exchange V2 queries

The SDK provides two different clients for interacting with the Injective Exchange:

- `ChainClient`: Use this client if you need to interact with Exchange V1. This client maintains compatibility with the original exchange implementation and is suitable for existing applications that haven't migrated to V2 yet. Note that this client will not include any new endpoints added to the Exchange module - for access to new features, you should migrate to V2.

- `ChainClientV2`: Use this client for all new applications or when you need to interact with Exchange V2 features. This client provides access to the latest exchange functionality and improvements, including all new endpoints added to the Exchange module.

Example usage:
```go
// For Exchange V1
client := chainclient.NewChainClient(...)

// For Exchange V2
clientV2 := chainclient.NewChainClientV2(...)
```

### Markets Assistant

The SDK provides a Markets Assistant to help you interact with markets in both V1 and V2. Here's how to create instances for each version:

```go
// For Exchange V1 markets
marketsAssistant, err := chain.NewMarketsAssistant(ctx, client)  // ChainClient instance
if err != nil {
    // Handle error
}

// For Exchange V2 markets
marketsAssistantV2, err := chain.NewHumanReadableMarketsAssistant(ctx, clientV2)  // ChainClientV2 instance
if err != nil {
    // Handle error
}
```

The Markets Assistant provides helper methods to:
- Fetch market information
- Get market prices
- Query orderbooks
- Access market statistics

Make sure to use the correct version of the Markets Assistant that matches your ChainClient version to ensure compatibility. The V1 assistant (`NewMarketsAssistant`) will only work with V1 markets, while the V2 assistant (`NewHumanReadableMarketsAssistant`) provides access to V2 markets and their features.

### Format Differences

There are important format differences between V1 and V2 endpoints:

- **Exchange V1**: All values (amounts, prices, margins, notionals) are returned in chain format (raw numbers)
- **Exchange V2**: Most values are returned in human-readable format for better usability:
  - Amounts, prices, margins, and notionals are in human-readable format
  - Deposit-related information remains in chain format to maintain consistency with the Bank module

This format difference is one of the key improvements in V2, making it easier to work with market data without manual conversion.

## Updating Exchange API proto and client

```bash
$ make copy-exchange-client
```

(you have to clone [this repo](https://github.com/InjectiveLabs/injective-indexer) into `../injective-indexer`)

---

## Publishing Tagged Release

```bash
$ git add .
$ git commit -m "bugfix"
$ git tag -a v1.1.1
$ git push origin master --tags
```

---

## ⛑ Support

Reach out to us at one of the following places!

- Website at <a href="https://injective.com" target="_blank">`injective.com`</a>
- Twitter at <a href="https://twitter.com/InjectiveLabs" target="_blank">`@InjectiveLabs`</a>

---

## License

Copyright © 2020 - 2025 Injective Labs Inc. (https://injectivelabs.org/)

<a href="https://drive.google.com/uc?export=view&id=1-fPQRh_D_dnun2yTtSsPW5MypVBOVYJP"><img src="https://drive.google.com/uc?export=view&id=1-fPQRh_D_dnun2yTtSsPW5MypVBOVYJP" style="width: 300px; max-width: 100%; height: auto" />

Originally released by Injective Labs Inc. under: <br />
Apache License <br />
Version 2.0, January 2004 <br />
http://www.apache.org/licenses/
