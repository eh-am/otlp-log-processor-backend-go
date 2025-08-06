package parserpipeline_test

import (
	"testing"

	"dash0.com/otlp-log-processor-backend/internal/parserpipeline"
	"github.com/alecthomas/assert/v2"
)

type mockDuplicateParser struct{}

// Parse just duplicates the input
func (m *mockDuplicateParser) Parse(input []byte) (string, error) {
	return string(input) + string(input), nil
}

func TestParserPipeline(t *testing.T) {
	mockParser := mockDuplicateParser{}

	pipeline := parserpipeline.NewPipeline(&mockParser, &mockParser, &mockParser)

	got, err := pipeline.Parse([]byte("1"))
	assert.NoError(t, err)

	want := "11111111"
	assert.Equal(t, want, got)
}
