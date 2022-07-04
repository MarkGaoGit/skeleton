package service

import (
	"skeleton/app/global/variable"
	"skeleton/app/model"
	"time"
)

// UserInfoAll 用户的所有信息
type UserInfoAll struct {
	Id            int64     `json:"id"`
	Phone         string    `json:"phone"`
	UserName      string    `json:"user_name"`
	Password      string    `json:"password"`
	Status        int       `json:"status"`
	RegTime       time.Time `json:"reg_time"`
	LastLoginTime time.Time `json:"last_login_time"`
	CreatedTime   time.Time `json:"created_time"`
	UpdateTime    time.Time `json:"update_time"`
	Uid           int64     `json:"uid"`
	OpenId        string    `json:"open_id"`
	BindTime      time.Time `json:"bind_time"`
	VipLevel      int       `json:"vip_level"`
	Photo         string    `json:"photo"`
}

// GetUserList 用户列表查询
func GetUserList(userId int64, page int, size int) (interface{}, int64) {
	var users []model.UserModel
	var count int64

	where := &model.UserModel{
		Id: userId,
	}

	variable.GormDbMysql.Model(model.UserModel{}).Where(where).Count(&count)

	result := variable.GormDbMysql.Where(where).Limit(size).Offset((page - 1) * size).Order("id DESC").Find(&users)

	if result.RowsAffected > 0 {
		return users, count
	} else {
		return variable.DefaultReturnData, 0
	}
}

// GetUserById 根据用户id获取用户信息
func GetUserById(userId int64) interface{} {
	var users UserInfoAll

	row := variable.GormDbMysql.
		Model(model.UserModel{}).
		Select("t_users_info_affiliated.*, t_users_info.*").
		Where(model.UserModel{
			Id: userId,
		}).
		Joins("LEFT JOIN t_users_info_affiliated ON t_users_info_affiliated.uid = t_users_info.id").
		Limit(1).
		Find(&users)

	if row.RowsAffected > 0 {
		return users
	} else {
		return variable.DefaultReturnData
	}

}
