package common

import (
	"errors"
	"net/http"
	"time"

	"google.golang.org/grpc/metadata"
)

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
		return errors.New("error getting a new cookie from the server")
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
		return errors.New("error getting a new cookie from the server")
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

func (assistant *DisabledCookieAssistant) Metadata(_ MetadataProvider) (string, error) {
	return "", nil
}

func (*DisabledCookieAssistant) RealMetadata() metadata.MD {
	return metadata.Pairs()
}

func (*DisabledCookieAssistant) ProcessResponseMetadata(header metadata.MD) {}
