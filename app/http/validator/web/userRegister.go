package web

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/global/consts"
	"skeleton/app/http/controller/web"
	"skeleton/app/http/validator/core/dataTransfer"
	"skeleton/app/utils/response"
)

type UserRegister struct {
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	UserName string `form:"userName" json:"userName" binding:"required,min=4,max=200"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=10"`
	OpenId   string `form:"openId" json:"openId" binding:"required,len=32"`
	Photo    string `form:"photo" json:"photo" binding:"required,url"`
}

func (r UserRegister) CheckParams(c *gin.Context) {
	if err := c.ShouldBind(&r); err != nil {
		response.ValidatorError(c, err)
		return
	}

	extraAddBindDataContext := dataTransfer.DataAddContext(r, consts.ValidatorPrefix, c)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(c, "UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).UserRegister(extraAddBindDataContext)
	}
}
