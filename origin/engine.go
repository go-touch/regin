package origin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-touch/regin/base"
	"net/http"
)

type EngineDispatcher struct {
	origin   *gin.Engine
	response *ResponseHandler
}

// 定义RouterHandler
var Engine *EngineDispatcher

func init() {
	Engine = &EngineDispatcher{
		origin:   gin.Default(),
		response: Response,
	}
}

// Run HttpServer
func (ed *EngineDispatcher) HttpServer(server base.WebServer) {
	// 解决跨域
	ed.origin.Use(func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	})

	// 注册路由
	/*ed.origin.Any("/:module/:controller/:action", func(c *gin.Context) {
		// Error catch.
		defer func() {
			if r := recover(); r != nil {
				// Error catch.
				server.ErrorCatch(Error(r))
				// 打印
				fmt.Println(Error(r))
				// Exception response.
				result := base.ResultInvoker.CreateJson(200, "")
				result.SetData("code", 10000)
				result.SetData("msg", "There's something wrong with the server.")
				_ = ed.response.Output(c, result)
			}
		}()
		_ = ed.response.Output(c, server.Work(base.RequestInvoker.Factory(c)))
	})*/

	// 注册路由
	ed.origin.Any("/api/:module/:controller", func(c *gin.Context) {
		// Error catch.
		defer func() {
			if r := recover(); r != nil {
				// Error catch.
				server.ErrorCatch(Error(r))
				// 打印
				fmt.Println(Error(r))
				// Exception response.
				result := base.ResultInvoker.CreateJson(200, "")
				result.SetData("code", 10000)
				result.SetData("msg", "There's something wrong with the server.")
				_ = ed.response.Output(c, result)
			}
		}()
		_ = ed.response.Output(c, server.Work(base.RequestInvoker.Factory(c)))
	})

	// 静态资源
	ed.origin.Static("/apidoc", "./apidoc")
	ed.origin.LoadHTMLFiles("apidoc/index.html")

	_ = ed.origin.Run(server.Addr()) // Run HttpServer
}

// 获取key
func (ed *EngineDispatcher) routerName(moduleName string, joinPathStr string) {




}

// 异常捕获
func (ed *EngineDispatcher) errorCatch(server base.WebServer) {
	if r := recover(); r != nil {
		// Error catch.
		server.ErrorCatch(Error(r))
		// 打印
		fmt.Println(Error(r))
		// Exception response.
		result := base.ResultInvoker.CreateJson(200, "")
		result.SetData("code", 10000)
		result.SetData("msg", "There's something wrong with the server.")
		//_ = ed.response.Output(c, result)
	}
}
