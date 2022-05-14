// API 后端接收到 User-Agent 数据后可以暂时不做处理，但是后续有特殊的业务需求时，可以针对某个客户端具体到版本，进行特殊的数据处理。
// 常见的使用场景，是废弃客户端：例如一个银行 APP，升级了交易时的加密算法，低于 5.0 版本的客户端因为安全原因，必须废弃。针对此情况，可通过后端 API 判断 User-Agent 标头，对低于 5.0 的版本的客户端请求，返回专属的数据，如 APP 首页的第一个 Banner 显示请升级客户端，安全升级无法使用的提示。
// 现实生产中，有些客户端用户会关闭系统的应用自动更新功能，多版本客户端是无法避免的问题。有了 User-Agent ，我们可以更加灵活的做针对性处理。
//
package middlewares

import (
	"errors"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 User-Agent 标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}
