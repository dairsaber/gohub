package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/pkg/response"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册 API 路由
	routes.RegisterApiRoutes(router)

	// 配置 404 路由
	setup404Handler(router)

	// TODO 到时候将静态路径放在配置里面
	router.StaticFS("/public/uploads", http.Dir("./public/uploads"))

}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(middlewares.Logger(), middlewares.Recovery(), middlewares.ForceUA())
}

func setup404Handler(router *gin.Engine) {

	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			response.Abort404(c, "路由未定义，请确认 url 和请求方法是否正确。")

		}
	})

}
