package factory

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/core/container"
	"skeleton/app/http/validator/core/interf"
)

// Create 表单参数验证器工厂
func Create(key string) func(context *gin.Context) {
	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}

	return nil
}
