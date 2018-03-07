package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"reflect"
)

const (
	tag = "config"
)

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

func unmarshal(sec *ini.Section, rv reflect.Value) {
	rType := rv.Type()

	for i := 0; i < rType.NumField(); i++ {
		key := rType.Field(i).Tag.Get(tag)
		secKey, err := sec.GetKey(key)
		if err != nil {
			log.Println(err.Error())
		}

		if rv.Field(i).CanSet() {
			rv.Field(i).Set(reflect.ValueOf(secKey.Value()))
		}

	}

}
