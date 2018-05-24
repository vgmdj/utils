package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSql(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(checkOp("="), true)
	ast.Equal(checkOp("!"), false)

}
