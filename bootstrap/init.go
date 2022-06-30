package bootstrap

import (
	"fmt"
	"skeleton/app/global/variable"
	"skeleton/app/utils/yml_config"
)

func init() {

	//初始化 yml的配置
	variable.ConfigYml = yml_config.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()

	fmt.Println(variable.ConfigYml.GetInt("Mysql.Port"))
}