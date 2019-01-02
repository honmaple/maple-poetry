/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: serializer.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-12 15:04:57 (CST)
 Last Update: Tuesday 2018-12-25 13:35:07 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

import (
	"reflect"
)

// SerializerType ..
type SerializerType interface {
	Serializer() map[string]interface{}
}

// Serializer ..
func Serializer(instance interface{}) interface{} {
	switch reflect.TypeOf(instance).Kind() {
	case reflect.Slice:
		v := reflect.ValueOf(instance)
		ins := make([]map[string]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			ins[i] = v.Index(i).Interface().(SerializerType).Serializer()
		}
		return ins
	default:
		return instance.(SerializerType).Serializer()
	}
}
