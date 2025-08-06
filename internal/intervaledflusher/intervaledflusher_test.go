package intervaledflusher_test

import (
	"runtime"
	"testing"
	"time"

	"dash0.com/otlp-log-processor-backend/internal/intervaledflusher"
	"github.com/alecthomas/assert/v2"
	"github.com/benbjohnson/clock"
)

type fakeFlusher struct {
	count int
}

func (ff *fakeFlusher) Flush() {
	ff.count++
}
func (ff *fakeFlusher) Reset() {
	ff.count = 0
}

func TestIntervaledKeyCounter(t *testing.T) {
	mockClock := clock.NewMock()
	ff := &fakeFlusher{}

	intervaledFlusher := intervaledflusher.NewIntervaledFlusher(
		mockClock,
		time.Second,
		ff,
	)

	intervaledFlusher.Start()

	runtime.Gosched()
	mockClock.Add(5 * time.Second)

	assert.Equal(t, 5, ff.count, "expect flusher to have been called n times")
}
