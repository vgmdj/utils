package config

import (
	"fmt"
	"reflect"
)

const (
	TagFlag = "check"
)

func Check(c interface{}) (nullList []string, err error) {
	v := reflect.ValueOf(c)
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		tv := vt.Field(i).Tag.Get(TagFlag)
		if tv == "no" {
			continue
		}

		if v.Field(i).Kind() == reflect.Struct {
			l, _ := Check(v.Field(i).Interface())
			nullList = append(nullList, l...)
		}

		if v.Field(i).Kind() != reflect.Bool && v.Field(i).IsZero() {
			nullList = append(nullList, vt.Field(i).Name)
		}

	}

	if len(nullList) != 0 {
		return nullList, fmt.Errorf("%v can not be null", nullList)
	}

	return

}
