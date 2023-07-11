package parser

import (
	"encoding/csv"
	"strings"
)

type CSVParser struct{}

func (p CSVParser) Parse(input string) (output map[string]string, err error) {
	reader := csv.NewReader(strings.NewReader(input))
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	output = make(map[string]string)
	for _, line := range lines {
		if len(line) == 2 {
			output[line[0]] = line[1]
		}
	}
	return output, nil
}
