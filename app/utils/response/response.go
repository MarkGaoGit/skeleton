package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 业务返回成功
func Success(c *gin.Context, data interface{}) {
	ReturnJson(c, http.StatusOK, http.StatusOK, "success", data)
}

// Fail 失败返回
func Fail(c *gin.Context, businessCode int, err error, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, businessCode, err.Error(), data)
}

// ReturnJson 返回Json格式的数据
func ReturnJson(c *gin.Context, httpCode int, businessCode int, msg string, result interface{}) {
	c.JSON(httpCode, gin.H{
		"code": businessCode,
		"msg":  msg,
		"data": result,
	})
}
