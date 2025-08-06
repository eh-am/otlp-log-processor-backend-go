package manager

import (
	"io"
	"log/slog"
	"time"

	"dash0.com/otlp-log-processor-backend/internal/intervaledflusher"
	"dash0.com/otlp-log-processor-backend/internal/keycounter"
	"dash0.com/otlp-log-processor-backend/internal/parser"
	"dash0.com/otlp-log-processor-backend/internal/parserpipeline"
	"github.com/benbjohnson/clock"
)

type Config struct {
	Interval time.Duration
	Pipeline []parser.ParserConfig
}

type FlusherCounter interface {
	Flush()
	Add(key string)
}

type IntervaledFlusher interface {
	Start()
	Stop()
}

type Manager struct {
	config  *Config
	parser  parserpipeline.Parser
	counter FlusherCounter
	flusher IntervaledFlusher
}

// New creates a Manager
// Which ties together different ideas
// 1. A list of parsers
// 2. A key counter
// 3. A (periodic) flusher
func New(config *Config, clock clock.Clock, writer io.Writer) (*Manager, error) {
	parsers := make([]parserpipeline.Parser, len(config.Pipeline))

	// Instantiate each parser
	for i, c := range config.Pipeline {
		p, err := parser.NewParserCreator(&c)
		parsers[i] = p

		if err != nil {
			return nil, err
		}
	}

	// Create the key counter
	kc := keycounter.NewKeyCounter("", writer)

	// Create the final pipeline which also implements Parse()
	pipelineParser := parserpipeline.NewPipeline(parsers...)
	intervaledFlusher := intervaledflusher.NewIntervaledFlusher(clock, config.Interval, kc)

	// TODO: should we allow the user start separately?
	intervaledFlusher.Start()

	return &Manager{
		config:  config,
		parser:  pipelineParser,
		flusher: intervaledFlusher,
		counter: kc,
	}, nil
}

// Run never fails, in case data is malformed
// It will just print it
func (m *Manager) Run(data []byte) {
	value, err := m.parser.Parse(data)
	if err != nil {
		slog.Error("failed to parse data", "err", err)
	}

	m.counter.Add(value)
}

func (m *Manager) Stop() {
	m.flusher.Stop()
}
