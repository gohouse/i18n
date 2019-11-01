package i18n

import "github.com/gohouse/e"

type IParser interface {
	SetOptions(opts *Options)
	Parse() e.E
	Load(key string, defaultVal ...string) interface{}
}
