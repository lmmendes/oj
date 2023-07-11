package parser

import "fmt"

type LogParser struct{}

func (p LogParser) Parse(input string) (output map[string]string, err error) {
	return nil, fmt.Errorf("LogParser not implemented")
}
