package chars

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRestTime(t *testing.T) {
	ast := assert.New(t)

	rt := RestTime(22, 50, 0, 50, true)
	t1 := time.Date(2018, time.July, 1, 23, 55, 0, 0, time.UTC)
	ast.Equal(rt.IsWorkingTime(t1), false)
	ast.Equal(rt.IsRestTime(t1), true)
	ast.Equal(rt.WaitTime(t1),time.Minute*56)

	t2 := time.Date(2018, time.July, 1, 0, 50, 0, 0, time.UTC)
	ast.Equal(rt.IsWorkingTime(t2), false)
	ast.Equal(rt.IsRestTime(t2), true)
	ast.Equal(rt.WaitTime(t2),time.Minute)

	rt.SetRestTime(8, 10, 14, 0, false)
	t3 := time.Date(2018, time.July, 1, 12, 55, 0, 0, time.UTC)
	ast.Equal(rt.IsWorkingTime(t3), false)
	ast.Equal(rt.IsRestTime(t3), true)
	ast.Equal(rt.WaitTime(t3),time.Minute*66)

	t4 := time.Date(2018, time.July, 1, 14, 1, 0, 0, time.UTC)
	ast.Equal(rt.IsWorkingTime(t4), true)
	ast.Equal(rt.IsRestTime(t4), false)
	ast.Equal(rt.WaitTime(t4),time.Duration(0))

}
