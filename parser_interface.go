package i18n

import "github.com/gohouse/e"

type IParser interface {
	SetOptions(opts *Options)
	Parse() e.Error
	LoadWithDefault(key string, defaultVal ...string) interface{}
	Load(keys ...string) interface{}
}
