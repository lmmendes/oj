package parser

var Parsers = map[string]Parser{
	"default": DefaultParser{},
	"kv":      KeyValueParser{Separator: ":"},
	"csv":     CSVParser{},
	"yaml":    YAMLParser{},
	"xml":     XMLParser{},
	"log":     LogParser{},
}
