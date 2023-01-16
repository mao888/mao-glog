package glog

import (
	"context"
	"sync"
)

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

//Level logger Level
type Level string

const (
	//DebugLevel has verbose message
	DebugLevel Level = "debug"
	//InfoLevel is default log level
	InfoLevel Level = "info"
	//WarnLevel is for logging messages about possible issues
	WarnLevel Level = "warn"
	//ErrorLevel is for logging errors
	ErrorLevel Level = "error"
	//FatalLevel is for logging fatal messages. The system shutdown after logging the message.
	FatalLevel Level = "fatal"
)

type ctxKey string

const (
	TrackKey ctxKey = "trace_id"
	TimeKey  string = "s_time"
)

var (
	logger   = newConsoleLogger()
	onceInit sync.Once
)

// Init initialize the package level logger instance
// v2.0 版本采用全局logger实例，不再返回实例
// 并发安全的，但多次调用仅会生效一次，应仅全局初始化一次
func Init(options ...Option) (err error) {
	//支持配置修改logger接口实现，目前只支持zap
	onceInit.Do(func() {
		conf := defaultConfiguration
		for _, opt := range options {
			opt.apply(&conf)
		}
		logger = newZapLogger(conf)
	})
	return err
}

// C 从 ctx 中获取打印字段 ，返回带有 ctx 字段的 logger 实例。
func C(ctx context.Context) *zapLogger {
	if ctx == nil {
		return logger
	}
	fields := make([]interface{}, 0, 2)
	if traceID := ctx.Value(TrackKey); traceID != nil {
		fields = append(fields, string(TrackKey), traceID)
	}
	if len(fields) == 0 {
		return logger
	}
	return logger.withFields(fields)
}

func Time(sec float64) *zapLogger {
	return logger.withField(TimeKey, sec)
}

func IsDebug() bool {
	return logger.IsDebug()
}

func ChangeFileStdoutLevel(level Level) {
	logger.ChangeFileStdoutLevel(level)
}

func ChangeConsoleStdoutLevel(level Level) {
	logger.ChangeConsoleStdoutLevel(level)
}
