package main

import (
	"flag"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestParseConfigReadsControlToken(t *testing.T) {
	oldCommandLine := flag.CommandLine
	oldArgs := os.Args
	t.Cleanup(func() {
		flag.CommandLine = oldCommandLine
		os.Args = oldArgs
	})

	flag.CommandLine = flag.NewFlagSet(t.Name(), flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)
	os.Args = []string{
		"detector",
		"--" + flagMQDetectorKafkaBrokers, "broker-1:9092,broker-2:9092",
		"--" + flagMQDetectorRawTopic, "raw",
		"--" + flagMQDetectorLatestTopic, "latest",
		"--" + flagMQDetectorFullNodes, "http://node-1:9999,http://node-2:9999",
		"--" + flagMQDetectorControlToken, "secret",
		"--" + flagMQDetectorRequestTimeout, "2s",
		"--" + flagMQDetectorMessageTimeout, "3s",
	}

	cfg, err := parseConfig()
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
