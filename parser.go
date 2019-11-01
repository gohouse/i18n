package i18n

import (
	"github.com/gohouse/e"
	"sync"
)

type Parser struct {
	parser *sync.Map
	opts   *Options
}

var onceParser sync.Once
var parserInit *Parser

func NewParser() *Parser {
	onceParser.Do(func() {
		parserInit = &Parser{
			parser: new(sync.Map),
		}
	})
	return parserInit
}

func (p *Parser) Register(parser string, obj IParser) {
	p.parser.Store(parser, obj)
}

func (p *Parser) Getter(parser string) IParser {
	val, ok := p.parser.Load(parser)
	if ok {
		return val.(IParser)
	}
	return nil
}

func (p *Parser) Load(key string, defaultVal ...string) interface{} {
	var dp = p.Getter(p.opts.DefaultParser)
	dp.SetOptions(p.opts)
	return dp.Load(key, defaultVal...)
}

func (p *Parser) Parse(opts *Options) e.E {
	//fmt.Printf("%#v\n",p.parser)
	var parser = p.Getter(p.opts.DefaultParser)
	if parser==nil {
		return e.New("未注册解析器")
	}
	parser.SetOptions(opts)
	return parser.Parse()
}
