package common

import (
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc/metadata"
)

func TestMainnetKubernetesLoadBalancedCookieAssistant(t *testing.T) {
	assistant := MainnetKubernetesCookieAssistant()
	expectedCookie := fmt.Sprintf("GCLB=CMOO2-DdvKWMqQE; path=/; HttpOnly; expires=%s", time.Now().Add(5*time.Hour).Format("Mon, 02-Jan-2006 15:04:05 MST"))

	localMetadata := metadata.Pairs("set-cookie", expectedCookie)

	assistant.ProcessResponseMetadata(localMetadata)
	assistantMetadata := assistant.RealMetadata()
	cookieInfo := assistantMetadata.Get("cookie")

	if len(cookieInfo) > 0 {
		cookie := cookieInfo[0]
		if cookie != expectedCookie {
			t.Fatalf("The parsed cookie is different than the expected cookie")
		}
	} else {
		t.Fatalf("The cookie was not parsed")
	}
}

func TestMainnetKubernetesLoadBalancedCookieAssistantRemovesExpiredCookie(t *testing.T) {
	assistant := MainnetKubernetesCookieAssistant()
	tt := time.Now()
	closeExpirationTime := tt.Add(-30 * time.Second)

	expiredCookie := fmt.Sprintf(
		"GCLB=CMOO2-DdvKWMqQE; path=/; HttpOnly; expires=%s",
		closeExpirationTime.Format("Mon, 02-Jan-2006 15:04:05 MST"),
	)

	localMetadata := metadata.Pairs("set-cookie", expiredCookie)

	assistant.ProcessResponseMetadata(localMetadata)
	cookieInfo := assistant.RealMetadata()

	if len(cookieInfo) > 0 {
		t.Fatalf("The expired cookie was not removed")
	}

}

func TestTestnetKubernetesLoadBalancedCookieAssistant(t *testing.T) {
	assistant := TestnetKubernetesCookieAssistant()
	expectedCookie := fmt.Sprintf(
		"grpc-cookie=d97c7a00bcb7bc8b69b26fb0303b60d4; Expires=%s; Max-Age=172800; Path=/; Secure; HttpOnly",
		time.Now().Add(5*time.Hour).Format("Mon, 02-Jan-06 15:04:05 MST"),
	)

	localMetadata := metadata.Pairs("set-cookie", expectedCookie)

	assistant.ProcessResponseMetadata(localMetadata)
	assistantMetadata := assistant.RealMetadata()
	cookieInfo := assistantMetadata.Get("cookie")

	if len(cookieInfo) > 0 {
		cookie := cookieInfo[0]
		if cookie != expectedCookie {
			t.Fatalf("The parsed cookie is different than the expected cookie")
		}
	} else {
		t.Fatalf("The cookie was not parsed")
	}
}

func TestTestnetKubernetesLoadBalancedCookieAssistantRemovesExpiredCookie(t *testing.T) {
	assistant := TestnetKubernetesCookieAssistant()
	tt := time.Now()
	closeExpirationTime := tt.Add(-30 * time.Second)

	expiredCookie := fmt.Sprintf(
		"grpc-cookie=d97c7a00bcb7bc8b69b26fb0303b60d4; Expires=%s; Max-Age=172800; Path=/; Secure; HttpOnly",
		closeExpirationTime.Format("Mon, 02-Jan-06 15:04:05 MST"),
	)

	localMetadata := metadata.Pairs("set-cookie", expiredCookie)

	assistant.ProcessResponseMetadata(localMetadata)
	cookieInfo := assistant.RealMetadata()

	if len(cookieInfo) > 0 {
		t.Fatalf("The expired cookie was not removed")
	}

}

func TestBareMetalLoadBalancedCookieAssistant(t *testing.T) {
	assistant := BareMetalLoadBalancedCookieAssistant{}
	expectedCookie := "lb=706c9476f31159d415afe3e9172972618b8ab99455c2daa799199540e6d31a1a; Path=/"

	localMetadata := metadata.Pairs("set-cookie", expectedCookie)

	assistant.ProcessResponseMetadata(localMetadata)
	assistantMetadata := assistant.RealMetadata()
	cookieInfo := assistantMetadata.Get("cookie")

	if len(cookieInfo) > 0 {
		cookie := cookieInfo[0]
		if cookie != expectedCookie {
			t.Fatalf("The parsed cookie is different than the expected cookie")
		}
	} else {
		t.Fatalf("The cookie was not parsed")
	}
}

func TestDisabledCookieAssistant(t *testing.T) {
	assistant := DisabledCookieAssistant{}
	providedCookie := "lb=706c9476f31159d415afe3e9172972618b8ab99455c2daa799199540e6d31a1a; Path=/"

	localMetadata := metadata.Pairs("set-cookie", providedCookie)

	assistant.ProcessResponseMetadata(localMetadata)
	assistantMetadata := assistant.RealMetadata()
	cookieInfo := assistantMetadata.Get("cookie")

	if len(cookieInfo) > 0 {
		t.Fatalf("The disabled cookie assistant should not return any cookie")
	}
}
