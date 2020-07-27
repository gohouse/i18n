package parser_json

import (
	"encoding/json"
	"github.com/gohouse/e"
	"github.com/gohouse/i18n"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ParserJson json解析器对象
type ParserJson struct {
	opts *i18n.Options
	// 示例: /zh_cn/error.json
	// {
	//  "params_format_error": "参数格式有误",
	//  "params_missing": "参数缺失",
	//  "err2": {
	//    "aa": "aaxx",
	//    "bb": "bbxx"
	//  }
	//}
	// map["zh_cn"]["params_format_error"]
	val map[string]map[string]interface{}
}

var _ i18n.IParser = &ParserJson{}

// NewParserJson 初始化json解析器
func NewParserJson() *ParserJson {
	return &ParserJson{val: make(map[string]map[string]interface{})}
}

// SetOptions 注入配置
func (pj *ParserJson) SetOptions(opts *i18n.Options) {
	pj.opts = opts
}

// Parse 执行解析
func (pj *ParserJson) Parse() e.Error {
	// 获取lang目录的所有文件并解析
	var s []string
	// 递归获取lang目录下的所有文件, 返回完整的文件路径数组
	fileAll, err := GetAllFile(pj.opts.LangDirectory, s)
	if err != nil {
		return e.New(err.Error())
	}

	// 去掉目录前缀, 获取语言和文件
	// 解析文件内容, 放入结果集中(pj.val)
	for _, item := range fileAll {
		// 获取除了目录外的文件名字
		fileSuf := strings.Replace(item, pj.opts.LangDirectory, "", 1)
		fileSuf = strings.TrimLeft(fileSuf, "/")
		//fileAllReal = append(fileAllReal, fileSuf)

		// 解析语言和文件名
		split := strings.Split(fileSuf, "/")
		if len(split) != 2 {
			return e.New("目录格式错误")
		}
		// 截取文件名作为一个key, 去掉后缀名
		fileNameStr := strings.TrimRight(split[1], ".json")

		// 读取文件内容为 []byte
		bytes, err := pj.ReadBytesFromFile(item)
		if err != nil {
			return e.New(err.Error())
		}

		// 解析json为map
		var js map[string]interface{}
		err = json.Unmarshal([]byte(string(bytes)), &js)
		if err != nil {
			return e.New(err.Error())
		}

		// 保存到 ParserJson.val 内存中
		langKey := StringToKey(split[0])
		if _, ok := pj.val[langKey]; !ok {
			pj.val[langKey] = make(map[string]interface{})
		}

		// 保存整个文件的内容为解析后的 interface{}
		if pj.opts.EnableFileAsKey {
			pj.val[langKey][fileNameStr] = js
		} else {
			pj.val[langKey] = js
		}
	}
	return nil
}

// StringToKey 中横线转换为下划线
func StringToKey(str string) string {
	return strings.Replace(str, "-", "_", -1)
}

// ReadBytesFromFile 读取文件内容
func (pj *ParserJson) ReadBytesFromFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// LoadWithDefault 获取内容
func (pj *ParserJson) LoadWithDefault(key string, defaultVal ...string) interface{} {
	if key == "" {
		return nil
	}
	var split []string
	// 如果key包含了点,则为多级调用
	if strings.Contains(key, ".") {
		split = strings.Split(key, ".")
	} else { // 如果key不包含点, 则就是直接一级调用
		split = []string{key}
	}

	// 取指定语言的配置
	var currentVal interface{} = pj.val[StringToKey(pj.opts.DefaultLang)]
	for _, item := range split {
		if v, ok := currentVal.(map[string]interface{}); ok {
			if v2, ok2 := v[item]; ok2 {
				currentVal = v2
			} else {
				currentVal = nil
			}
		} else {
			currentVal = nil
		}
	}

	// 如果没有取到且传入了默认值, 则返回默认值
	if currentVal == nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return currentVal
}

// Load 获取单个或所有的配置对象
func (pj *ParserJson) Load(keys ...string) interface{} {
	if len(keys) == 0 {
		return pj.val
	}
	var split = keys
	if len(keys) == 1 {
		// 如果key包含了点,则为多级调用
		if strings.Contains(keys[0], ".") {
			split = strings.Split(keys[0], ".")
		}
	}

	// 取指定语言的配置
	var currentVal interface{} = pj.val[StringToKey(pj.opts.DefaultLang)]
	for _, item := range split {
		if v, ok := currentVal.(map[string]interface{}); ok {
			if v2, ok2 := v[item]; ok2 {
				currentVal = v2
			} else {
				currentVal = nil
			}
		} else {
			currentVal = nil
		}
	}

	return currentVal
}

// GetAllFile 递归读取制定目录下的所有文件, 返回完整文件路径数组
func GetAllFile(dirname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := dirname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				log.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := dirname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
