package routes

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {

	// 测试一个v1的路由组

	v1 := r.Group("/v1")

	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}

}
