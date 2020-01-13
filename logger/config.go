package logger

import (
	"os"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	timeKey       = "time"
	levelKey      = "level"
	nameKey       = "logger"
	callerKey     = "caller"
	msgKey        = "msg"
	stacktraceKey = "stacktrace"
)

var (
	DefaultConfig = &Config{
		WriteToFile:    false,
		WriteToConsole: true,
		FileName:       "",
		MaxSize:        0,
		MaxBackups:     0,
		MaxDays:        0,
		TimeKey:        timeKey,
		LevelKey:       levelKey,
		NameKey:        nameKey,
		CallerKey:      callerKey,
		MessageKey:     msgKey,
		StacktraceKey:  stacktraceKey,
	}
)

type Config struct {
	WriteToFile    bool
	WriteToConsole bool
	FileName       string
	MaxDays        int
	MaxSize        int
	MaxBackups     int

	TimeKey       string
	LevelKey      string
	NameKey       string
	CallerKey     string
	MessageKey    string
	StacktraceKey string
}

func (c *Config) NewAtomicLevel() {

}

func (c *Config) NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        c.TimeKey,
		LevelKey:       c.LevelKey,
		NameKey:        c.NameKey,
		CallerKey:      c.CallerKey,
		MessageKey:     c.MessageKey,
		StacktraceKey:  c.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //lower case encoding
		EncodeTime:     EchoTimeEncoder,               //time format, eg:2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

}

func (c *Config) SetWriter() zapcore.WriteSyncer {
	result := make([]zapcore.WriteSyncer, 0)
	if c.WriteToFile {
		result = append(result, setFileWriter(c.FileName, c.MaxSize, c.MaxBackups, c.MaxDays))
	}

	if c.WriteToConsole {
		result = append(result, zapcore.AddSync(os.Stdout))
	}

	return zapcore.NewMultiWriteSyncer(result...)
}

func setFileWriter(fileName string, maxSize, maxBackups, maxDays int) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxDays,
		LocalTime:  true,
	}

	return zapcore.AddSync(hook)
}

func EchoTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
