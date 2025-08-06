package manager_test

import (
	"bytes"
	"os"
	"runtime"
	"testing"
	"time"

	"dash0.com/otlp-log-processor-backend/internal/manager"
	"dash0.com/otlp-log-processor-backend/internal/parser"
	"github.com/alecthomas/assert/v2"
	"github.com/benbjohnson/clock"
)

func loadTestfile(t *testing.T) string {
	b, err := os.ReadFile("./testdata/out")
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}

func TestManagerConfigMultipleOperations(t *testing.T) {
	var buf bytes.Buffer

	want := loadTestfile(t)
	mockClock := clock.NewMock()
	config := &manager.Config{
		Interval: 5 * time.Second,
		Operations: []parser.ParserConfig{
			{
				Kind:        "regex",
				MatchString: `\{.*\}`,
			},
			{
				Kind:            "json",
				FieldOfInterest: "foo",
			},
		},
	}

	m, err := manager.New(config, mockClock, &buf)
	assert.NoError(t, err)
	m.Run([]byte(`"my log body 1" - {"foo":"bar", "baz":"qux"}`))
	m.Run([]byte(`"my log body 2" - {"foo":"qux", "baz":"qux"}`))
	m.Run([]byte(`"my log body 3" - {"baz":"qux"}`))
	m.Run([]byte(`"my log body 4" - {"foo":"baz"}`))
	m.Run([]byte(`"my log body 5" - {"foo":"baz", "baz":"qux"}`))

	// Pass the necessary time
	runtime.Gosched()
	mockClock.Add(5 * time.Second)

	assert.Equal(t, want, buf.String())
}

func TestManagerJSON(t *testing.T) {
	var buf bytes.Buffer

	want := loadTestfile(t)
	mockClock := clock.NewMock()
	config := &manager.Config{
		Interval: 5 * time.Second,
		Operations: []parser.ParserConfig{
			{
				Kind:            "json",
				FieldOfInterest: "foo",
			},
		},
	}

	m, err := manager.New(config, mockClock, &buf)
	assert.NoError(t, err)
	m.Run([]byte(`{"foo":"bar", "baz":"qux"}`))
	m.Run([]byte(`{"foo":"qux", "baz":"qux"}`))
	m.Run([]byte(`{"baz":"qux"}`))
	m.Run([]byte(`{"foo":"baz"}`))
	m.Run([]byte(`{"foo":"baz", "baz":"qux"}`))

	// Pass the necessary time
	runtime.Gosched()
	mockClock.Add(5 * time.Second)

	assert.Equal(t, want, buf.String())
}

func TestManagerRegex(t *testing.T) {
	var buf bytes.Buffer

	mockClock := clock.NewMock()
	config := &manager.Config{
		Interval: 5 * time.Second,
		Operations: []parser.ParserConfig{
			{
				Kind:        "regex",
				MatchString: `(\d+)`,
			},
		},
	}

	m, err := manager.New(config, mockClock, &buf)
	assert.NoError(t, err)
	m.Run([]byte(`"my log body 1"`))
	m.Run([]byte(`"my log body 2"`))
	m.Run([]byte(`"my log body 2"`))
	m.Run([]byte(`"my log body 3"`))
	m.Run([]byte(`"my log body 4"`))

	// Pass the necessary time
	runtime.Gosched()
	mockClock.Add(5 * time.Second)

	assert.Equal(t, `"1" - 1
"2" - 2
"3" - 1
"4" - 1
`, buf.String())
}
