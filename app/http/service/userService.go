package service

import (
	"skeleton/app/global/variable"
	"skeleton/app/model"
)

// GetUserList 用户列表查询
func GetUserList(userId int64, page int, size int) (interface{}, int64) {
	var users []model.UserModel
	var count int64

	where := &model.UserModel{
		Id : userId,
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
	var users model.UserModel
	row := variable.GormDbMysql.Where(model.UserModel{
		Id: userId,
	}).Limit(1).Find(&users)

	if row.RowsAffected > 0 {
		return users
	} else {
		return variable.DefaultReturnData
	}

}