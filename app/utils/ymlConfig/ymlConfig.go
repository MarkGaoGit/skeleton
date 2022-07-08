package ymlConfig

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"skeleton/app/global/variable"
	"skeleton/app/utils/ymlConfig/ymlconfigInterf"
	"time"
)

var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

func CreateYamlFactory(fileName ...string) ymlconfigInterf.YmlConfigInterf {
	v := viper.New()
	v.AddConfigPath(variable.BasePath + "/config")
	if len(fileName) == 0 {
		v.SetConfigName("dev")
	} else {
		v.SetConfigName(fileName[0])
	}
	//设置配置文件类型(后缀)为 yml
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("初始化配置文件发生错误，服务启动需使用：dev、test、stage、prod其中一个配置" + err.Error())
	}

	return &ymlConfig{
		viper: v,
	}

}

type ymlConfig struct {
	viper *viper.Viper
}

func (y *ymlConfig) ConfigFileChangeListen() {
	y.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				lastChangeTime = time.Now()
			}
		}
	})
	y.viper.WatchConfig()
}

// Clone 允许 clone 一个相同功能的结构体
func (y *ymlConfig) Clone(fileName string) ymlconfigInterf.YmlConfigInterf {
	// 这里存在一个深拷贝，需要注意，避免拷贝的结构体操作对原始结构体造成影响
	var ymlC = *y
	var ymlConfViper = *(y.viper)
	(&ymlC).viper = &ymlConfViper

	(&ymlC).viper.SetConfigName(fileName)
	if err := (&ymlC).viper.ReadInConfig(); err != nil {
		//日志
		//variable.ZapLog.Error(my_errors.ErrorsConfigInitFail, zap.Error(err))
	}
	return &ymlC
}

func (y *ymlConfig) Get(keyName string) interface{} {
	return y.viper.Get(keyName)
}

func (y *ymlConfig) GetString(keyName string) string {
	return y.viper.GetString(keyName)
}

func (y *ymlConfig) GetBool(keyName string) bool {
	return y.viper.GetBool(keyName)
}

func (y *ymlConfig) GetInt(keyName string) int {
	return y.viper.GetInt(keyName)
}

func (y *ymlConfig) GetInt32(keyName string) int32 {
	return y.viper.GetInt32(keyName)
}

func (y *ymlConfig) GetInt64(keyName string) int64 {
	return y.viper.GetInt64(keyName)
}

func (y *ymlConfig) GetFloat64(keyName string) float64 {
	return y.viper.GetFloat64(keyName)
}

func (y *ymlConfig) GetDuration(keyName string) time.Duration {
	return y.viper.GetDuration(keyName)
}

func (y *ymlConfig) GetStringSlice(keyName string) []string {
	return y.viper.GetStringSlice(keyName)
}
