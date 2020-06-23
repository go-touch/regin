package core

import (
	"github.com/go-touch/regin/utils"
	"fmt"
)

type Logger struct {
	pattern string // local、remote
	logPath string
}

// Init Logger
func (l *Logger) Init(pattern string, logPath string) error {
	l.pattern = pattern
	l.logPath = logPath
	return nil
}

// Record log message.
func (l *Logger) Record(stackString string) error {
	switch l.pattern {
	case "local":
		return l.Local(stackString)
	case "remote":
		return l.Local(stackString)
	default:
		return l.Local(stackString)
	}
}

// Record log message by local.
func (l *Logger) Local(stackString string) error {
	fileName := utils.Log.GetFileName(l.logPath)
	if err := utils.Log.Writer(fileName, stackString); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
