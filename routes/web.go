package routes

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/global/consts"
	"skeleton/app/http/controller/web"
	_ "skeleton/app/http/controller/web"
	validatorFactory "skeleton/app/http/validator/core/factory"
	"skeleton/app/utils/response"
)

// InitWebRouters 初始化路由
func InitWebRouters() *gin.Engine {
	r := gin.Default()

	//默认路由
	r.GET("/", func(c *gin.Context) {
		response.Success(c, "This is Web")
	})

	manager := r.Group("/manager")
	{
		manager.GET("/user/list", (&web.Users{}).UserLists)
		manager.GET("/user/:uid", (&web.Users{}).UserDetail)
		manager.POST("/user/register", validatorFactory.Create(consts.ValidatorPrefix+"UserRegister"))
	}

	return r
}
