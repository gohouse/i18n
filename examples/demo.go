package main

import (
	"fmt"
	"github.com/gohouse/i18n"
)

func main() {
	var lang = i18n.NewI18n()
	lang.Register("json",i18n.NewI18nDefault())

	var cur = lang.Getter("json")
	err := cur.SetFile("/Users/fizz/go/src/github.com/gohouse/i18n/examples/zh-cn/error.json")
	if err!=nil {
		panic(err.Error())
	}
	res := cur.Load("err_params_format")
	fmt.Println(res)
}
