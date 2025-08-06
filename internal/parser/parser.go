package parser

import (
	"errors"
	"fmt"
)

var (
	ErrFailedToCreateRegexpParser = errors.New("failed to create a regexp parser")
)

// ParserConfig contains the configurations for a new parser
type ParserConfig struct {
	// TODO: add more types
	// One of 'json' | 'regex'
	Kind string

	// Only needed if Kind is 'regex'
	MatchString string
}

type Parser interface {
	Parse()
	Name() string
}

func NewParserCreator(cfg *ParserConfig) (Parser, error) {
	switch cfg.Kind {
	case "json":
		return NewJSONParser(), nil
	case "regex":
		regexpParser, err := NewRegexpParser(cfg.MatchString)
		if err != nil {
			return nil, fmt.Errorf("%w: '%w'", ErrFailedToCreateRegexpParser, err)
		}
		return regexpParser, nil
	default:
		// TODO: improve error
		return nil, fmt.Errorf("invalid parser: '%s'", cfg.Kind)
	}
}
