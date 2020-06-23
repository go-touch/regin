package utils

import (
	"strings"
)

type DataHandler struct {

}

// 定义DataHandler
var Handler *DataHandler

func init() {
	Handler = &DataHandler{}
}

func StringJoinByDot(str ...string) string {
	return strings.Join(str, ".")
}