package parser_test

import (
	"fmt"
	"testing"

	"dash0.com/otlp-log-processor-backend/internal/parser"
	assert "github.com/alecthomas/assert/v2"
)

func TestParserCreator(t *testing.T) {

	tests := []struct {
		Config parser.ParserConfig
		// TODO: maybe we should check directly what instance is?
		WantName string
	}{
		{
			Config: parser.ParserConfig{
				Kind: "json",
			},
			WantName: "json",
		},
	}

	for _, test := range tests {
		test := test
		testName := fmt.Sprintf("parsing kind: '%s'", test.Config.Kind)

		t.Run(testName, func(t *testing.T) {
			got, err := parser.NewParserCreator(&test.Config)
			assert.NoError(t, err, "should not error when creating a parser")

			assert.Equal(t, test.WantName, got.Name())

		})
	}
	assert.Equal(t, 10, 10)
}
