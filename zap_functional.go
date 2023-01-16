package glog

import "context"

// Debugf 打印用于调试的格式化信息，日志 level 为 DebugLevel 时才会打印
func Debugf(ctx context.Context, format string, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Debugf(format, args...)
}

// Debug 打印用于调试的信息，日志 level 为 DebugLevel 时才会打印
func Debug(ctx context.Context, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Debug(args...)
}

// Infof 打印格式化标准信息，默认的日志信息
func Infof(ctx context.Context, format string, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Infof(format, args...)
}

// Info 打印标准信息，默认的日志信息
func Info(ctx context.Context, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Info(args...)
}

// Warnf 打印比 Info 更重要，但暂不需要单独人为介入的格式化信息
func Warnf(ctx context.Context, format string, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Warnf(format, args...)
}

// Warn 打印比 Info 更重要，但暂不需要单独人为介入的信息
func Warn(ctx context.Context, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Warn(args...)
}

// Errorf 打印格式化错误信息，通常需要人为介入
func Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Errorf(format, args...)
}

// Error 打印错误信息，通常需要人为介入
func Error(ctx context.Context, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Error(args...)
}

// Panicf 打印格式化错误信息，然后 panics
func Panicf(ctx context.Context, format string, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Panicf(format, args...)
}

// Panic 打印错误信息，然后 panics
func Panic(ctx context.Context, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Panic(args...)
}

// Fatalf 打印错误信息，然后调用 os.Exit(1), 业务不使用
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Fatalf(format, args...)
}

// Fatal 打印错误信息，然后调用 os.Exit(1), 业务不使用
func Fatal(ctx context.Context, args ...interface{}) {
	logger.C(ctx).sugaredLogger.Fatal(args...)
}

// InfoWithField 打印带自定义字段的日志
func InfoWithField(ctx context.Context, fields Fields, args ...interface{}) {
	if fields == nil {
		logger.C(ctx).sugaredLogger.Info(args...)
		return
	}
	newFields := make([]interface{}, 0, 2*len(fields))
	for key, val := range fields {
		newFields = append(newFields, key, val)
	}
	logger.C(ctx).withFields(newFields).sugaredLogger.Info(args...)
}
