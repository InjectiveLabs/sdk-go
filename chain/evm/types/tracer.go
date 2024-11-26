package types

import (
	"os"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/params"
)

const (
	TracerAccessList = "access_list"
	TracerJSON       = "json"
	TracerStruct     = "struct"
	TracerMarkdown   = "markdown"
	TracerFirehose   = "firehose"
)

// NewTracer creates a new Logger tracer to collect execution traces from an
// EVM transaction.
func NewTracer(tracer string, msg *core.Message, rules params.Rules) *tracers.Tracer {
	// TODO: enable additional log configuration
	logCfg := &logger.Config{
		Debug: true,
	}

	var hooks *tracing.Hooks

	switch tracer {
	case TracerAccessList:
		preCompiles := vm.DefaultActivePrecompiles(rules)
		hooks = logger.NewAccessListTracer(msg.AccessList, msg.From, *msg.To, preCompiles).Hooks()
	case TracerJSON:
		hooks = logger.NewJSONLogger(logCfg, os.Stderr)
	case TracerMarkdown:
		hooks = logger.NewMarkdownLogger(logCfg, os.Stdout).Hooks() // TODO: Stderr ?
	case TracerStruct:
		hooks = logger.NewStructLogger(logCfg).Hooks()
	default:
		// Use noop tracer by default
		hooks, _ = tracers.LiveDirectory.New("noop", nil)
	}

	return &tracers.Tracer{
		Hooks: hooks,
	}
}

// TxTraceResult is the result of a single transaction trace during a block trace.
type TxTraceResult struct {
	Result interface{} `json:"result,omitempty"` // Trace results produced by the tracer
	Error  string      `json:"error,omitempty"`  // Trace failure produced by the tracer
}
