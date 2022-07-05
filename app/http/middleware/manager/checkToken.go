package manager

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"skeleton/app/global/consts"
	"skeleton/app/model"
	"skeleton/app/utils/redisMe"
	"skeleton/app/utils/response"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required"`
}

// CheckToken 检查是否登陆
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerParams := HeaderParams{}
		if err := c.ShouldBindHeader(&headerParams); err != nil {
			response.Fail(c, consts.RequestErrorHeader, err, nil)
			return
		} else {
			userInfo, er := (new(redisMe.Client)).GetKey(consts.CacheUserTokenPrefix + headerParams.Authorization)
			if er != nil || userInfo == "" {
				response.Fail(
					c,
					consts.BusinessErrorUserNotLogin,
					errors.New(consts.BusinessErrorMap[consts.BusinessErrorUserNotLogin]),
					nil,
				)
				return
			} else {
				var userLoginInfo model.UserModel
				_ = json.Unmarshal([]byte(userInfo), &userLoginInfo)

				c.Set("userLoginInfo", &userLoginInfo)
				c.Next()
			}
		}
	}
}
