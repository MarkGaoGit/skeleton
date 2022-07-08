package sysLogHook

import (
	"go.uber.org/zap/zapcore"
)

// ZapLogHandler 系统运行日志钩子函数
// 1.单条日志就是一个结构体格式，本函数拦截每一条日志，可以进行后续处理，例如：推送到阿里云日志管理面板、ElasticSearch 日志库等
func ZapLogHandler(entry zapcore.Entry) error {

	//这里启动一个协程，不影响程序性能，
	go func(paramEntry zapcore.Entry) {
		// todo 可以在这里继续处理系统日志 推送elk 等等
	}(entry)
	return nil
}
