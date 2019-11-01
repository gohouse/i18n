package i18n

import (
	"sync"
)

// Options 配置
type Options struct {
	DefaultParser  string
	DefaultLang    string
	LangDirectory  string
	CacheDirectory string
}

// Option 配置驱动
type Option func(*Options)

type I18n struct {
	opts   *Options
	parser *Parser
}

var once sync.Once
var i18nInit *I18n

// NewI18n 初始化包对象函数
func NewI18n(opt ...Option) *I18n {
	once.Do(func() {
		i18nInit = &I18n{
			opts: &Options{},
		}
	})
	// 初始化配置
	i18nInit.initOption(opt...)

	// 初始化解析器
	i18nInit.initParser()

	return i18nInit
}

func (i *I18n) initOption(opt ...Option) {
	for _, o := range opt {
		o(i.opts)
	}
}

func (i *I18n) initParser() {
	i.parser = NewParser()
	err := i.parser.Parse(i.opts)
	if err!=nil {
		panic(err.ErrorWithStack())
	}
}

func (i *I18n) Load(key string, defaultVal ...string) interface{} {
	return i.parser.Load(key, defaultVal...)
}

// LangDirectory 存放不同语言的目录
func LangDirectory(ld string) Option {
	return func(o *Options) {
		o.LangDirectory = ld
	}
}

// DefaultLang 默认语言
func DefaultLang(l string) Option {
	return func(o *Options) {
		o.DefaultLang = l
	}
}

// DefaultParser 默认配置解析器
func DefaultParser(p string) Option {
	return func(o *Options) {
		o.DefaultParser = p
	}
}

// CacheDirectory 配置解析器后的缓存目录
func CacheDirectory(p string) Option {
	return func(o *Options) {
		o.DefaultParser = p
	}
}
