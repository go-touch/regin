package core

import (
	"fmt"
	"github.com/go-touch/regin/utils"
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
