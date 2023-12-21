package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"path"
	"runtime"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials"
)

const (
	SessionRenewalOffset = 2 * time.Minute
)

func cookieByName(cookies []*http.Cookie, key string) *http.Cookie {
	for _, c := range cookies {
		if c.Name == key {
			return c
		}
	}
	return nil
}

type MetadataProvider struct {
	f func() metadata.MD
}

func NewMetadataProvider(f func() metadata.MD) MetadataProvider {
	return MetadataProvider{f: f}
}

func (provider *MetadataProvider) metadata() metadata.MD {
	return provider.f()
}

type CookieAssistant interface {
	Metadata(provider MetadataProvider) (string, error)
}

type ExpiringCookieAssistant struct {
	expirationKey string
	timeLayout    string
	cookie        string
}

func (assistant *ExpiringCookieAssistant) initializeCookie(provider MetadataProvider) error {
	md := provider.metadata()
	cookieInfo := md.Get("set-cookie")

	if len(cookieInfo) == 0 {
		return fmt.Errorf("error getting a new cookie from the server")
	}

	assistant.cookie = cookieInfo[0]
	return nil
}

func (assistant *ExpiringCookieAssistant) checkCookieExpiration() {
	//borrow http request to parse cookie
	header := http.Header{}
	header.Add("Cookie", assistant.cookie)
	request := http.Request{Header: header}
	cookies := request.Cookies()
	cookie := cookieByName(cookies, assistant.expirationKey)

	if cookie != nil {
		expirationTime, err := time.Parse(assistant.timeLayout, cookie.Value)

		if err == nil {
			timestampDiff := time.Until(expirationTime)
			if timestampDiff < SessionRenewalOffset {
				assistant.cookie = ""
			}
		}
	}
}

func (assistant *ExpiringCookieAssistant) Metadata(provider MetadataProvider) (string, error) {
	if assistant.cookie == "" {
		err := assistant.initializeCookie(provider)
		if err != nil {
			return "", err
		}
	}

	cookie := assistant.cookie
	assistant.checkCookieExpiration()

	return cookie, nil
}

func TestnetKubernetesCookieAssistant() ExpiringCookieAssistant {
	assistant := ExpiringCookieAssistant{}
	assistant.expirationKey = "Expires"
	assistant.timeLayout = "Mon, 02-Jan-06 15:04:05 MST"

	return assistant
}

func MainnetKubernetesCookieAssistant() ExpiringCookieAssistant {
	assistant := ExpiringCookieAssistant{}
	assistant.expirationKey = "expires"
	assistant.timeLayout = "Mon, 02-Jan-2006 15:04:05 MST"

	return assistant
}

type BareMetalLoadBalancedCookieAssistant struct {
	cookie string
}

func (assistant *BareMetalLoadBalancedCookieAssistant) initializeCookie(provider MetadataProvider) error {
	md := provider.metadata()
	cookieInfo := md.Get("set-cookie")

	if len(cookieInfo) == 0 {
		return fmt.Errorf("error getting a new cookie from the server")
	}

	assistant.cookie = cookieInfo[0]
	return nil
}

func (assistant *BareMetalLoadBalancedCookieAssistant) Metadata(provider MetadataProvider) (string, error) {
	if assistant.cookie == "" {
		err := assistant.initializeCookie(provider)
		if err != nil {
			return "", err
		}
	}

	return assistant.cookie, nil
}

type DisabledCookieAssistant struct{}

func (assistant *DisabledCookieAssistant) Metadata(provider MetadataProvider) (string, error) {
	return "", nil
}

type Network struct {
	LcdEndpoint             string
	TmEndpoint              string
	ChainGrpcEndpoint       string
	ChainStreamGrpcEndpoint string
	ChainTlsCert            credentials.TransportCredentials
	ExchangeGrpcEndpoint    string
	ExplorerGrpcEndpoint    string
	ExchangeTlsCert         credentials.TransportCredentials
	ExplorerTlsCert         credentials.TransportCredentials
	ChainId                 string
	Fee_denom               string
	Name                    string
	chainCookieAssistant    CookieAssistant
	exchangeCookieAssistant CookieAssistant
	explorerCookieAssistant CookieAssistant
}

func (network *Network) ChainMetadata(provider MetadataProvider) (string, error) {
	return network.chainCookieAssistant.Metadata(provider)
}

func (network *Network) ExchangeMetadata(provider MetadataProvider) (string, error) {
	return network.exchangeCookieAssistant.Metadata(provider)
}

func (network *Network) ExplorerMetadata(provider MetadataProvider) (string, error) {
	return network.explorerCookieAssistant.Metadata(provider)
}

func getFileAbsPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), relativePath)
}

func LoadNetwork(name string, node string) Network {
	switch name {

	case "devnet-1":
		return Network{
			LcdEndpoint:             "https://devnet-1.lcd.injective.dev",
			TmEndpoint:              "https://devnet-1.tm.injective.dev:443",
			ChainGrpcEndpoint:       "tcp://devnet-1.grpc.injective.dev:9900",
			ChainStreamGrpcEndpoint: "tcp://devnet-1.grpc.injective.dev:9999",
			ExchangeGrpcEndpoint:    "tcp://devnet-1.api.injective.dev:9910",
			ExplorerGrpcEndpoint:    "tcp://devnet-1.api.injective.dev:9911",
			ChainId:                 "injective-777",
			Fee_denom:               "inj",
			Name:                    "devnet-1",
			chainCookieAssistant:    &DisabledCookieAssistant{},
			exchangeCookieAssistant: &DisabledCookieAssistant{},
			explorerCookieAssistant: &DisabledCookieAssistant{},
		}
	case "devnet":
		return Network{
			LcdEndpoint:             "https://devnet.lcd.injective.dev",
			TmEndpoint:              "https://devnet.tm.injective.dev:443",
			ChainGrpcEndpoint:       "tcp://devnet.injective.dev:9900",
			ChainStreamGrpcEndpoint: "tcp://devnet.injective.dev:9999",
			ExchangeGrpcEndpoint:    "tcp://devnet.injective.dev:9910",
			ExplorerGrpcEndpoint:    "tcp://devnet.api.injective.dev:9911",
			ChainId:                 "injective-777",
			Fee_denom:               "inj",
			Name:                    "devnet",
			chainCookieAssistant:    &DisabledCookieAssistant{},
			exchangeCookieAssistant: &DisabledCookieAssistant{},
			explorerCookieAssistant: &DisabledCookieAssistant{},
		}
	case "testnet":
		validNodes := []string{"lb", "sentry"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}

		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, chainStreamGrpcEndpoint, exchangeGrpcEndpoint, explorerGrpcEndpoint string
		var chainTlsCert, exchangeTlsCert, explorerTlsCert credentials.TransportCredentials
		var chainCookieAssistant, exchangeCookieAssistant, explorerCookieAssistant CookieAssistant
		if node == "lb" {
			lcdEndpoint = "https://testnet.sentry.lcd.injective.network:443"
			tmEndpoint = "https://testnet.sentry.tm.injective.network:443"
			chainGrpcEndpoint = "testnet.sentry.chain.grpc.injective.network:443"
			chainStreamGrpcEndpoint = "testnet.sentry.chain.stream.injective.network:443"
			exchangeGrpcEndpoint = "testnet.sentry.exchange.grpc.injective.network:443"
			explorerGrpcEndpoint = "testnet.sentry.explorer.grpc.injective.network:443"
			chainTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			exchangeTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			explorerTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
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
			chainTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			exchangeTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			explorerTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			chainCookieAssistant = &DisabledCookieAssistant{}
			exchangeCookieAssistant = &DisabledCookieAssistant{}
			explorerCookieAssistant = &DisabledCookieAssistant{}
		}

		return Network{
			LcdEndpoint:             lcdEndpoint,
			TmEndpoint:              tmEndpoint,
			ChainGrpcEndpoint:       chainGrpcEndpoint,
			ChainStreamGrpcEndpoint: chainStreamGrpcEndpoint,
			ChainTlsCert:            chainTlsCert,
			ExchangeGrpcEndpoint:    exchangeGrpcEndpoint,
			ExchangeTlsCert:         exchangeTlsCert,
			ExplorerGrpcEndpoint:    explorerGrpcEndpoint,
			ExplorerTlsCert:         explorerTlsCert,
			ChainId:                 "injective-888",
			Fee_denom:               "inj",
			Name:                    "testnet",
			chainCookieAssistant:    chainCookieAssistant,
			exchangeCookieAssistant: exchangeCookieAssistant,
			explorerCookieAssistant: explorerCookieAssistant,
		}
	case "mainnet":
		validNodes := []string{"lb"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}
		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, chainStreamGrpcEndpoint, exchangeGrpcEndpoint, explorerGrpcEndpoint string
		var chainTlsCert, exchangeTlsCert, explorerTlsCert credentials.TransportCredentials
		var chainCookieAssistant, exchangeCookieAssistant, explorerCookieAssistant CookieAssistant

		lcdEndpoint = "https://sentry.lcd.injective.network"
		tmEndpoint = "https://sentry.tm.injective.network:443"
		chainGrpcEndpoint = "sentry.chain.grpc.injective.network:443"
		chainStreamGrpcEndpoint = "sentry.chain.stream.injective.network:443"
		exchangeGrpcEndpoint = "sentry.exchange.grpc.injective.network:443"
		explorerGrpcEndpoint = "sentry.explorer.grpc.injective.network:443"
		chainTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		exchangeTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		explorerTlsCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
		chainCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
		exchangeCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
		explorerCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}

		return Network{
			LcdEndpoint:             lcdEndpoint,
			TmEndpoint:              tmEndpoint,
			ChainGrpcEndpoint:       chainGrpcEndpoint,
			ChainStreamGrpcEndpoint: chainStreamGrpcEndpoint,
			ChainTlsCert:            chainTlsCert,
			ExchangeGrpcEndpoint:    exchangeGrpcEndpoint,
			ExchangeTlsCert:         exchangeTlsCert,
			ExplorerGrpcEndpoint:    explorerGrpcEndpoint,
			ExplorerTlsCert:         explorerTlsCert,
			ChainId:                 "injective-1",
			Fee_denom:               "inj",
			Name:                    "mainnet",
			chainCookieAssistant:    chainCookieAssistant,
			exchangeCookieAssistant: exchangeCookieAssistant,
			explorerCookieAssistant: explorerCookieAssistant,
		}
	}

	return Network{}
}

// NewNetwork returns a new Network instance with all cookie assistants disabled.
// It can be used to setup a custom environment from scratch.
func NewNetwork() Network {
	return Network{
		chainCookieAssistant:    &DisabledCookieAssistant{},
		exchangeCookieAssistant: &DisabledCookieAssistant{},
		explorerCookieAssistant: &DisabledCookieAssistant{},
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
