package chars

import "sort"

func IsDuplicates(a []string) bool {
	sort.Strings(a)

	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if i > 0 && a[i-1] == a[i] {
			return true
		}
	}
	return false
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sort.Strings(a)

	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

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
