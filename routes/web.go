package routes

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/http/controller/web"
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
		manager.GET("/user/list", web.UserLists)
		manager.GET("/user/:uid", web.UserDetail)
	}

	return r
}
