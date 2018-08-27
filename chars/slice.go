package chars

import "sort"

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
func IsContain(array []interface{}, value interface{}) bool {
	if len(array) == 0 {
		return false
	}

	for _, v := range array {
		if value == v {
			return true
		}
	}

	return false
}
