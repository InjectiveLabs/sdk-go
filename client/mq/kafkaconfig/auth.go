// Package kafkaconfig configures authenticated Kafka connections for the MQ tools.
package kafkaconfig

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
)

// OptionsFromEnv preserves plaintext connections when all three variables are unset.
// Setting any variable requires the complete TLS and SCRAM-SHA-512 configuration.
func OptionsFromEnv() ([]kgo.Opt, error) {
	username := os.Getenv("MQ_KAFKA_USERNAME")
	password := os.Getenv("MQ_KAFKA_PASSWORD")
	caFile := os.Getenv("MQ_KAFKA_CA_FILE")
	if username == "" && password == "" && caFile == "" {
		return nil, nil
	}
	if username == "" || password == "" || caFile == "" {
		return nil, errors.New("MQ_KAFKA_USERNAME, MQ_KAFKA_PASSWORD and MQ_KAFKA_CA_FILE must be set together")
	}

	caPEM, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read MQ Kafka CA file: %w", err)
	}
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(caPEM) {
		return nil, errors.New("MQ_KAFKA_CA_FILE must contain a PEM certificate")
	}

	return []kgo.Opt{
		kgo.DialTLSConfig(&tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    roots,
		}),
		kgo.SASL(scram.Auth{User: username, Pass: password}.AsSha512Mechanism()),
	}, nil
}
