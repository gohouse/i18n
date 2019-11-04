
package main

import (
	"fmt"
	"github.com/gohouse/i18n"
	// 这里不要忘记引入默认的json驱动
	_ "github.com/gohouse/i18n/parser_json"
)

func main() {
	lang := i18n.NewI18n(
		// 这里指定语言文件路径
		i18n.LangDirectory("/go/src/github.com/gohouse/i18n/examples/language"),

		// 这里如果不i设置, 则默认使用zh-cn
		//i18n.DefaultLang("zh-cn"),

		// 这里如果不i设置, 则默认使用 json,可以自定义解析器和配置文件格式
		//i18n.DefaultParser("json"),
	)

	// 加载error.json文件内的具体配置项, 多级加载, 使用.连接
	test := lang.Load("error.test")
	test2 := lang.Load("error.err2.bb.cc")

	fmt.Println(test)
	fmt.Println(test2)
}