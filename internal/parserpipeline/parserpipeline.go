package parserpipeline

type Parser interface {
	Parse([]byte) (string, error)
}

func NewPipeline(parsers ...Parser) func([]byte) (string, error) {
	return func(input []byte) (string, error) {
		//		var out string
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
