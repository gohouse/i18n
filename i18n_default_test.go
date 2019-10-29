package i18n

import (
	"testing"
)

func TestNewI18nDefault(t *testing.T) {
	var d = NewI18nDefault()
	d.SetRaw(`{
  "err_params_format": "参数格式有误",
  "err_params_missing": "参数缺失"
}`)
	res := d.Load("err_params_format")
	t.Log(res)
}
func TestNewI18nDefault2(t *testing.T) {
	var d = NewI18nDefault()
	d.SetRaw(`{
  "err_params_format": "参数格式有误",
  "err_params_missing": "参数缺失",
  "err2": {
    "aa": "aaxx",
    "bb": "bbxx"
  }
}`)
	res := d.Load("err2.aa")
	t.Log(res)
}
