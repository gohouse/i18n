package i18n

import "sync"

type I18n struct {
	*sync.Map
}

var once sync.Once
var i18nInit *I18n

func NewI18n() *I18n {
	once.Do(func() {
		i18nInit = &I18n{Map: new(sync.Map)}
	})
	return i18nInit
}

func (i *I18n) Register(dataFormat string, obj Ii18n) {
	i.Store(dataFormat, obj)
}

func (i *I18n) Getter(dataFormat string) Ii18n {
	val, ok := i.Load(dataFormat)
	if ok {
		return val.(Ii18n)
	}
	return nil
}
