package bootstrap

import (
	"log"
	"os"
	"skeleton/app/global/variable"
	"skeleton/app/http/validator/common/registerValidator"
	"skeleton/app/sysLogHook"
	"skeleton/app/utils/gormMe"
	"skeleton/app/utils/validatorTranslation"
	"skeleton/app/utils/ymlConfig"
	"skeleton/app/utils/zapFactory"
)

func init() {
	args := os.Args
	if len(args) < 2 {
		log.Println("没有输入启动环境 默认使用 dev")
		args = append(args, "dev")
	}

	//初始化 yml的配置
	variable.ConfigYml = ymlConfig.CreateYamlFactory(args[1])
	variable.ConfigYml.ConfigFileChangeListen()

	//注册参数验证
	registerValidator.WebRegisterValidator()

	//初始化数据库连接
	if dbMysql, err := gormMe.GetOneMysqlClient(""); err != nil {
		log.Fatal("数据库初始化失败！" + err.Error())
	} else {
		variable.GormDbMysql = dbMysql
	}

	//日志组件初始化
	variable.ZapLog = zapFactory.CreateZapFactory(sysLogHook.ZapLogHandler)

	if err := validatorTranslation.InitTrans("zh"); err != nil {
		log.Fatal("参数验证器语言初始化错误：" + err.Error())
	}
}
