package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// Logger 全局日志实例
	Logger = logrus.New()
)

// Init 初始化日志系统
func Init() {
	// 设置日志格式为JSON格式
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置日志输出到标准错误输出
	Logger.SetOutput(os.Stderr)

	// 根据环境变量设置日志级别
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

// Debug 记录Debug级别日志
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Debugf 记录Debug级别格式化日志
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Info 记录Info级别日志
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Infof 记录Info级别格式化日志
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Warn 记录Warn级别日志
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warnf 记录Warn级别格式化日志
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Error 记录Error级别日志
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Errorf 记录Error级别格式化日志
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Fatal 记录Fatal级别日志并退出程序
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf 记录Fatal级别格式化日志并退出程序
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// WithError 添加错误上下文
func WithError(err error) *logrus.Entry {
	return Logger.WithError(err)
}

// WithField 添加单个字段
func WithField(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

// WithFields 添加多个字段
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

// SetLevel 设置日志级别
func SetLevel(level logrus.Level) {
	Logger.SetLevel(level)
}

// init 包初始化时自动初始化日志
func init() {
	Init()
}
