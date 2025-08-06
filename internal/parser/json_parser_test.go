package parser_test

import (
	"testing"

	"dash0.com/otlp-log-processor-backend/internal/parser"
	"github.com/alecthomas/assert/v2"
)

func TestJSONParser(t *testing.T) {
	parser := parser.NewJSONParser()

	lm, err := parser.Parse([]byte(`{"foo":"bar", "baz":"qux"}`))
	assert.NoError(t, err)

	want := map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
	}

	assert.Equal(t, want, lm)
}
