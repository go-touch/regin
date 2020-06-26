package service

import (
	"errors"
	"runtime"
)

// Init system component.
func (a *Application) initCore() {
	// Init Config
	a.config.Init(a.GetAttribute("ConfigPath"))

	// Init logger.
	pattern := ""
	if pattern = a.GetConfig("server.error.pattern").ToString(); pattern == "" {
		pattern = "local"
	}
	a.logger.Init(pattern, a.GetAttribute("RuntimeLogPath"))
}

// The application's  exception message when work.
func (a *Application) SetError(err error) {
	a.err = err
}

// The application's  exception message when work.
func (a *Application) GetError() error {
	return a.err
}

// Error Catch.
func (a *Application) ErrorCatch() {
	if r := recover(); r != nil {
		var array [4096]byte
		buf := array[:]
		runtime.Stack(buf, false)
		stackString := a.exception.Stack(r, buf)

		// Handle error log.
		if openLog := a.GetConfig("main.errorLog.isOpen").ToString(); openLog == "1" {
			_ = a.logger.Record(stackString)
		}

		// Set error message.
		a.SetError(errors.New("there's something wrong with the system"))
	}
}
