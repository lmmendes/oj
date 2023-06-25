package parser

import (
	"strings"
)

type Parser interface {
	Parse(input string) (output map[string]string, err error)
}

type DefaultParser struct{}

func (p DefaultParser) Parse(input string) (output map[string]string, err error) {
	output = make(map[string]string)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			output[parts[0]] = parts[1]
		}
	}
	return output, nil
}
