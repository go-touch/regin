package service

import (
	"github.com/go-touch/regin/base"
	"github.com/go-touch/regin/app/core"
)

type Application struct {
	attribute map[string]string   // System Attribute
	config    *core.Config        // Config Storage
	logger    *core.Logger        // core Logger
	exception *core.Exception     // core Exception
	router    *base.RouterStorage // Router Storage
	err       error               // Runtime error
}

// Define BaseServer.
var App *Application

func init() {
	App = &Application{
		attribute: make(map[string]string),
		config:    &core.Config{},
		exception: &core.Exception{},
		logger:    &core.Logger{},
		router:    base.Router,
	}
	defer App.ErrorCatch()
	App.init()
	App.initCore()
	App.initExpand()
}

// Get application ConfigValue.
func (a *Application) GetConfig(args ...string) *base.ConfigValue {
	return a.config.GetConfig(args...)
}

// Get Logger.
func (a *Application) GetLogger() *core.Logger {
	return a.logger
}

// Get Exception
func (a *Application) GetException() *core.Exception {
	return a.exception
}

// Get application config.
func (a *Application) GetRouter() *base.RouterStorage {
	return a.router
}
