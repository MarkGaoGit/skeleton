package ginRelease

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"skeleton/app/global/variable"
	"skeleton/app/utils/response"
)

// ReleaseRouter 根据 gin 路由包官方的建议，gin 路由引擎如果在生产模式使用，官方建议设置为 release 模式
// 官方原版提示说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// 这里我们将按照官方指导进行生产模式精细化处理
// 如果部署到生产环境，请使用以下模式：
// 1.生产模式(release) 和开发模式的变化主要是禁用 gin 记录接口访问日志，
// 2.go服务就必须使用nginx作为前置代理服务，这样也方便实现负载均衡
// 3.如果程序发生 panic 等异常使用自定义的 panic 恢复中间件拦截、记录到日志
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式禁用 gin 输出接口访问日志
	// 具体的接口日志自己在业务中另外输出
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()
	// 载入gin的中间件，关键是第二个中间件，对它进行了自定义重写，将可能的 panic 异常等，统一使用 zaplog 接管，保证全局日志打印统一
	engine.Use(gin.Logger(), CustomRecovery())
	return engine
}

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即可
		// 这里的 err 数据类型为 ：runtime.boundsError  ，需要转为普通数据类型才可以输出
		response.ErrorSystem(c, "", fmt.Sprintf("%s", err))
	})
}

//PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	errStr := string(b)
	err = errors.New(errStr)
	variable.ZapLog.Error("服务器内部发生代码执行错误", zap.String("msg", errStr))
	return len(errStr), err
}
