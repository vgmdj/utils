package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l *Logger
)

type Logger struct {
	logger *zap.Logger
	c      *Config
}

func init() {
	c := DefaultConfig
	l = &Logger{
		c: c,
	}
	l.logger = NewLogger(l.c)
	defer l.logger.Sync()
}

// NewLogger create a new zap logger
func NewLogger(c *Config) *zap.Logger {
	c.Init()
	zc := c.NewEncoderConfig()
	ops := c.SetWriter()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zc), ops, zap.NewAtomicLevel())
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.PanicLevel), zap.AddStacktrace(zap.ErrorLevel))

}

// Reset reset the global logger
func Reset(c *Config) {
	l = &Logger{
		c: c,
	}
	l.logger = NewLogger(l.c)
	defer l.logger.Sync()
}

// With add global logger field
func With(key string, value interface{}) {
	l.logger.With(zap.Any(key, value))
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	l.logger.Panic(fmt.Sprint(args...))
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(template, args...))
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(template, args...))
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(template, args...))
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	l.logger.Panic(fmt.Sprintf(template, args...))
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(template, args...))
}

// Infowf info with fields
func Infowf(msg string, field ...zap.Field) {
	l.logger.Info(msg, field...)
}

// Warnwf warn with field
func Warnwf(msg string, field ...zap.Field) {
	l.logger.Warn(msg, field...)
}

// Panicwf panic with field
func Panicwf(msg string, field ...zap.Field) {
	l.logger.Panic(msg, field...)
}

// patal fatal with field
func Fatalwf(msg string, field ...zap.Field) {
	l.logger.Fatal(msg, field...)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return l.logger.Sync()
}
