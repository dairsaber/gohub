package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {

	// 测试一个v1的路由组

	v1 := r.Group("/v1")

	{
		// 注册一个路由
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World",
			})
		})
	}

}
