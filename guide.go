package regin

import (
	"github.com/go-touch/regin/app"
	"github.com/go-touch/regin/origin"
)

type Guider struct{}

// 定义Guider
var Guide *Guider

func init() {
	Guide = &Guider{}
}

// 开启http服务
func (e *Guider) HttpService() {
	origin.Engine.HttpServer(app.NewHttp())
}
