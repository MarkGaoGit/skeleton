package gormMe

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"skeleton/app/global/variable"
)

// GetOneMysqlClient 获取一个Mysql数据库实例
func GetOneMysqlClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		variable.ConfigYml.GetString("Mysql.User"),
		variable.ConfigYml.GetString("Mysql.Password"),
		variable.ConfigYml.GetString("Mysql.Host"),
		variable.ConfigYml.GetInt("Mysql.Port"),
		variable.ConfigYml.GetString("Mysql.Databases"),
		variable.ConfigYml.GetString("Mysql.Charset"),
	)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	return client, err
}
