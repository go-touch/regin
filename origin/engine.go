package origin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-touch/regin/base"
	"github.com/unrolled/secure"
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
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	})
	// 注册路由
	ed.origin.Any("/:module/:controller/:action", func(c *gin.Context) {
		// Error catch.
		defer func() {
			if err := server.GetError(); err != nil {
				result := base.ResultInvoker.CreateJson(200, "")
				result.SetData("code", 10000)
				result.SetData("msg", err.Error())
				_ = ed.response.Output(c, result)
			}
		}()
		defer server.ErrorCatch()
		_ = ed.response.Output(c, server.Work(base.RequestInvoker.Factory(c)))
	})

	// 注册路由
	ed.origin.Any("/:module/:controller", func(c *gin.Context) {
		// Error catch.
		defer func() {
			if err := server.GetError(); err != nil {
				result := base.ResultInvoker.CreateJson(200, "")
				result.SetData("code", 10000)
				result.SetData("msg", err.Error())
				_ = ed.response.Output(c, result)
			}
		}()
		defer server.ErrorCatch()
		_ = ed.response.Output(c, server.Work(base.RequestInvoker.Factory(c)))
	})
	_ = ed.origin.Run(server.Addr()) // Run HttpServer
}

// Run HttpsServer
func (ed *EngineDispatcher) HttpsServer(server base.WebServer) {
	// Register middleware.
	ed.origin.Use(func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "127.0.0.1:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			panic(err.Error())
		}
		c.Next()
	})
	// Register router.
	ed.origin.Any("/:module/:action", func(c *gin.Context) {
		request := base.RequestInvoker.Factory(c)
		result := server.Work(request)
		_ = ed.response.Output(c, result)
	})
	// Run server.
	_ = ed.origin.RunTLS("127.0.0.1:8080", "cert.pem", "key.pem")
}
