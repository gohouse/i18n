# i18n
golang i18n, golang实现的多语言解析使用

## 安装
- go mod
```shell script
require github.com/gohouse/i18n master
```
- go get
```shell script
go get github.com/gohouse/i18n
```

## 使用
可以查看包内的示例代码: [https://github.com/gohouse/i18n/blob/master/examples/demo.go](https://github.com/gohouse/i18n/blob/master/examples/demo.go)  

添加语言文件
```shell script
# 创建文件夹
mkdir -p ~/go/src/gopro/language/zh-cn ~/go/src/gopro/language/en-us

# 编写中文语言文件
cat >>~/go/src/gopro/language/zh-cn/error.json<<EOF
{
  "test": "测试",
  "params_format_error": "参数格式有误",
  "params_missing": "参数缺失",
  "err2": {
    "aa": "aaxx",
    "bb": {
      "cc": "cc"
    }
  }
}
EOF

# 编写英文语言文件
cat >>~/go/src/gopro/language/en-us/error.json<<EOF
{
  "test": "just a test",
  "params_format_error": "Incorrect parameters format",
  "params_missing": "Missing parameters",
  "err2": {
    "aa": "aaxx",
    "bb": {
      "cc": "cc"
    }
  }
}
EOF
```
使用demo
```go
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
    
    // 加载error.json文件内的具体配置项
	test := lang.Load("error.test")
	test2 := lang.Load("error.err2.bb.cc")

	fmt.Println(test)
	fmt.Println(test2)
}
```
结果
```shell script
测试
cc
```

## 额外说明
i18n默认提供了json解析器, 同时, 提供了解析器接口, 可以自由定制其他格式的解析器, 如yml,ini,toml等