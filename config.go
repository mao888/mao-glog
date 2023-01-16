package glog

// DefaultConfiguration 默认配置
var defaultConfiguration = configuration{
	ConsoleStdoutEnable:       false,
	ConsoleStdoutIsJSONFormat: false,
	ConsoleStdoutLevel:        DebugLevel,
	FileStdoutEnable:          true,
	FileStdoutIsJSONFormat:    true,
	FileStdoutLevel:           DebugLevel,
	FileStdoutFileLocation:    "./log.log",
	FileStdoutLogMaxSize:      256,
	FileStdoutCompress:        true,
	FileStdoutLogMaxAge:       30,
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type configuration struct {
	//  是否开启控制台日志输出
	ConsoleStdoutEnable bool
	//  控制台日志是否是 JSON 格式
	ConsoleStdoutIsJSONFormat bool
	// Level 控制台日志等级
	ConsoleStdoutLevel Level

	// 是否开启文件日志输出
	FileStdoutEnable bool
	// 文件日志是否是 JSON 格式
	FileStdoutIsJSONFormat bool
	// Level 文件日志等级
	FileStdoutLevel Level
	// 写入文件位置
	FileStdoutFileLocation string
	FileStdoutLogMaxSize   int
	FileStdoutCompress     bool
	FileStdoutLogMaxAge    int

	//全局字段
	GlobalField Fields

	//keys
	CoverDefaultKey CoverDefaultKey
}

type Option interface {
	apply(*configuration)
}

type funcOption struct {
	f func(*configuration)
}

func (f *funcOption) apply(conf *configuration) {
	f.f(conf)
}

func newFuncOption(f func(*configuration)) *funcOption {
	return &funcOption{f: f}
}

// WithConsoleStdout 打开控制台输出，生产环境不建议使用
func WithConsoleStdout() Option {
	return newFuncOption(func(c *configuration) {
		c.ConsoleStdoutEnable = true
	})
}

//Deprecated WithConsoleLevel 控制台打印级别设置，使用WithLevel同时设置控制台和文件的日志级别
func WithConsoleLevel(level Level) Option {
	return newFuncOption(func(c *configuration) {
		c.ConsoleStdoutLevel = level
	})
}

// WithConsoleFormatJson 控制台打印格式设置为 json
func WithConsoleFormatJson() Option {
	return newFuncOption(func(c *configuration) {
		c.ConsoleStdoutIsJSONFormat = true
	})
}

// WithLevel 设置日志级别
func WithLevel(level Level) Option {
	return newFuncOption(func(c *configuration) {
		c.FileStdoutLevel = level
		c.ConsoleStdoutLevel = level
	})
}

// WithFileLevel 设置文件输出日志级别
func WithFileLevel(level Level) Option {
	return newFuncOption(func(c *configuration) {
		c.FileStdoutLevel = level
	})
}

// WithFileLocation 设置日志文件名称，包含路径
func WithFileLocation(filename string) Option {
	return newFuncOption(func(c *configuration) {
		c.FileStdoutFileLocation = filename
	})
}

// WithLogMaxSize 设置单个文件的大小, 单位MB
func WithLogMaxSize(maxSize int) Option {
	return newFuncOption(func(c *configuration) {
		c.FileStdoutLogMaxSize = maxSize
	})
}

// WithLogMaxAge 设置文件保存天数
func WithLogMaxAge(maxAge int) Option {
	return newFuncOption(func(c *configuration) {
		c.FileStdoutLogMaxAge = maxAge
	})
}

// WithOffCompress 设置不进行文件压缩
func WithOffCompress() Option {
	return newFuncOption(func(c *configuration) {
		c.FileStdoutCompress = false
	})
}

// WithCustomizedGlobalField 文件日志添加全局字段
// 具体项目中显式传入的日志字段，用来初始化全局logger
// 如 容器编号、环境、项目名称 等字段
func WithCustomizedGlobalField(fields Fields) Option {
	return newFuncOption(func(c *configuration) {
		c.GlobalField = fields
	})
}

type CoverDefaultKey struct {
	LevelKey      string
	TimeKey       string
	CallerKey     string
	MessageKey    string
	StacktraceKey string
}

// WithCoverDefaultKey 设置覆盖默认字段的字段名
func WithCoverDefaultKey(keys CoverDefaultKey) Option {
	return newFuncOption(func(c *configuration) {
		c.CoverDefaultKey = keys
	})
}
