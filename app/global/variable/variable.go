package variable

import (
	"os"
	"skeleton/app/utils/yml_config/ymlconfig_interf"
)

var (
	BasePath string
	ConfigYml	ymlconfig_interf.YmlConfigInterf	// 全局配置文件指针
)

func init() {
	BasePath, _ = os.Getwd()
}