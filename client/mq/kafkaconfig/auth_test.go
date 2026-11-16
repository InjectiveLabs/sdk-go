package kafkaconfig

import (
	"crypto/tls"
	"encoding/pem"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl"
)

func TestOptionsFromEnvRejectsPartialConfiguration(t *testing.T) {
	for mask := 0; mask < 7; mask++ {
		t.Run(string(rune('0'+mask)), func(t *testing.T) {
			for i, name := range []string{"MQ_KAFKA_USERNAME", "MQ_KAFKA_PASSWORD", "MQ_KAFKA_CA_FILE"} {
				value := ""
				if mask&(1<<i) != 0 {
					value = "configured"
				}
				t.Setenv(name, value)
			}
			opts, err := OptionsFromEnv()
			if mask == 0 {
				if err != nil || len(opts) != 0 {
					t.Fatalf("plaintext configuration changed: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatal("partial authentication configuration was accepted")
			}
		})
	}
}

func TestOptionsFromEnvRejectsInvalidCA(t *testing.T) {
	t.Setenv("MQ_KAFKA_USERNAME", "reader")
	t.Setenv("MQ_KAFKA_PASSWORD", "test-password")
	caFile := filepath.Join(t.TempDir(), "ca.pem")
	t.Setenv("MQ_KAFKA_CA_FILE", caFile)
	if _, err := OptionsFromEnv(); err == nil {
		t.Fatal("missing CA file was accepted")
	}
	if err := os.WriteFile(caFile, []byte("invalid PEM"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := OptionsFromEnv(); err == nil {
		t.Fatal("invalid CA file was accepted")
	}
}

func TestOptionsFromEnvVerifiesTLSAndUsesSCRAMSHA512(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	server.Config.ErrorLog = log.New(io.Discard, "", 0)
	server.StartTLS()
	defer server.Close()

	caFile := filepath.Join(t.TempDir(), "ca.pem")
	if err := os.WriteFile(caFile, pem.EncodeToMemory(&pem.Block{
		Type: "CERTIFICATE", Bytes: server.Certificate().Raw,
	}), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("MQ_KAFKA_USERNAME", "reader")
	t.Setenv("MQ_KAFKA_PASSWORD", "test-password")
	t.Setenv("MQ_KAFKA_CA_FILE", caFile)
	opts, err := OptionsFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	client, err := kgo.NewClient(opts...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	mechanisms := client.OptValue(kgo.SASL).([]sasl.Mechanism)
	if len(mechanisms) != 1 || mechanisms[0].Name() != "SCRAM-SHA-512" {
		t.Fatal("expected SCRAM-SHA-512 authentication")
	}
	config := client.OptValue(kgo.DialTLSConfig).(*tls.Config)
	if config.MinVersion != tls.VersionTLS12 || config.InsecureSkipVerify {
		t.Fatal("expected verified TLS 1.2 or later")
	}
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", server.Listener.Addr().String(), config)
	if err != nil {
		t.Fatalf("CA-trusted broker rejected: %v", err)
	}
	if err := conn.Close(); err != nil {
		t.Fatal(err)
	}
	wrongHost := config.Clone()
	wrongHost.ServerName = "wrong-broker.example"
	conn, err = tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", server.Listener.Addr().String(), wrongHost)
	if err == nil {
		_ = conn.Close()
		t.Fatal("broker hostname mismatch was accepted")
	}
}
