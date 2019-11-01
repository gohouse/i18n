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
		//LangDirectory:  "/Users/fizz/go/src/github.com/gohouse/i18n/examples/language",
		LangDirectory:  "../examples/language",
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
//func TestNewI18nDefault2(t *testing.T) {
//	var d = NewI18nDefault()
//	d.SetRaw(`{
//  "err_params_format": "参数格式有误",
//  "err_params_missing": "参数缺失",
//  "err2": {
//    "aa": "aaxx",
//    "bb": "bbxx"
//  }
//}`)
//	res := d.Load("err2.aa")
//	t.Log(res)
//}
