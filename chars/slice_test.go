package chars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsContain(t *testing.T) {
	ast := assert.New(t)

	var array1 []interface{}
	array1 = append(array1,"a")
	array1 = append(array1,1)

	t1 := "a"
	ast.Equal(true,IsContain(array1,t1))

	t2 := 1.23
	ast.Equal(false,IsContain(array1,t2))


	array2 := []string{"1","2","a","b"}
	t3 := 1
	ast.Equal(false,IsContain(array2,t3))

	t4 := "1"
	ast.Equal(true,IsContain(array2,t4))

}
