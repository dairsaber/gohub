package routes

import (
	controllers "gohub/app/http/controllers/api/v1"

	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {

	// 测试一个v1的路由组

	v1 := r.Group("/v1")

	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			// 用手机注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			// 用Email注册
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			// 登录
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// 刷新token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码 这边是提供给用户自己重置密码的 要在未登录的情况下的重置
			pwc := new(auth.PasswordController)
			authGroup.PUT("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.PUT("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}

		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

		usersGroup := v1.Group("/users")
		{
			// 用户列表分页
			usersGroup.GET("", uc.Index)
			// 更新用户本人profile
			usersGroup.PUT("", middlewares.AuthJWT(), uc.UpdateProfile)
			// 更新邮箱
			usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
			// 更新手机
			usersGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
		}

		// 分类
		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			// 分页分类
			cgcGroup.GET("", cgc.Index)
			// 创建分类
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			// 更新分类
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			// 删除分类
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		// 话题
		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.GET("", middlewares.AuthJWT(), tpc.Index)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
		}
	}

}
