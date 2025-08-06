package parser

import (
	"fmt"
	"regexp"
)

type RegexpParser struct {
	name string
	re   *regexp.Regexp
}

func NewRegexpParser(matchString string) (*RegexpParser, error) {
	if matchString == "" {
		return nil, fmt.Errorf("invalid regexp: '%s'", matchString)
	}
	// Compile the Regex
	re, err := regexp.Compile(matchString)
	if err != nil {
		return nil, fmt.Errorf("invalid regexp: '%s' %w", matchString, err)
	}

	return &RegexpParser{
		name: "regex",
		re:   re,
	}, nil
}

func (p *RegexpParser) Parse() {

}

func (p *RegexpParser) Name() string {
	return p.name
}
