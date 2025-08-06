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
		WantErr  error
	}{
		{
			Config: parser.ParserConfig{
				Kind: "json",
			},
			WantName: "json",
		},
		{
			Config: parser.ParserConfig{
				Kind:        "regex",
				MatchString: `\{.*\}`,
			},
			WantName: "regex",
		},
		// Empty Regex
		{
			Config: parser.ParserConfig{
				Kind: "regex",
			},
			WantName: "regex",
			WantErr:  parser.ErrFailedToCreateRegexpParser,
		},

		// Invalid Regex
		{
			Config: parser.ParserConfig{
				Kind:        "regex",
				MatchString: `(`,
			},
			WantName: "regex",
			WantErr:  parser.ErrFailedToCreateRegexpParser,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("parsing kind:%s", test.Config.Kind)

		t.Run(testName, func(t *testing.T) {
			got, err := parser.NewParserCreator(&test.Config)

			if test.WantErr != nil {
				assert.IsError(t, err, test.WantErr)
			} else {
				assert.Equal(t, test.WantName, got.Name())
			}
		})
	}
}
