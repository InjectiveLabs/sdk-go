package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials"
)

const (
	MainnetTokensListURL = "https://github.com/InjectiveLabs/injective-lists/raw/master/tokens/mainnet.json" // nolint:gosec // not credentials, just the link to the public tokens list
	TestnetTokensListURL = "https://github.com/InjectiveLabs/injective-lists/raw/master/tokens/testnet.json" // nolint:gosec // not credentials, just the link to the public tokens list
	DevnetTokensListURL  = "https://github.com/InjectiveLabs/injective-lists/raw/master/tokens/devnet.json"  // nolint:gosec // not credentials, just the link to the public tokens list
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
	RealMetadata() metadata.MD
	ProcessResponseMetadata(header metadata.MD)
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
	// borrow http request to parse cookie
	header := http.Header{}
	header.Add("Cookie", assistant.cookie)
	request := http.Request{Header: header}
	cookies := request.Cookies()
	cookie := cookieByName(cookies, assistant.expirationKey)

	if cookie != nil {
		expirationTime, err := time.Parse(assistant.timeLayout, cookie.Value)

		if err == nil {
			timestampDiff := time.Until(expirationTime)
			if timestampDiff < 0 {
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

func (assistant *ExpiringCookieAssistant) RealMetadata() metadata.MD {
	newMetadata := metadata.Pairs()
	assistant.checkCookieExpiration()
	if assistant.cookie != "" {
		newMetadata.Append("cookie", assistant.cookie)
	}
	return newMetadata
}

func (assistant *ExpiringCookieAssistant) ProcessResponseMetadata(header metadata.MD) {
	cookieInfo := header.Get("set-cookie")
	if len(cookieInfo) > 0 {
		assistant.cookie = cookieInfo[0]
	}
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

func (assistant *BareMetalLoadBalancedCookieAssistant) RealMetadata() metadata.MD {
	newMetadata := metadata.Pairs()
	if assistant.cookie != "" {
		newMetadata.Append("cookie", assistant.cookie)
	}
	return newMetadata
}

func (assistant *BareMetalLoadBalancedCookieAssistant) ProcessResponseMetadata(header metadata.MD) {
	cookieInfo := header.Get("set-cookie")
	if len(cookieInfo) > 0 {
		assistant.cookie = cookieInfo[0]
	}
}

type DisabledCookieAssistant struct{}

func (assistant *DisabledCookieAssistant) Metadata(provider MetadataProvider) (string, error) {
	return "", nil
}

func (assistant *DisabledCookieAssistant) RealMetadata() metadata.MD {
	return metadata.Pairs()
}

func (assistant *DisabledCookieAssistant) ProcessResponseMetadata(header metadata.MD) {}

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

func LoadNetwork(name, node string) Network {
	switch name {

	case "local":
		return Network{
			LcdEndpoint:             "http://localhost:10337",
			TmEndpoint:              "http://localhost:26657",
			ChainGrpcEndpoint:       "tcp://localhost:9900",
			ChainStreamGrpcEndpoint: "tcp://localhost:9999",
			ExchangeGrpcEndpoint:    "tcp://localhost:9910",
			ExplorerGrpcEndpoint:    "tcp://localhost:9911",
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
			ChainGrpcEndpoint:       "tcp://devnet-1.grpc.injective.dev:9900",
			ChainStreamGrpcEndpoint: "tcp://devnet-1.grpc.injective.dev:9999",
			ExchangeGrpcEndpoint:    "tcp://devnet-1.api.injective.dev:9910",
			ExplorerGrpcEndpoint:    "tcp://devnet-1.api.injective.dev:9911",
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
			ChainGrpcEndpoint:       "tcp://devnet.injective.dev:9900",
			ChainStreamGrpcEndpoint: "tcp://devnet.injective.dev:9999",
			ExchangeGrpcEndpoint:    "tcp://devnet.injective.dev:9910",
			ExplorerGrpcEndpoint:    "tcp://devnet.api.injective.dev:9911",
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
		validNodes := []string{"lb", "sentry", "sentry0", "sentry1", "sentry3"}
		if !contains(validNodes, node) {
			panic(fmt.Sprintf("invalid node %s for %s", node, name))
		}
		var lcdEndpoint, tmEndpoint, chainGrpcEndpoint, chainStreamGrpcEndpoint, exchangeGrpcEndpoint, explorerGrpcEndpoint string
		var chainTLSCert, exchangeTLSCert, explorerTLSCert credentials.TransportCredentials
		var chainCookieAssistant, exchangeCookieAssistant, explorerCookieAssistant CookieAssistant

		if node == "lb" || node == "sentry" {
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
		} else {
			lcdEndpoint = fmt.Sprintf("http://%s.injective.network:10337", node)
			tmEndpoint = fmt.Sprintf("http://%s.injective.network:26657", node)
			chainGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.network:9900", node)
			chainStreamGrpcEndpoint = "sentry.chain.stream.injective.network:443"
			exchangeGrpcEndpoint = fmt.Sprintf("tcp://%s.injective.network:9910", node)
			explorerGrpcEndpoint = "sentry.explorer.grpc.injective.network:443"
			chainTLSCert = insecure.NewCredentials()
			explorerTLSCert = insecure.NewCredentials()
			explorerTLSCert = credentials.NewServerTLSFromCert(&tls.Certificate{})
			chainCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
			exchangeCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
			explorerCookieAssistant = &BareMetalLoadBalancedCookieAssistant{}
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
func ProtocolAndAddress(listenAddr string) (protocol, address string) {
	protocol, address = "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
