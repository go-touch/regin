package service

import (
	"github.com/go-touch/regin/utils"
	"os"
	"strings"
)

// Init Application.
func (a *Application) init() {
	// Run Mode Include: dev、test、prod
	if a.attribute["RunMode"] = os.Getenv("REGIN_RUNMODE"); a.attribute["RunMode"] == "" {
		a.attribute["RunMode"] = "dev"
	}

	// Application Path ... flag get param yet
	a.attribute["AppPath"], a.err = os.Getwd()
	if a.err != nil {
		panic(a.err)
	}

	// Batch Init System Path.
	a.BatchInitPath(map[string]string{
		"ConfigPath":     "config." + a.attribute["RunMode"],
		"RuntimePath":    "runtime",
		"RuntimeLogPath": "runtime.log",
	})
}

// Batch Init System Path.
func (a *Application) BatchInitPath(pathMap map[string]string) {
	for key, value := range pathMap {
		a.attribute[key] = utils.File.JoinPath(a.attribute["AppPath"], utils.File.JoinPath(strings.Split(value, ".")...))
		if a.err = a.DirExists(a.attribute[key]); a.err != nil {
			panic(a.err)
		}
	}
}

// Judge dir is exist.
func (a *Application) DirExists(path string) (err error) {
	_, err = os.Stat(path)
	return err
}

// Get system attribute.
func (a *Application) GetAttribute(key string) string {
	if attribute, ok := a.attribute[key]; ok {
		return attribute
	}
	return ""
}
