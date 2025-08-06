package parser_test

import (
	"testing"

	"dash0.com/otlp-log-processor-backend/internal/parser"
	"github.com/alecthomas/assert/v2"
)

func TestRegexpParser(t *testing.T) {
	parser, err := parser.NewRegexpParser(`\{.*\}`)
	assert.NoError(t, err)

	lm, err := parser.Parse([]byte(`"my log body 1" - {"foo":"bar", "baz":"qux"}`))
	assert.NoError(t, err)

	want := `{"foo":"bar", "baz":"qux"}`
	assert.Equal(t, want, lm)
}
