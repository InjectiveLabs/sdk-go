package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/credentials"
)

const (
	// nolint:gosec // not credentials, just the link to the public tokens list
	MainnetTokensListURL = "https://github.com/InjectiveLabs/injective-lists/raw/master/json/tokens/mainnet.json"
	// nolint:gosec // not credentials, just the link to the public tokens list
	TestnetTokensListURL = "https://github.com/InjectiveLabs/injective-lists/raw/master/json/tokens/testnet.json"
	// nolint:gosec // not credentials, just the link to the public tokens list
	DevnetTokensListURL = "https://github.com/InjectiveLabs/injective-lists/raw/master/json/tokens/devnet.json"
)

func cookieByName(cookies []*http.Cookie, key string) *http.Cookie {
	for _, c := range cookies {
		if c.Name == key {
			return c
		}
	}
	return nil
}

type Network struct {
	LcdEndpoint             string
	TmEndpoint              string
	ChainGrpcEndpoint       string
	ChainStreamGrpcEndpoint string
	ChainTLSCert            credentials.TransportCredentials
	ExchangeGrpcEndpoint    string
	ExplorerGrpcEndpoint    string
	ExchangeTLSCert         credentials.TransportCredentials
	ExplorerTLSCert         credentials.TransportCredentials
	ChainId                 string
	FeeDenom                string
	Name                    string
	ChainCookieAssistant    CookieAssistant
	ExchangeCookieAssistant CookieAssistant
	ExplorerCookieAssistant CookieAssistant
	OfficialTokensListURL   string
}

// nolint:cyclomatic // we leave the function as it is for clarity
// revive:disable:function-length // we leave the function as it is for clarity
func LoadNetwork(name, node string) Network {
	switch name {
	case "local":
		return Network{
			LcdEndpoint:             "http://localhost:10337",
			TmEndpoint:              "http://localhost:26657",
			ChainGrpcEndpoint:       "localhost:9900",
			ChainStreamGrpcEndpoint: "localhost:9999",
			ExchangeGrpcEndpoint:    "localhost:9910",
			ExplorerGrpcEndpoint:    "localhost:9911",
			ChainId:                 "injective-1",
			FeeDenom:                "inj",
			Name:                    "local",
			ChainCookieAssistant:    &DisabledCookieAssistant{},
			ExchangeCookieAssistant: &DisabledCookieAssistant{},
			ExplorerCookieAssistant: &DisabledCookieAssistant{},
			OfficialTokensListURL:   MainnetTokensListURL,
		}
	case "devnet-1":
		return Network{
			LcdEndpoint:             "https://devnet-1.lcd.injective.dev",
			TmEndpoint:              "https://devnet-1.tm.injective.dev:443",
			ChainGrpcEndpoint:       "devnet-1.grpc.injective.dev:9900",
			ChainStreamGrpcEndpoint: "devnet-1.grpc.injective.dev:9999",
			ExchangeGrpcEndpoint:    "devnet-1.api.injective.dev:9910",
			ExplorerGrpcEndpoint:    "devnet-1.api.injective.dev:9911",
			ChainId:                 "injective-777",
			FeeDenom:                "inj",
			Name:                    "devnet-1",
			ChainCookieAssistant:    &DisabledCookieAssistant{},
			ExchangeCookieAssistant: &DisabledCookieAssistant{},
			ExplorerCookieAssistant: &DisabledCookieAssistant{},
			OfficialTokensListURL:   DevnetTokensListURL,
		}
	case "devnet":
		return Network{
			LcdEndpoint:             "https://devnet.lcd.injective.dev",
			TmEndpoint:              "https://devnet.tm.injective.dev:443",
			ChainGrpcEndpoint:       "devnet.injective.dev:9900",
			ChainStreamGrpcEndpoint: "devnet.injective.dev:9999",
			ExchangeGrpcEndpoint:    "devnet.injective.dev:9910",
			ExplorerGrpcEndpoint:    "devnet.api.injective.dev:9911",
			ChainId:                 "injective-777",
			FeeDenom:                "inj",
			Name:                    "devnet",
			ChainCookieAssistant:    &DisabledCookieAssistant{},
			ExchangeCookieAssistant: &DisabledCookieAssistant{},
			ExplorerCookieAssistant: &DisabledCookieAssistant{},
			OfficialTokensListURL:   DevnetTokensListURL,
		}
	case "testnet":
		validNodes := []string{"lb", "sentry"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}

		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, chainStreamGrpcEndpoint, exchangeGrpcEndpoint, explorerGrpcEndpoint string
		var chainTLSCert, exchangeTLSCert, explorerTLSCert credentials.TransportCredentials
		var chainCookieAssistant, exchangeCookieAssistant, explorerCookieAssistant CookieAssistant
		if node == "lb" {
			lcdEndpoint = "https://testnet.sentry.lcd.injective.network:443"
			tmEndpoint = "https://testnet.sentry.tm.injective.network:443"
			chainGrpcEndpoint = "testnet.sentry.chain.grpc.injective.network:443"
			chainStreamGrpcEndpoint = "testnet.sentry.chain.stream.injective.network:443"
			exchangeGrpcEndpoint = "testnet.sentry.exchange.grpc.injective.network:443"
			explorerGrpcEndpoint = "testnet.sentry.explorer.grpc.injective.network:443"
			chainTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			exchangeTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			explorerTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			chainCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
			exchangeCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
			explorerCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
		} else if node == "sentry" {
			lcdEndpoint = "https://testnet.lcd.injective.network:443"
			tmEndpoint = "https://testnet.tm.injective.network:443"
			chainGrpcEndpoint = "testnet.chain.grpc.injective.network:443"
			chainStreamGrpcEndpoint = "testnet.chain.stream.injective.network:443"
			exchangeGrpcEndpoint = "testnet.exchange.grpc.injective.network:443"
			explorerGrpcEndpoint = "testnet.explorer.grpc.injective.network:443"
			chainTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			exchangeTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			explorerTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			chainCookieAssistant = &DisabledCookieAssistant{}
			exchangeCookieAssistant = &DisabledCookieAssistant{}
			explorerCookieAssistant = &DisabledCookieAssistant{}
		}

		return Network{
			LcdEndpoint:             lcdEndpoint,
			TmEndpoint:              tmEndpoint,
			ChainGrpcEndpoint:       chainGrpcEndpoint,
			ChainStreamGrpcEndpoint: chainStreamGrpcEndpoint,
			ChainTLSCert:            chainTLSCert,
			ExchangeGrpcEndpoint:    exchangeGrpcEndpoint,
			ExchangeTLSCert:         exchangeTLSCert,
			ExplorerGrpcEndpoint:    explorerGrpcEndpoint,
			ExplorerTLSCert:         explorerTLSCert,
			ChainId:                 "injective-888",
			FeeDenom:                "inj",
			Name:                    "testnet",
			ChainCookieAssistant:    chainCookieAssistant,
			ExchangeCookieAssistant: exchangeCookieAssistant,
			ExplorerCookieAssistant: explorerCookieAssistant,
			OfficialTokensListURL:   TestnetTokensListURL,
		}
	case "mainnet":
		validNodes := []string{"lb"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}
		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, chainStreamGrpcEndpoint, exchangeGrpcEndpoint, explorerGrpcEndpoint string
		var chainTLSCert, exchangeTLSCert, explorerTLSCert credentials.TransportCredentials
		var chainCookieAssistant, exchangeCookieAssistant, explorerCookieAssistant CookieAssistant

		lcdEndpoint = "https://sentry.lcd.injective.network"
		tmEndpoint = "https://sentry.tm.injective.network:443"
		chainGrpcEndpoint = "sentry.chain.grpc.injective.network:443"
		chainStreamGrpcEndpoint = "sentry.chain.stream.injective.network:443"
		exchangeGrpcEndpoint = "sentry.exchange.grpc.injective.network:443"
		explorerGrpcEndpoint = "sentry.explorer.grpc.injective.network:443"
		chainTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		exchangeTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		explorerTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		chainCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
		exchangeCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
		explorerCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}

		return Network{
			LcdEndpoint:             lcdEndpoint,
			TmEndpoint:              tmEndpoint,
			ChainGrpcEndpoint:       chainGrpcEndpoint,
			ChainStreamGrpcEndpoint: chainStreamGrpcEndpoint,
			ChainTLSCert:            chainTLSCert,
			ExchangeGrpcEndpoint:    exchangeGrpcEndpoint,
			ExchangeTLSCert:         exchangeTLSCert,
			ExplorerGrpcEndpoint:    explorerGrpcEndpoint,
			ExplorerTLSCert:         explorerTLSCert,
			ChainId:                 "injective-1",
			FeeDenom:                "inj",
			Name:                    "mainnet",
			ChainCookieAssistant:    chainCookieAssistant,
			ExchangeCookieAssistant: exchangeCookieAssistant,
			ExplorerCookieAssistant: explorerCookieAssistant,
			OfficialTokensListURL:   MainnetTokensListURL,
		}

	default:
		panic(fmt.Sprintf("invalid network %s", name))
	}
}

// NewNetwork returns a new Network instance with all cookie assistants disabled.
// It can be used to setup a custom environment from scratch.
func NewNetwork() Network {
	return Network{
		ChainCookieAssistant:    &DisabledCookieAssistant{},
		ExchangeCookieAssistant: &DisabledCookieAssistant{},
		ExplorerCookieAssistant: &DisabledCookieAssistant{},
		OfficialTokensListURL:   MainnetTokensListURL,
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DialerFunc(_ context.Context, addr string) (net.Conn, error) {
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
func ProtocolAndAddress(listenAddr string) (protocol, address string) {
	protocol, address = "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
