package bootstrap

import (
	"log"
	"skeleton/app/global/variable"
	"skeleton/app/http/validator/common/registerValidator"
	"skeleton/app/utils/gormMe"
	"skeleton/app/utils/validatorTranslation"
	"skeleton/app/utils/ymlConfig"
)

func init() {

	//初始化 yml的配置
	variable.ConfigYml = ymlConfig.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()

	//注册参数验证
	registerValidator.WebRegisterValidator()

	//初始化数据库连接
	if dbMysql, err := gormMe.GetOneMysqlClient(); err != nil {
		log.Fatal("数据库初始化失败！" + err.Error())
	} else {
		variable.GormDbMysql = dbMysql
	}

	if err := validatorTranslation.InitTrans("zh"); err != nil {
		log.Fatal("参数验证器语言初始化错误：" + err.Error())
	}
}
