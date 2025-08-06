package parser

import "fmt"

// ParserConfig contains the configurations for a new parser
type ParserConfig struct {
	// TODO: add more types
	// One of 'json' | 'regex'
	Kind string
}

type Parser interface {
	Parse()
	Name() string
}

func NewParserCreator(cfg *ParserConfig) (Parser, error) {
	switch cfg.Kind {
	case "json":
		return NewJSONParser(), nil
	default:
		// TODO: improve error
		return nil, fmt.Errorf("invalid parser: '%s'", cfg.Kind)

	}
}

//func (p *Parser) Parse() {
//	// TODO: return a map?
//
//}
