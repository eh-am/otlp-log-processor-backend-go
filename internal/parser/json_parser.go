package parser

type JSONParser struct {
	name string
}

func NewJSONParser() *JSONParser {
	return &JSONParser{
		name: "json",
	}
}

func (p *JSONParser) Parse() {

}

func (p *JSONParser) Name() string {
	return p.name
}
