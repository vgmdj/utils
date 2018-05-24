package chars

import (
	"math"
)

func BytesToInt(b []byte) int {
	bti := 0
	for k, v := range b {
		bti += (int(v) - 48) * int(math.Pow10(len(b)-k-1))
	}

	return bti
}

func TakeLeftInt(num int, n int) int {
	for num >= int(math.Pow10(n)) {
		num /= 10
	}

	return num
}

//TODO finish this function
func TekeRightInt(num int, n int) int {

	return 0
}
