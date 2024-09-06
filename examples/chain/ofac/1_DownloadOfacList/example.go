package main

import (
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
)

func main() {
	err := chainclient.DownloadOfacList()
	if err != nil {
		panic(err)
	}
}
