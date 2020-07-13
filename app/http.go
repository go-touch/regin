package app

import (
	"errors"
	"github.com/go-touch/regin/app/service"
	"github.com/go-touch/regin/base"
	"github.com/go-touch/regin/utils"
	"runtime"
)

type Http struct {
	base.WebServer
	*service.Application
	addrAndPort string
	err         error
}

// Create an new Http Server.
func NewHttp() *Http {
	server := &Http{
		Application: service.App,
	}

	// 获取
	if host := service.App.GetConfig("server.main.httpHost").ToString(); host == "" {
		panic("服务器端口未设置")
	} else {
		server.addrAndPort = host
	}
	return server
}

// Http server work method.
func (h *Http) Work(request *base.Request) *base.Result {
	// Get action.
	paramArray := []string{request.Param("module"), request.Param("controller")}
	if request.Param("action") != "" {
		paramArray = append(paramArray, request.Param("action"))
	}
	actionKey := utils.StringJoinByDot(paramArray...)
	action, err := h.GetRouter().GetGeneral(actionKey)
	if err != nil {
		panic(err)
	}

	// Call Before Action
	if result := action.BeforeExec(request); result.Status != 200 || result.GetData("code") != 0 {
		return result
	}

	// Call Action
	return action.Exec(request)
}

// Http server work method.
func (h *Http) Addr() string {
	return h.addrAndPort
}

func (h *Http) GetError() error {
	return h.err
}

func (h *Http) ErrorCatch() {
	if r := recover(); r != nil {
		var array [4096]byte
		buf := array[:]
		runtime.Stack(buf, false)
		stackString := h.GetException().Stack(r, buf)

		// Handle error log.
		if openLog := h.GetConfig("server.error.log").ToBool(); openLog == true {
			_ = h.GetLogger().Record(stackString)
		}

		// Set error message.
		h.err = errors.New("there's something wrong with the system")
	}
}
