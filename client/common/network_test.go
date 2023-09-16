package common

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func TestMainnetKubernetesLoadBalancedCookieAssistant(t *testing.T) {
	assistant := MainnetKubernetesCookieAssistant()
	expectedCookie := "GCLB=CMOO2-DdvKWMqQE; path=/; HttpOnly; expires=Sat, 16-Sep-2023 18:26:00 GMT"

	providerFunc := func() metadata.MD {
		md := metadata.Pairs("set-cookie", expectedCookie)
		return md
	}

	provider := NewMetadataProvider(providerFunc)
	cookie, err := assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != expectedCookie {
		t.Fatalf("The parsed cookie is different than the expected cookie")
	}
}

func TestMainnetKubernetesLoadBalancedCookieAssistantRemovesExpiredCookie(t *testing.T) {
	assistant := MainnetKubernetesCookieAssistant()
	tt := time.Now()
	closeExpirationTime := tt.Add(30 * time.Second)

	soonToExpireCookie := fmt.Sprintf(
		"GCLB=CMOO2-DdvKWMqQE; path=/; HttpOnly; expires=%s",
		closeExpirationTime.Format("Mon, 02-Jan-2006 15:04:05 MST"),
	)

	providerFunc := func() metadata.MD {
		md := metadata.Pairs("set-cookie", soonToExpireCookie)
		return md
	}

	provider := NewMetadataProvider(providerFunc)
	cookie, err := assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != soonToExpireCookie {
		t.Fatalf("The parsed cookie is different than the expected cookie")
	}

	nextExpirationTime := tt.Add(5 * time.Minute)
	secondCookie := fmt.Sprintf(
		"GCLB=CMOO2-DdvKWMqQE; path=/; HttpOnly; expires=%s",
		nextExpirationTime.Format("Mon, 02-Jan-2006 15:04:05 MST"),
	)

	providerFunc = func() metadata.MD {
		md := metadata.Pairs("set-cookie", secondCookie)
		return md
	}

	provider = NewMetadataProvider(providerFunc)
	cookie, err = assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != secondCookie {
		t.Fatalf("The expired cookie was not removed")
	}

	farFutureTime := tt.Add(5 * time.Hour)
	notRequiredCookie := fmt.Sprintf(
		"GCLB=CMOO2-DdvKWMqQE; path=/; HttpOnly; expires=%s",
		farFutureTime.Format("Mon, 02-Jan-2006 15:04:05 MST"),
	)

	providerFunc = func() metadata.MD {
		md := metadata.Pairs("set-cookie", notRequiredCookie)
		return md
	}

	provider = NewMetadataProvider(providerFunc)
	cookie, err = assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != secondCookie {
		t.Fatalf("The cookie assistant removed a cookie that was not expired")
	}

}

func TestTestnetKubernetesLoadBalancedCookieAssistant(t *testing.T) {
	assistant := TestnetKubernetesCookieAssistant()
	expectedCookie := "grpc-cookie=d97c7a00bcb7bc8b69b26fb0303b60d4; Expires=Sun, 17-Sep-23 13:18:08 GMT; Max-Age=172800; Path=/; Secure; HttpOnly"

	providerFunc := func() metadata.MD {
		md := metadata.Pairs("set-cookie", expectedCookie)
		return md
	}

	provider := NewMetadataProvider(providerFunc)
	cookie, err := assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != expectedCookie {
		t.Fatalf("The parsed cookie is different than the expected cookie")
	}
}

func TestTestnetKubernetesLoadBalancedCookieAssistantRemovesExpiredCookie(t *testing.T) {
	assistant := TestnetKubernetesCookieAssistant()
	tt := time.Now()
	closeExpirationTime := tt.Add(30 * time.Second)

	soonToExpireCookie := fmt.Sprintf(
		"grpc-cookie=d97c7a00bcb7bc8b69b26fb0303b60d4; Expires=%s; Max-Age=172800; Path=/; Secure; HttpOnly",
		closeExpirationTime.Format("Mon, 02-Jan-06 15:04:05 MST"),
	)

	providerFunc := func() metadata.MD {
		md := metadata.Pairs("set-cookie", soonToExpireCookie)
		return md
	}

	provider := NewMetadataProvider(providerFunc)
	cookie, err := assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != soonToExpireCookie {
		t.Fatalf("The parsed cookie is different than the expected cookie")
	}

	nextExpirationTime := tt.Add(5 * time.Minute)
	secondCookie := fmt.Sprintf(
		"grpc-cookie=d97c7a00bcb7bc8b69b26fb0303b60d4; Expires=%s; Max-Age=172800; Path=/; Secure; HttpOnly",
		nextExpirationTime.Format("Mon, 02-Jan-06 15:04:05 MST"),
	)

	providerFunc = func() metadata.MD {
		md := metadata.Pairs("set-cookie", secondCookie)
		return md
	}

	provider = NewMetadataProvider(providerFunc)
	cookie, err = assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != secondCookie {
		t.Fatalf("The expired cookie was not removed")
	}

	farFutureTime := tt.Add(5 * time.Hour)
	notRequiredCookie := fmt.Sprintf(
		"grpc-cookie=d97c7a00bcb7bc8b69b26fb0303b60d4; Expires=%s; Max-Age=172800; Path=/; Secure; HttpOnly",
		farFutureTime.Format("Mon, 02-Jan-06 15:04:05 MST"),
	)

	providerFunc = func() metadata.MD {
		md := metadata.Pairs("set-cookie", notRequiredCookie)
		return md
	}

	provider = NewMetadataProvider(providerFunc)
	cookie, err = assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != secondCookie {
		t.Fatalf("The cookie assistant removed a cookie that was not expired")
	}

}

func TestBareMetalLoadBalancedCookieAssistant(t *testing.T) {
	assistant := BareMetalLoadBalancedCookieAssistant{}
	expectedCookie := "lb=706c9476f31159d415afe3e9172972618b8ab99455c2daa799199540e6d31a1a; Path=/"

	providerFunc := func() metadata.MD {
		md := metadata.Pairs("set-cookie", expectedCookie)
		return md
	}

	provider := NewMetadataProvider(providerFunc)
	cookie, err := assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != expectedCookie {
		t.Fatalf("The parsed cookie is different than the expected cookie")
	}
}

func TestDisabledCookieAssistant(t *testing.T) {
	assistant := DisabledCookieAssistant{}
	providedCookie := "lb=706c9476f31159d415afe3e9172972618b8ab99455c2daa799199540e6d31a1a; Path=/"
	providerFunc := func() metadata.MD {
		md := metadata.Pairs("set-cookie", providedCookie)
		return md
	}

	provider := NewMetadataProvider(providerFunc)
	cookie, err := assistant.Metadata(provider)

	if err != nil {
		t.Errorf("Error parsing the cookie string %v", err)
	}

	if cookie != "" {
		t.Fatalf("The parsed cookie is different than the expected cookie")
	}
}
