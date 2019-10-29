package i18n

type Ii18n interface {
	SetFile(filePath string) error
	SetRaw(conf_str string) error
	Load(key string, defaultVal ...string) interface{}
}