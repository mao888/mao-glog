// 封装zap日志

package glog

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type zapLogger struct {
	cfg                configuration
	sugaredLogger      *zap.SugaredLogger
	fileAtomicLevel    *zap.AtomicLevel
	consoleAtomicLevel *zap.AtomicLevel
}

var _ LoggerInterface = (*zapLogger)(nil)

func getEncoder(isJSON bool, keys CoverDefaultKey) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	if keys.TimeKey != "" {
		encoderConfig.TimeKey = keys.TimeKey
	}
	if keys.LevelKey != "" {
		encoderConfig.LevelKey = keys.LevelKey
	}
	if keys.CallerKey != "" {
		encoderConfig.CallerKey = keys.CallerKey
	}
	if keys.MessageKey != "" {
		encoderConfig.MessageKey = keys.MessageKey
	}
	if keys.StacktraceKey != "" {
		encoderConfig.StacktraceKey = keys.StacktraceKey
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	// 开启 level 染色
	encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level Level) zapcore.Level {
	switch level {
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case DebugLevel:
		return zapcore.DebugLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger(config configuration) *zapLogger {
	cores := make([]zapcore.Core, 0)

	logger := &zapLogger{cfg: config}

	if config.ConsoleStdoutEnable {
		lvl := zap.NewAtomicLevel()
		lvl.SetLevel(getZapLevel(config.ConsoleStdoutLevel))
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleStdoutIsJSONFormat, config.CoverDefaultKey), writer, lvl)
		cores = append(cores, core)
		logger.consoleAtomicLevel = &lvl
	}

	if config.FileStdoutEnable {
		lvl := zap.NewAtomicLevel()
		lvl.SetLevel(getZapLevel(config.FileStdoutLevel))
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileStdoutFileLocation,
			MaxSize:  config.FileStdoutLogMaxSize,
			Compress: config.FileStdoutCompress,
			MaxAge:   config.FileStdoutLogMaxAge,
		})
		core := zapcore.NewCore(getEncoder(config.FileStdoutIsJSONFormat, config.CoverDefaultKey), writer, lvl)
		cores = append(cores, core)
		logger.fileAtomicLevel = &lvl
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger.sugaredLogger = zap.New(combinedCore,
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCaller(), zap.AddCallerSkip(1),
	).Sugar()
	// 添加全局字段
	fields := make([]interface{}, 0, 2*len(config.GlobalField))
	for k, v := range config.GlobalField {
		fields = append(fields, k, v)
	}
	return logger.withFields(fields)
}

func newConsoleLogger() *zapLogger {
	var core zapcore.Core
	config := configuration{
		ConsoleStdoutEnable:       true,
		ConsoleStdoutIsJSONFormat: false,
		ConsoleStdoutLevel:        InfoLevel,
		FileStdoutEnable:          false,
	}
	logger := &zapLogger{cfg: config}

	if config.ConsoleStdoutEnable {
		lvl := zap.NewAtomicLevel()
		lvl.SetLevel(getZapLevel(config.ConsoleStdoutLevel))
		writer := zapcore.Lock(os.Stdout)
		core = zapcore.NewCore(getEncoder(config.ConsoleStdoutIsJSONFormat, config.CoverDefaultKey), writer, lvl)
		logger.consoleAtomicLevel = &lvl
	}

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger.sugaredLogger = zap.New(core,
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCaller(), zap.AddCallerSkip(1),
	).Sugar()
	return logger
}

// Debugf 打印用于调试的格式化信息，日志 level 为 DebugLevel 时才会打印
func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

// Debug 打印用于调试的信息，日志 level 为 DebugLevel 时才会打印
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

// Infof 打印格式化标准信息，默认的日志信息
func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

// Info 打印标准信息，默认的日志信息
func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

// Warnf 打印比 Info 更重要，但暂不需要单独人为介入的格式化信息
func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

// Warn 打印比 Info 更重要，但暂不需要单独人为介入的信息
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

// Errorf 打印格式化错误信息，通常需要人为介入
func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

// Error 打印错误信息，通常需要人为介入
func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

// Panicf 打印格式化错误信息，然后 panics
func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Panicf(format, args...)
}

// Panic 打印错误信息，然后 panics
func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

// Fatalf 打印错误信息，然后调用 os.Exit(1), 业务不使用
func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

// Fatal 打印错误信息，然后调用 os.Exit(1), 业务不使用
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

// InfoWithField 打印带自定义字段的日志
func (l *zapLogger) InfoWithField(fields Fields, args ...interface{}) {
	if fields == nil {
		l.Info(args...)
		return
	}
	newFields := make([]interface{}, 0, 2*len(fields))
	for key, val := range fields {
		newFields = append(newFields, key, val)
	}
	l.withFields(newFields).sugaredLogger.Info(args...)
}

func (l *zapLogger) IsDebug() bool {
	if l.cfg.FileStdoutEnable {
		return l.fileAtomicLevel.Level() == getZapLevel(DebugLevel)
	}
	return l.consoleAtomicLevel.Level() == getZapLevel(DebugLevel)
}

func (l *zapLogger) ChangeFileStdoutLevel(level Level) {
	if l.cfg.FileStdoutEnable {
		l.fileAtomicLevel.SetLevel(getZapLevel(level))
	}
}
func (l *zapLogger) ChangeConsoleStdoutLevel(level Level) {
	if l.cfg.ConsoleStdoutEnable {
		l.consoleAtomicLevel.SetLevel(getZapLevel(level))
	}
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	fields := make([]interface{}, 0, 2)
	if traceID := ctx.Value(TrackKey); traceID != nil {
		fields = append(fields, string(TrackKey), traceID)
	}
	if len(fields) == 0 {
		return l
	}
	return l.withFields(fields)
}

func (l *zapLogger) Time(sec float64) *zapLogger {
	return l.withField(TimeKey, sec)
}

func (l *zapLogger) clone() *zapLogger {
	tmp := *l
	return &tmp
}

func (l *zapLogger) withFields(fields []interface{}) *zapLogger {
	z := l.clone()
	z.sugaredLogger = l.sugaredLogger.With(fields...)
	return z
}

func (l *zapLogger) withField(key interface{}, value interface{}) *zapLogger {
	z := l.clone()
	z.sugaredLogger = l.sugaredLogger.With(key, value)
	return z
}
