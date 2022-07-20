package service

import (
	"encoding/json"
	"errors"
	"skeleton/app/global/consts"
	"skeleton/app/global/variable"
	"skeleton/app/model"
	"skeleton/app/utils/encrypt"
	"skeleton/app/utils/localTime"
	"skeleton/app/utils/redisMe"
	"skeleton/app/utils/stringMe"
	"time"
)

// UserInfoAll 用户的所有信息
type UserInfoAll struct {
	Id            int64               `json:"id"`
	Phone         string              `json:"phone"`
	UserName      string              `json:"user_name"`
	Status        int                 `json:"status"`
	RegTime       localTime.LocalTime `json:"reg_time"`
	LastLoginTime localTime.LocalTime `json:"last_login_time"`
	CreatedTime   localTime.LocalTime `json:"created_time"`
	UpdatedTime   localTime.LocalTime `json:"updated_time"`
	Uid           int64               `json:"uid"`
	OpenId        string              `json:"open_id"`
	BindTime      localTime.LocalTime `json:"bind_time"`
	VipLevel      int                 `json:"vip_level"`
	Photo         string              `json:"photo"`
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

// UserRegister 用户注册
func UserRegister(user map[string]string) (int64, error) {
	var err error

	userBaseInfo := &model.UserModel{
		UserName:    user["userName"],
		Phone:       user["phone"],
		Password:    user["password"],
		RegTime:     user["createdTime"],
		CreatedTime: user["createdTime"],
		UpdatedTime: user["createdTime"],
	}

	//事务开启
	db := variable.GormDbMysql.Begin()

	saved := db.Omit("id", "last_login_time").Create(&userBaseInfo)
	if saved.RowsAffected < 1 {
		err = errors.New(consts.BusinessErrorMap[consts.BusinessErrorUserSave] + saved.Error.Error())
		db.Rollback()
		return userBaseInfo.Id, err
	}

	userAffiliatedData := &model.UserAffiliated{
		Uid:      userBaseInfo.Id,
		OpenId:   user["openId"],
		BindTime: user["createdTime"],
		Photo:    user["photo"],
	}
	saved = db.Create(userAffiliatedData)
	if saved.RowsAffected < 1 {
		err = errors.New(consts.BusinessErrorMap[consts.BusinessErrorUserSave] + saved.Error.Error())
		db.Rollback()
		return userBaseInfo.Id, err
	}

	db.Commit()
	return userBaseInfo.Id, nil
}

// UserLogin 用户登陆
// @desc 并没有对token未失效就登陆进行检查 只是redis的使用方法示例
func UserLogin(phone, password, LoginTime string) (interface{}, error) {
	var err error
	var userM model.UserModel

	where := &model.UserModel{
		Phone: phone,
	}

	res := variable.GormDbMysql.Where(where).Limit(1).Find(&userM)
	if res.RowsAffected == 0 {
		err = errors.New(consts.BusinessErrorMap[consts.BusinessErrorUserNotFund] + res.Error.Error())
		return nil, err
	}

	if userM.Password != encrypt.Sha256(password) {
		err = errors.New(consts.BusinessErrorMap[consts.BusinessErrorUserPassword])
		return nil, err
	}

	//生成一个随机的touKen
	token := encrypt.MD5(stringMe.RandomStr(20) + time.Now().String())

	//用户信息json后存入redis
	jsonData, _ := json.Marshal(&userM)
	err = (new(redisMe.Client)).SetKey(
		consts.CacheUserTokenPrefix+token,
		string(jsonData),
		time.Duration(consts.CacheUserTokenTTL)*time.Second,
	)
	if err != nil {
		return nil, err
	}

	// 正常的Done方法应该再具体的协程函数中
	variable.RoutineWg.Add(1)
	go variable.GormDbMysql.Model(&userM).Update("last_login_time", LoginTime)
	variable.RoutineWg.Done()

	return map[string]interface{}{
		"token": token,
		"ttl":   consts.CacheUserTokenTTL,
	}, nil
}
