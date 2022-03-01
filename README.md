# Injective Protocol Golang SDK 🌟

---

## 📚 Getting Started

Clone the repository locally and install needed dependencies

```bash
$ git clone git@github.com:InjectiveLabs/sdk-go.git
$ cd sdk-go
$ go install ./...
```

## Run examples
```bash
# import pk into keyring if you use keyring
injectived keys unsafe-import-eth-key inj-user 5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e

# run chain example
go run examples/chain/0_MsgSend.go

# run exchange example
go run examples/exchange/derivative_exchange_rpc/0_GetOrderbook.go
```

---

## Updating Exchange API proto and client

```bash
$ make copy-exchange-client
```

(you have to clone [this repo](https://github.com/InjectiveLabs/injective-exchange) into `../injective-exchange`)

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

- Website at <a href="https://injectiveprotocol.com" target="_blank">`injectiveprotocol.com`</a>
- Twitter at <a href="https://twitter.com/InjectiveLabs" target="_blank">`@InjectiveLabs`</a>

---

## 🔐 License
