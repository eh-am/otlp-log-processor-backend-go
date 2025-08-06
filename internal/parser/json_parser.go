package parser

import (
	"encoding/json"
	"fmt"
)

type JSONParser struct {
	name            string
	fieldOfInterest string
}

// TODO: should we use this?
//const MaxBodySize int64 = 52428800 // 50MB

func NewJSONParser(fieldOfInterest string) *JSONParser {
	return &JSONParser{
		name:            "json",
		fieldOfInterest: fieldOfInterest,
	}
}

func (p *JSONParser) Parse(data []byte) (string, error) {
	//	var parsed LogMap
	var parsed map[string]interface{}

	if err := json.Unmarshal(data, &parsed); err != nil {
		return "", err
	}

	// TODO: is this type assertion harmful?
	return fmt.Sprintf("%v", parsed[p.fieldOfInterest]), nil
}

func (p *JSONParser) Name() string {
	return p.name
}
