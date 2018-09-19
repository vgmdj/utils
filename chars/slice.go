package chars

import (
	"sort"

	"github.com/vgmdj/utils/logger"
)

//IsDuplicates 数组内重复元素判断
func IsDuplicates(a []string) bool {
	counter := make(map[string]struct{})

	for _, v := range a {
		if _, ok := counter[v]; ok {
			return true
		}
		counter[v] = struct{}{}

	}

	return false
}

//RemoveDuplicatesAndEmpty 去除重复元素
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sort.Strings(a)

	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

//IsContain 是否包含元素
func IsContain(array interface{}, value interface{}) bool {
	switch array.(type) {
	default:
		logger.Error("not support value type")
		return false

	case []string:
		for _, v := range array.([]string) {
			if v == value {
				return true
			}
		}

	case []int:
		for _, v := range array.([]int) {
			if v == value {
				return true
			}
		}

	case []int64:
		for _, v := range array.([]int64) {
			if v == value {
				return true
			}
		}

	case []float64:
		for _, v := range array.([]float64) {
			if v == value {
				return true
			}
		}

	case []byte:
		for _, v := range array.([]byte) {
			if v == value {
				return true
			}
		}

	case []interface{}:
		for _, v := range array.([]interface{}) {
			if v == value {
				return true
			}
		}

	}

	return false

}
