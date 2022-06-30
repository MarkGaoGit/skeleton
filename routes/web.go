package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitWebRouters 初始化路由
func InitWebRouters() *gin.Engine {
	r := gin.Default()

	//默认路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "This is Web Api",
			"data": new(interface{}),
		})
	})

	return r
}
