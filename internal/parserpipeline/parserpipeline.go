package parserpipeline

type Parser interface {
	Parse([]byte) (string, error)
}

type ParserPipeline struct {
	name string
	fn   func([]byte) (string, error)
}

func NewPipeline(parsers ...Parser) *ParserPipeline {
	return &ParserPipeline{
		name: "pipeline",
		fn:   generatePipeline(parsers...),
	}
}

func generatePipeline(parsers ...Parser) func([]byte) (string, error) {
	return func(input []byte) (string, error) {
		curr := input

		for _, parser := range parsers {
			out, err := parser.Parse(curr)
			if err != nil {
				return "", err
			}

			// TODO: this conversion is unnecessary
			curr = []byte(out)
		}

		return string(curr), nil
	}
}

func (p *ParserPipeline) Name() string {
	return p.name
}

func (p *ParserPipeline) Parse(data []byte) (string, error) {
	return p.fn(data)
}
