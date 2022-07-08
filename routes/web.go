package routes

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/global/consts"
	"skeleton/app/global/variable"
	"skeleton/app/http/controller/web"
	_ "skeleton/app/http/controller/web"
	mManager "skeleton/app/http/middleware/manager"
	validatorFactory "skeleton/app/http/validator/core/factory"
	"skeleton/app/utils/ginRelease"
	"skeleton/app/utils/response"
)

// InitWebRouters 初始化路由
func InitWebRouters() *gin.Engine {
	var r *gin.Engine
	if variable.ConfigYml.GetString("AppEnv") != "dev" {
		//非开发模式
		//gin.DisableConsoleColor()
		//f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		//gin.DefaultWriter = io.MultiWriter(f)
		r = ginRelease.ReleaseRouter()
	}

	r = gin.Default()

	//默认路由
	r.GET("/", func(c *gin.Context) {
		response.Success(c, "This is Web")
	})

	manager := r.Group("/manager").
		Use(mManager.CheckToken())
	{
		manager.GET("/user/list", (&web.Users{}).UserLists)
		manager.GET("/user/:uid", (&web.Users{}).UserDetail)
		manager.POST("/user/register", validatorFactory.Create(consts.ValidatorPrefix+"UserRegister"))
	}

	//不需要中间件
	r.POST("/user/login", validatorFactory.Create(consts.ValidatorPrefix+"UserLogin"))

	return r
}
