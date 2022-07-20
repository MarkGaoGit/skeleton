package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"skeleton/app/global/consts"
	"skeleton/app/global/variable"
	"skeleton/app/utils/validatorTranslation"
	"strings"
)

// Success 业务返回成功
func Success(c *gin.Context, data interface{}) {
	ReturnJson(c, http.StatusOK, http.StatusOK, "success", data)
}

// Fail 失败返回
func Fail(c *gin.Context, businessCode int, err error, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, businessCode, err.Error(), data)
	c.Abort()
}

// ReturnJson 返回Json格式的数据
func ReturnJson(c *gin.Context, httpCode int, businessCode int, msg string, result interface{}) {
	variable.RoutineWg.Wait()
	c.JSON(httpCode, gin.H{
		"code": businessCode,
		"msg":  msg,
		"data": result,
	})
}

// ErrorSystem 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, http.StatusInternalServerError, msg, data)
	c.Abort()
}

// ValidatorError 翻译表单参数验证器出现的校验错误
func ValidatorError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		wrongParam := validatorTranslation.RemoveTopStruct(errs.Translate(validatorTranslation.Trans))
		ReturnJson(c, http.StatusBadRequest, consts.RequestErrorParams, "参数错误", wrongParam)
	} else {
		errStr := err.Error()
		// multipart:nextpart:eof 错误表示验证器需要一些参数，但是调用者没有提交任何参数
		if strings.ReplaceAll(strings.ToLower(errStr), " ", "") == "multipart:nextpart:eof" {
			ReturnJson(c, http.StatusBadRequest, consts.RequestErrorParams, "参数错误", nil)
		} else {
			ReturnJson(c, http.StatusBadRequest, consts.RequestErrorParams, "参数错误", gin.H{"tips": errStr})
		}
	}
	c.Abort()
}
