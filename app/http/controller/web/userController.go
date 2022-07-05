package web

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/global/consts"
	"skeleton/app/http/service"
	"skeleton/app/utils/encrypt"
	"skeleton/app/utils/response"
	"strconv"
)

type Users struct{}

// UserLists Lists 用户列表
func (u *Users) UserLists(c *gin.Context) {
	userId := c.GetInt64("userId")
	page := c.GetInt("page")
	size := c.GetInt("size")

	result, count := service.GetUserList(userId, page, size)
	response.Success(c, map[string]interface{}{
		"items": result,
		"pages": map[string]interface{}{
			"page":  page,
			"size":  size,
			"total": count,
		},
	})
}

// UserDetail 用户详细信息
func (u *Users) UserDetail(c *gin.Context) {
	userId := c.Param("uid")
	uid, _ := strconv.Atoi(userId)
	user := service.GetUserById(int64(uid))
	response.Success(c, user)
}

// UserRegister 用户注册
func (u *Users) UserRegister(c *gin.Context) {
	phone := c.GetString(consts.ValidatorPrefix + "phone")
	userName := c.GetString(consts.ValidatorPrefix + "userName")
	password := c.GetString(consts.ValidatorPrefix + "password")
	openId := c.GetString(consts.ValidatorPrefix + "openId")
	photo := c.GetString(consts.ValidatorPrefix + "photo")
	createdTime := c.GetString(consts.ValidatorPrefix + "createdTime")

	userId, err := service.UserRegister(map[string]string{
		"phone":       phone,
		"userName":    userName,
		"password":    encrypt.Sha256(password),
		"openId":      openId,
		"photo":       photo,
		"createdTime": createdTime,
	})

	if err != nil {
		response.Fail(c, consts.BusinessErrorUserSave, err, nil)
	} else {
		response.Success(c, map[string]int64{
			"userId": userId,
		})
	}
}

// UserLogin 用户登陆
func (u *Users) UserLogin(c *gin.Context) {
	phone := c.GetString(consts.ValidatorPrefix + "phone")
	password := c.GetString(consts.ValidatorPrefix + "password")
	loginTime := c.GetString(consts.ValidatorPrefix + "createdTime")

	login, err := service.UserLogin(phone, password, loginTime)
	if err != nil {
		response.Fail(c, consts.BusinessErrorUserLoginFail, err, nil)
	} else {
		response.Success(c, login)
	}
}
