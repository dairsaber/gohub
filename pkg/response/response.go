package response

import (
	"gohub/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getResponseBody(success bool, data any, msg string, errors ...any) gin.H {

	responseBody := gin.H{"success": success, "message": msg}

	if success {
		responseBody["data"] = data
	} else {
		if len(errors) > 0 {
			responseBody["errors"] = errors
		}
	}

	return responseBody
}

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
// 执行某个『没有具体返回数据』的『变更』操作成功后调用，例如删除、修改密码、修改手机号
func Success(c *gin.Context) {
	JSON(c, getResponseBody(true, nil, "操作成功"))
}

// Data 响应 200 和带 data 键的 JSON 数据
// 执行『更新操作』成功后调用，例如更新话题，成功后返回已更新的话题
func Data(c *gin.Context, data any) {
	JSON(c, getResponseBody(true, data, "操作成功"))
}

// Created 响应 201 和带 data 键的 JSON 数据
// 执行『新增操作』成功后调用，例如更新话题，成功后返回已更新的话题
func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, getResponseBody(true, data, "操作成功"))
}

// CreatedJSON 响应 201 和 JSON 数据
func CreatedJSON(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, getResponseBody(true, data, "操作成功"))
}

// Abort404 响应 404，未传参 msg 时使用默认消息
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound,
		getResponseBody(false, nil, defaultMessage("数据不存在，请确定请求正确", msg...)))
}

// Abort403 响应 403，未传参 msg 时使用默认消息
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden,
		getResponseBody(false, nil, defaultMessage("权限不足，请确定您有对应的权限", msg...)))

}

// Abort500 响应 500，未传参 msg 时使用默认消息
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, getResponseBody(false, nil, defaultMessage("服务器内部错误，请稍后再试", msg...)))
}

// BadRequest 响应 400，传参 err 对象，未传参 msg 时使用默认消息
// 在解析用户请求，请求的格式或者方法不符合预期时调用
func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	c.AbortWithStatusJSON(http.StatusBadRequest,
		getResponseBody(false, nil, defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
			err.Error(),
		))
}

// Error 响应 404 或 422，未传参 msg 时使用默认消息
// 处理请求时出现错误 err，会附带返回 error 信息，如登录错误、找不到 ID 对应的 Model
func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	// error 类型为『数据库未找到内容』
	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
		getResponseBody(false, nil, defaultMessage("请求处理失败，请查看 error 的值", msg...),
			err.Error(),
		))
}

// ValidationError 处理表单验证不通过的错误，返回的 JSON 示例：
//         {
//             "errors": {
//                 "phone": [
//                     "手机号为必填项，参数名称 phone",
//                     "手机号长度必须为 11 位的数字"
//                 ]
//             },
//             "message": "请求验证不通过，具体请查看 errors"
//         }
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
		getResponseBody(false, nil, "请求验证不通过，具体请查看 errors", errorsToSlice(errors)...))
}

// Unauthorized 响应 401，未传参 msg 时使用默认消息
// 登录失败、jwt 解析失败时调用
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized,
		getResponseBody(false, nil, defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...)))
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}

// 将errors 按field分组 返回error
func errorsToSlice(errors map[string][]string) (errorSlice []any) {
	for k, v := range errors {

		errorSlice = append(errorSlice, map[string]any{
			"field":         k,
			"errorMessages": v,
		})
	}

	return
}