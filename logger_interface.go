package glog

import "context"

// LoggerInterface is our contract for the logger
// 日志实现约定
type LoggerInterface interface {
	// IsDebug 判断日志级别
	IsDebug() bool
	// ChangeConsoleStdoutLevel 动态修改控制台输出等级
	ChangeConsoleStdoutLevel(level Level)

	// ChangeFileStdoutLevel 动态修改写入文件日志等级
	ChangeFileStdoutLevel(level Level)

	Debugf(format string, args ...interface{})

	Debug(args ...interface{})

	Infof(format string, args ...interface{})

	Info(args ...interface{})

	Warnf(format string, args ...interface{})

	Warn(args ...interface{})

	Errorf(format string, args ...interface{})

	Error(args ...interface{})

	Fatalf(format string, args ...interface{})

	Fatal(args ...interface{})

	Panicf(format string, args ...interface{})

	Panic(args ...interface{})

	C(ctx context.Context) *zapLogger
}
