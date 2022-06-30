package registerValidator

import (
	"skeleton/app/core/container"
	"skeleton/app/global/consts"
	"skeleton/app/http/validator/web"
)

// WebRegisterValidator 注册web的验证
func WebRegisterValidator() {
	containers := container.CreateContainersFactory()

	var key string

	key = consts.ValidatorPrefix + "UserRegister"
	containers.Set(key, web.UserRegister{})
}
