package bootstrap

import (
	"log"
	"skeleton/app/global/variable"
	"skeleton/app/utils/gorm_me"
	"skeleton/app/utils/yml_config"
)

func init() {

	//初始化 yml的配置
	variable.ConfigYml = yml_config.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()

	//初始化数据库连接
	if dbMysql, err := gorm_me.GetOneMysqlClient(); err != nil {
		log.Fatal("数据库初始化失败！" + err.Error())
	} else {
		variable.GormDbMysql = dbMysql
	}
}
