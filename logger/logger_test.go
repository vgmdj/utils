package logger

import (
	"testing"
	"time"
)

type TestCase1 struct {
	ID   string
	Name string
	Msg  string
}

func TestLogger(t *testing.T) {
	SetAsync()
	SetLevel(LevelDebug)
	SetLogFuncCall(true)


	Info("this is info")
	Debug("this is debug")
	Error("this is error")
	Warning("this is warning")

	case1 := TestCase1{
		ID:   "123",
		Name: "name",
		Msg:  "this is case1 msg",
	}

	Info(case1)
	Debug(case1)
	Error(case1)
	Warning(case1)

	time.Sleep(time.Second)
}
