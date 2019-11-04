package main

import (
	"fmt"
	"github.com/gohouse/i18n"
	_ "github.com/gohouse/i18n/parser_json"
)

func main() {
	lang := i18n.NewI18n(
		i18n.LangDirectory("/Users/fizz/go/src/github.com/gohouse/i18n/examples/language"),
		//i18n.DefaultLang("zh-cn"),
		//i18n.DefaultParser("json"),
	)

	//i18n.NewParser().Register("json",parser_json.NewParserJson())

	test := lang.Load("error.test")
	test2 := lang.Load("error.err2.bb.cc")

	fmt.Println(test)
	fmt.Println(test2)
}
