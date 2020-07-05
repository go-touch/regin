package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Logger struct {
}

var Log *Logger

func init() {
	Log = &Logger{}
}

// 获取文件名称
func (l *Logger) GetFileName(logPath string, path ...string) string {
	fileName := ""                          // 文件名
	separator := string(filepath.Separator) // 分隔符
	datetime := time.Now().String()[0:10]   // 时间戳

	// 拼接filename
	fileName += logPath + separator
	if path != nil {
		fileName += strings.Join(path, separator) + separator
	}
	fileName += datetime + ".txt"
	return fileName
}

// 写入日志
func (l *Logger) Writer(fileName string, content string) error {
	// 打开文件句柄
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 666)

	if err != nil {
		return err
	}
	defer file.Close()

	// 日志文件格式:log包含时间及文件行数
	logger := log.New(file, "", log.LstdFlags)
	logger.Println(content)
	return nil
}
