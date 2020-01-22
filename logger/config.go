package logger

import (
	"os"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	TimeKey           = "time"
	LevelKey          = "level"
	NameKey           = "logger"
	CallerKey         = "caller"
	MsgKey            = "msg"
	StacktraceKey     = "stacktrace"
	DefaultFileName   = "/tmp/logger.log"
	DefaultMaxSize    = 10 //10MB
	DefaultMaxBackups = 10 //10个备份文件
	DefaultMaxDays    = 30 //保留30天
)

var (
	DefaultConfig = &Config{
		WriteToFile:    false,
		WriteToConsole: true,
		FileName:       DefaultFileName,
		MaxSize:        DefaultMaxSize,
		MaxBackups:     DefaultMaxBackups,
		MaxDays:        DefaultMaxDays,
		TimeKey:        TimeKey,
		LevelKey:       LevelKey,
		NameKey:        NameKey,
		CallerKey:      CallerKey,
		MessageKey:     MsgKey,
		StacktraceKey:  StacktraceKey,
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

// Init default out is console
func (c *Config) Init() {
	if !c.WriteToConsole && !c.WriteToFile {
		c.WriteToConsole = true
	}

	if c.MaxDays == 0 {
		c.MaxDays = DefaultConfig.MaxDays
	}

	if c.MaxSize == 0 {
		c.MaxSize = DefaultConfig.MaxSize
	}

	if c.MaxBackups == 0 {
		c.MaxBackups = DefaultConfig.MaxBackups
	}

	if c.WriteToFile && c.FileName == "" {
		c.FileName = DefaultConfig.FileName
	}

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
