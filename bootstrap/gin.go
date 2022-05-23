// 对gin的一些配置操作
package bootstrap

import (
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupGin() {
	setLoggerMode()
}

// 设置 gin 的运行模式，支持 debug, release, test
// release 会屏蔽调试信息，官方建议生产环境中使用
// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
// 故此设置为 release，有特殊情况手动改为 debug 即可
func setLoggerMode() {
	logLevel := config.GetString("app.log_level")

	switch logLevel {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	case "debug":
		fallthrough
	default:
		gin.SetMode(gin.DebugMode)

	}

}
