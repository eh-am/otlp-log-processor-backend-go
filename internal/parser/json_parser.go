package parser

import "encoding/json"

type JSONParser struct {
	name string
}

// TODO: should we use this?
//const MaxBodySize int64 = 52428800 // 50MB

func NewJSONParser() *JSONParser {
	return &JSONParser{
		name: "json",
	}
}

func (p *JSONParser) Parse(data []byte) (LogMap, error) {
	var parsed LogMap

	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}

func (p *JSONParser) Name() string {
	return p.name
}
