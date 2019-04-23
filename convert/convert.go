package convert

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

//ConvStruct2Map convert struct to map[string]interface{}
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
			tag := sv.Type().Field(i).Tag.Get("json")
			if tag == "-" {
				continue
			}

			name, opts := parseTag(tag)
			if !isValidTag(name) {
				continue
			}

			if strings.Contains(opts, "omitempty") && isEmptyValue(sv.Field(i)) {
				continue
			}
			target[name] = sv.Field(i).Interface()
		}
	}
	return target
}

func parseTag(tag string) (string, string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tag[idx+1:]
	}
	return tag, ""
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
		case !unicode.IsLetter(c) && !unicode.IsDigit(c):
			return false
		}
	}
	return true
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

//ConvIntf2Str ...
func ConvIntf2Str(source interface{}) string {
	return fmt.Sprintf("%v", source)
}

//ConvMapIntf2MapStr ...
func ConvMapIntf2MapStr(source map[string]interface{}) map[string]string {
	target := make(map[string]string)
	for k, v := range source {
		target[k] = ConvIntf2Str(v)
	}
	return target
}
