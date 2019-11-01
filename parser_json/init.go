package parser_json

import "github.com/gohouse/i18n"

const PARSER = "json"

func init() {
	i18n.NewParser().Register(PARSER, NewParserJson())
}
