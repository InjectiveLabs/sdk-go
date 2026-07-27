package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func TestParseConfigReadsControlToken(t *testing.T) {
	cmd := newDetectorTestCommand()

	setFlag(t, cmd, flagMQDetectorKafkaBrokers, "broker-1:9092,broker-2:9092")
	setFlag(t, cmd, flagMQDetectorRawTopic, "raw")
	setFlag(t, cmd, flagMQDetectorLatestTopic, "latest")
	setFlag(t, cmd, flagMQDetectorFullNodes, "http://node-1:9999,http://node-2:9999")
	setFlag(t, cmd, flagMQDetectorControlToken, "secret")
	setFlag(t, cmd, flagMQDetectorRequestTimeout, "2s")
	setFlag(t, cmd, flagMQDetectorMessageTimeout, "3s")

	cfg, err := parseConfig(cmd)
	if err != nil {
		t.Fatalf("parseConfig returned error: %v", err)
	}

	if !reflect.DeepEqual(cfg.KafkaBrokers, []string{"broker-1:9092", "broker-2:9092"}) {
		t.Fatalf("unexpected KafkaBrokers: %#v", cfg.KafkaBrokers)
	}

	if !reflect.DeepEqual(cfg.FullNodes, []string{"http://node-1:9999", "http://node-2:9999"}) {
		t.Fatalf("unexpected FullNodes: %#v", cfg.FullNodes)
	}

	if cfg.ControlToken != "secret" {
		t.Fatalf("unexpected ControlToken: %q", cfg.ControlToken)
	}

	if cfg.RequestTimeout != 2*time.Second {
		t.Fatalf("unexpected RequestTimeout: %s", cfg.RequestTimeout)
	}

	if cfg.MessageTimeout != 3*time.Second {
		t.Fatalf("unexpected MessageTimeout: %s", cfg.MessageTimeout)
	}
}

func newDetectorTestCommand() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.Flags().StringSlice(flagMQDetectorKafkaBrokers, []string{}, "")
	cmd.Flags().String(flagMQDetectorRawTopic, "", "")
	cmd.Flags().String(flagMQDetectorLatestTopic, "", "")
	cmd.Flags().StringSlice(flagMQDetectorFullNodes, []string{}, "")
	cmd.Flags().String(flagMQDetectorControlToken, "", "")
	cmd.Flags().Duration(flagMQDetectorRequestTimeout, 10*time.Second, "")
	cmd.Flags().Duration(flagMQDetectorMessageTimeout, 30*time.Second, "")

	return cmd
}

func setFlag(t *testing.T, cmd *cobra.Command, name, value string) {
	t.Helper()

	if err := cmd.Flags().Set(name, value); err != nil {
		t.Fatalf("failed to set flag %q: %v", name, err)
	}
}
