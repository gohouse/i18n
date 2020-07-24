package parser_json

import (
	"github.com/gohouse/i18n"
	"testing"
)

func TestNewI18nDefault(t *testing.T) {
	var pj = NewParserJson()
	pj.SetOptions(&i18n.Options{
		DefaultParser:  "json",
		DefaultLang:    "zh_cn",
		LangDirectory:  "/Users/mac/go/src/github.com/gohouse/i18n/examples/language",
		//LangDirectory:  "../examples/language",
		CacheDirectory: "",
	})
	err := pj.Parse()
	if err!=nil {
		t.Error(err.ErrorWithStack())
	}
	res := pj.LoadWithDefault("params_missing")
	res2 := pj.LoadWithDefault("err2.bb.cc")
	t.Log(res)
	t.Log(res2)
}
