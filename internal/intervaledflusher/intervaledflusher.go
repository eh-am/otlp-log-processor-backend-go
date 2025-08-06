package intervaledflusher

import (
	"log/slog"
	"time"

	"github.com/benbjohnson/clock"
)

type Flusher interface {
	Flush()
}

type IntervaledFlusher struct {
	flusher  Flusher
	interval time.Duration
	quitCh   chan struct{}
	clock    clock.Clock
}

// NewIntervaledFlusher creates a Intervaled Flusher
// Which calls its Flusher every duration time
func NewIntervaledFlusher(clock clock.Clock, interval time.Duration, flusher Flusher) *IntervaledFlusher {
	return &IntervaledFlusher{
		flusher:  flusher,
		interval: interval,
		quitCh:   make(chan struct{}),
		clock:    clock,
	}
}

// Starts a ticker in a go routine which Flushes every duration
func (ifl *IntervaledFlusher) Start() {
	slog.Info("Starting intervaled flusher", "interval", ifl.interval.String())
	ticker := ifl.clock.Ticker(ifl.interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				slog.Debug("flushing")
				ifl.flusher.Flush()
			case <-ifl.quitCh:
				ticker.Stop()
				return
			}
		}

	}()
}

// Stop closes the ticker
func (ifl *IntervaledFlusher) Stop() {
	close(ifl.quitCh)
}
