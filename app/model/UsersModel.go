package model

type UserModel struct {
	Id            int64  `gorm:"id" json:"id"`
	Phone         string `gorm:"phone" json:"phone"`
	UserName      string `gorm:"user_name" json:"user_name"`
	Password      string `gorm:"password" json:"-"`
	Status        int    `gorm:"status" json:"status"`
	RegTime       string `gorm:"reg_time" json:"reg_time"`
	LastLoginTime string `gorm:"last_login_time" json:"last_login_time"`
	CreatedTime   string `gorm:"created_time" json:"created_time"`
	UpdatedTime   string `gorm:"updated_time" json:"updated_time"`
}

// TableName 表名
func (u *UserModel) TableName() string {
	return "t_users_info"
}
