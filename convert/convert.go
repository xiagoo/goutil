package convert

import (
	"reflect"
)

func ConvStruct2Map(source interface{}) map[string]interface{} {
	target := make(map[string]interface{})
	st := reflect.TypeOf(source)
	sv := reflect.ValueOf(source)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		sv = sv.Elem()
	}
	if st.Kind() == reflect.Struct {
		for i := 0; i < sv.Type().NumField(); i++ {
			target[sv.Type().Field(i).Tag.Get("json")] = sv.Field(i).Interface()
		}
	}
	return target
}
