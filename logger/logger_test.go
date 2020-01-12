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
	Info("this is info")
	Error("this is error")
	Warn("this is warning")

	case1 := TestCase1{
		ID:   "123",
		Name: "name",
		Msg:  "this is case1 msg",
	}

	Info(case1)
	Warn(case1)
	Error(case1)

	time.Sleep(time.Second)
}
