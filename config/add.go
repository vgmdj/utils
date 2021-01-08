package config

import (
	"fmt"
	"reflect"

	"github.com/go-ini/ini"
)

const (
	tag = "config"
)

//AddConfig 添加配置
func (c *Conf) AddConfig(sec string, obj interface{}) (err error) {
	rv := reflect.ValueOf(obj).Elem()
	// dereference pointer
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Struct {
		// for each struct field on v
		strSec := c.GetSection(sec)

		unmarshal(strSec, rv)

	} else {
		return fmt.Errorf("v must point to a struct type")
	}

	return

}

func unmarshal(sec *ini.Section, rv reflect.Value) error {
	rType := rv.Type()

	for i := 0; i < rType.NumField(); i++ {
		key := rType.Field(i).Tag.Get(tag)
		secKey, err := sec.GetKey(key)
		if err != nil {
			return err
		}

		if rv.Field(i).CanSet() {
			rv.Field(i).Set(reflect.ValueOf(secKey.Value()))
		}

	}

}
