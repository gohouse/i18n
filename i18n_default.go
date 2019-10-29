package i18n

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type I18nDefault struct {
	val map[string]interface{}
}

func NewI18nDefault() *I18nDefault {
	return &I18nDefault{val: make(map[string]interface{})}
}

func (e *I18nDefault) SetFile(filePath string) error {
	bytes, err := ReadAll(filePath)
	if err != nil {
		return err
	}
	return e.SetRaw(string(bytes))
}

func ReadAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

// {
//  "err_params_format": "参数格式有误",
//  "err_params_missing": "参数缺失"
//}
func (e *I18nDefault) SetRaw(json_str string) error {
	err := json.Unmarshal([]byte(json_str), &e.val)
	return err
}

func (e *I18nDefault) Load(key string, defaultVal ...string) interface{} {
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

	var currentVal interface{} = e.val
	for _, item := range split {
		if v, ok := currentVal.(map[string]interface{}); ok {
			currentVal = v[item]
		} else {
			currentVal = nil
		}
	}
	return currentVal
}
