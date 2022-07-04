package model

import "time"

type UserAffiliated struct {
	Uid      int64     `gorm:"uid" json:"uid"`
	OpenId   string    `gorm:"open_id" json:"open_id"`
	BindTime time.Time `gorm:"bind_time" json:"bind_time"`
	VipLevel int       `gorm:"vip_level" json:"vip_level"`
	Photo    string    `gorm:"photo" json:"photo"`
}

// TableName 表名
func (u *UserAffiliated) TableName() string {
	return "t_users_info_affiliated"
}
