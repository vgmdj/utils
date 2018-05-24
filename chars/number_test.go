package chars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTakeLeftInt(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(TakeLeftInt(1234567, 3), 123)
}
