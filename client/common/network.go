package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"path"
	"runtime"
	"strings"

	"google.golang.org/grpc/credentials"
)

type Network struct {
	LcdEndpoint          string
	TmEndpoint           string
	ChainGrpcEndpoint    string
	ChainTlsCert         credentials.TransportCredentials
	ExchangeGrpcEndpoint string
	ExchangeTlsCert      credentials.TransportCredentials
	ChainId              string
	Fee_denom            string
	Name                 string
}

func getFileAbsPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), relativePath)
}

func LoadNetwork(name string, node string) Network {
	switch name {

	case "devnet-1":
		return Network{
			LcdEndpoint:          "https://devnet-1.lcd.injective.dev",
			TmEndpoint:           "https://devnet-1.tm.injective.dev:443",
			ChainGrpcEndpoint:    "tcp://devnet-1.grpc.injective.dev:9900",
			ExchangeGrpcEndpoint: "tcp://devnet-1.api.injective.dev:9910",
			ChainId:              "injective-777",
			Fee_denom:            "inj",
			Name:                 "devnet-1",
		}
	case "devnet":
		return Network{
			LcdEndpoint:          "https://devnet.lcd.injective.dev",
			TmEndpoint:           "https://devnet.tm.injective.dev:443",
			ChainGrpcEndpoint:    "tcp://devnet.injective.dev:9900",
			ExchangeGrpcEndpoint: "tcp://devnet.injective.dev:9910",
			ChainId:              "injective-777",
			Fee_denom:            "inj",
			Name:                 "devnet",
		}
	case "testnet":
		validNodes := []string{"sentry0", "sentry1", "k8s"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}

		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, exchangeGrpcEndpoint string
		var chainTlsCert, exchangeTlsCert credentials.TransportCredentials
		if node == "k8s" {
			certPath := getFileAbsPath("../cert/testnet.crt")
			lcdEndpoint = "https://k8s.testnet.lcd.injective.network"
			tmEndpoint = "https://k8s.testnet.tm.injective.network:443"
			chainGrpcEndpoint = "tcp://k8s.testnet.chain.grpc.injective.network:443"
			chainTlsCert = LoadTlsCert(certPath, chainGrpcEndpoint)
			exchangeGrpcEndpoint = "tcp://k8s.testnet.exchange.grpc.injective.network:443"
			exchangeTlsCert = LoadTlsCert(certPath, exchangeGrpcEndpoint)
		} else {
			lcdEndpoint = fmt.Sprintf("http://%s.injective.dev:10337", node)
			tmEndpoint = fmt.Sprintf("http://%s.injective.dev:26657", node)
			chainGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.dev:9900", node)
			exchangeGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.dev:9910", node)
		}

		return Network{
			LcdEndpoint:          lcdEndpoint,
			TmEndpoint:           tmEndpoint,
			ChainGrpcEndpoint:    chainGrpcEndpoint,
			ChainTlsCert:         chainTlsCert,
			ExchangeGrpcEndpoint: exchangeGrpcEndpoint,
			ExchangeTlsCert:      exchangeTlsCert,
			ChainId:              "injective-888",
			Fee_denom:            "inj",
			Name:                 "testnet",
		}
	case "mainnet":
		validNodes := []string{"k8s", "lb", "sentry0", "sentry1", "sentry2", "sentry3"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}
		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, exchangeGrpcEndpoint string
		var chainTlsCert, exchangeTlsCert credentials.TransportCredentials
		if node == "k8s" {
			certPath := getFileAbsPath("../cert/mainnet.crt")
			lcdEndpoint = fmt.Sprintf("https://%s.mainnet.lcd.injective.network", node)
			tmEndpoint = fmt.Sprintf("https://%s.mainnet.tm.injective.network:443", node)
			chainGrpcEndpoint = fmt.Sprintf("tcp://%s.mainnet.chain.grpc.injective.network:443", node)
			chainTlsCert = LoadTlsCert(certPath, chainGrpcEndpoint)
			exchangeGrpcEndpoint = fmt.Sprintf("tcp://%s.mainnet.exchange.grpc.injective.network:443", node)
			exchangeTlsCert = LoadTlsCert(certPath, exchangeGrpcEndpoint)
		} else if node == "lb" {
			lcdEndpoint = "https://k8s.global.mainnet.lcd.injective.network"
			tmEndpoint = "https://k8s.global.mainnet.tm.injective.network"
			chainGrpcEndpoint = "k8s.global.mainnet.chain.grpc.injective.network:443"
			exchangeGrpcEndpoint = "k8s.global.mainnet.exchange.grpc.injective.network:443"
			chainTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			exchangeTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		} else {
			lcdEndpoint = fmt.Sprintf("http://%s.injective.network:10337", node)
			tmEndpoint = fmt.Sprintf("http://%s.injective.network:26657", node)
			chainGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.network:9900", node)
			exchangeGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.network:9910", node)
		}

		return Network{
			LcdEndpoint:          lcdEndpoint,
			TmEndpoint:           tmEndpoint,
			ChainGrpcEndpoint:    chainGrpcEndpoint,
			ChainTlsCert:         chainTlsCert,
			ExchangeGrpcEndpoint: exchangeGrpcEndpoint,
			ExchangeTlsCert:      exchangeTlsCert,
			ChainId:              "injective-1",
			Fee_denom:            "inj",
			Name:                 "mainnet",
		}
	}

	return Network{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DialerFunc(ctx context.Context, addr string) (net.Conn, error) {
	return Connect(addr)
}

// Connect dials the given address and returns a net.Conn. The protoAddr argument should be prefixed with the protocol,
// eg. "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock"
func Connect(protoAddr string) (net.Conn, error) {
	proto, address := ProtocolAndAddress(protoAddr)
	conn, err := net.Dial(proto, address)
	return conn, err
}

// ProtocolAndAddress splits an address into the protocol and address components.
// For instance, "tcp://127.0.0.1:8080" will be split into "tcp" and "127.0.0.1:8080".
// If the address has no protocol prefix, the default is "tcp".
func ProtocolAndAddress(listenAddr string) (string, string) {
	protocol, address := "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
