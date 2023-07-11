package parser

import "strings"

type KeyValueParser struct {
	Separator string
}

func (p KeyValueParser) Parse(input string) (output map[string]string, err error) {
	output = make(map[string]string)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, p.Separator, 2)
		if len(parts) == 2 {
			output[parts[0]] = parts[1]
		}
	}
	return output, nil
}
