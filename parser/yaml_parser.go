package parser

import "gopkg.in/yaml.v2"

type YAMLParser struct{}

func (p YAMLParser) Parse(input string) (output map[string]string, err error) {
	err = yaml.Unmarshal([]byte(input), &output)
	return output, err
}
