package web

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/http/service"
	"skeleton/app/utils/response"
	"strconv"
)

// UserLists Lists 用户列表
func UserLists(c *gin.Context) {
	userId := c.GetInt64("userId")
	page := c.GetInt("page")
	size := c.GetInt("size")

	result, count := service.GetUserList(userId, page, size)
	response.Success(c, map[string]interface{} {
		"items": result,
		"pages": map[string]interface{} {
			"page": page,
			"size": size,
			"total": count,
		},
	})
}

// UserDetail 用户详细信息
func UserDetail(c *gin.Context) {
	userId := c.Param("uid")
	uid, _ := strconv.Atoi(userId)
	user := service.GetUserById(int64(uid))
	response.Success(c, user)
}