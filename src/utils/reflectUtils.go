package utils

import (
	"github.com/pkg/errors"
	"log"
	"reflect"
)

func ReflectField2Map(structName interface{}) map[string]string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make(map[string]string)
	for i := 0; i < fieldNum; i++ {
		result[t.Field(i).Name] = t.Field(i).Type.String()
	}
	return result

}

func ReflectField2List(structName interface{}) (ml []interface{}, err error) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil, errors.New("Check type error not Struct")
	}
	fieldNum := t.NumField()
	result := make(map[string]string)
	for i := 0; i < fieldNum; i++ {
		result[t.Field(i).Name] = t.Field(i).Type.String()
	}
	return nil, errors.New("Check type error not Struct")

}
