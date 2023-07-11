package parser

import "fmt"

type XMLParser struct{}

func (p XMLParser) Parse(input string) (output map[string]string, err error) {
	return nil, fmt.Errorf("XMLParser not implemented")
}
