package bootstrap

import (
	"log"
	"skeleton/app/global/variable"
	"skeleton/app/utils/gormMe"
	"skeleton/app/utils/ymlConfig"
)

func init() {

	//初始化 yml的配置
	variable.ConfigYml = ymlConfig.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()

	//初始化数据库连接
	if dbMysql, err := gormMe.GetOneMysqlClient(); err != nil {
		log.Fatal("数据库初始化失败！" + err.Error())
	} else {
		variable.GormDbMysql = dbMysql
	}
}
