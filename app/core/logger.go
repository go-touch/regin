package core

import (
	"fmt"
	"github.com/go-touch/regin/utils"
	"path/filepath"
	"strings"
	"time"
)

const (
	Info    string = "Info"
	Warning string = "Warning"
	Error   string = "Error"
)

type Logger struct {
	pattern string // local、remote
	logPath string
}

// Init Logger
func (l *Logger) Init(pattern string, logPath string) {
	l.pattern = pattern
	l.logPath = logPath
}

// Record log message by local.
func (l *Logger) Local(logType string, content interface{}, path ...string) error {
	fileName := l.Filename(logType, path...)
	if err := utils.Log.Writer(fileName, content); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

// 获取文件名称
func (l *Logger) Filename(logType string, path ...string) string {
	fileSlice := []string{l.logPath}

	if path != nil {
		fileSlice = append(fileSlice, path...)
	}
	// 文件扩展名
	extName := time.Now().Format("2006-01-02") + "-" + logType + ".txt"
	fileSlice = append(fileSlice, extName)
	return strings.Join(fileSlice, string(filepath.Separator))
}
