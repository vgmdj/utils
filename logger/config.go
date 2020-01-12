package logger

import (
	"time"

	"go.uber.org/zap"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
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
		RotationCount:  7,
		RotationTime:   time.Hour * 24,
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
	RotationCount  uint
	RotationTime   time.Duration

	TimeKey       string
	LevelKey      string
	NameKey       string
	CallerKey     string
	MessageKey    string
	StacktraceKey string
}

func (c *Config) NewZapConfig() zap.Config {
	ec := zapcore.EncoderConfig{
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
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	zc := zap.NewProductionConfig()
	zc.EncoderConfig = ec

	return zc
}

func (c *Config) SetWriter() []zap.Option {
	result := make([]zap.Option, 0)
	if c.WriteToFile {
		result = append(result, setFileWriter(c.FileName, c.RotationCount, c.RotationTime))
	}

	if c.WriteToConsole {

	}

	return result
}

func setFileWriter(fileName string, rotateCount uint, rotateTime time.Duration) zap.Option {
	hook, _ := rotateLogs.New(
		fileName+".%Y%m%d",
		rotateLogs.WithLinkName(fileName),
		rotateLogs.WithRotationCount(rotateCount),
		rotateLogs.WithRotationTime(rotateTime),
	)

	return zap.ErrorOutput(zapcore.AddSync(hook))
}

func EchoTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
