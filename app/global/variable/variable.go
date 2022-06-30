package variable

import (
	"gorm.io/gorm"
	"os"
	"skeleton/app/utils/ymlConfig/ymlconfigInterf"
)

var (
	BasePath          string                  //项目基础路径
	DateFormat        = "2006-01-02 15:04:05" //全局的时间格式
	DefaultReturnData = make([]int, 0)        //默认没有数据的时候返回一个空的整形切片

	ConfigYml ymlconfigInterf.YmlConfigInterf // 全局配置文件指针

	GormDbMysql *gorm.DB //数据库指针
)

func init() {
	BasePath, _ = os.Getwd()
}
