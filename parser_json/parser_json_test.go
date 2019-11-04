package parser_json

import (
	"github.com/gohouse/i18n"
	"testing"
)

func TestNewI18nDefault(t *testing.T) {
	var pj = NewParserJson()
	pj.SetOptions(&i18n.Options{
		DefaultParser:  "json",
		DefaultLang:    "zh-cn",
		LangDirectory:  "/Users/fizz/go/src/github.com/gohouse/i18n/examples/language",
		//LangDirectory:  "../examples/language",
		CacheDirectory: "",
	})
	err := pj.Parse()
	if err!=nil {
		t.Error(err.ErrorWithStack())
	}
	res := pj.Load("error.params_missing")
	res2 := pj.Load("error.err2.bb.cc")
	t.Log(res)
	t.Log(res2)
}
