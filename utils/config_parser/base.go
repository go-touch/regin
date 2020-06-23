package config_parser

type BaseParser interface {
	ParserToMap(string) map[string]interface{}
}