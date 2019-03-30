package slice

import (
	"fmt"
	"reflect"
)

//InArray haystack need slice or array, needle.kind should be int,float,string
func InArray(needle interface{}, haystack interface{}) (bool, error) {
	ht := reflect.TypeOf(haystack)
	if ht.Kind() != reflect.Slice && ht.Kind() != reflect.Array {
		return false, fmt.Errorf("haystack should be array or slice, type %s", ht.Kind().String())
	}
	hv := reflect.ValueOf(haystack)
	nt := reflect.TypeOf(needle)
	nv := reflect.ValueOf(needle)
	switch nt.Kind() {
	case reflect.String:
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
	case reflect.Float32, reflect.Float64:
	default:
		return false, fmt.Errorf("needle type %s error, should by int,float,string", nt.Kind().String())
	}
	for i := 0; i < hv.Len(); i++ {
		if hv.Index(i).Kind() != nt.Kind() {
			return false, fmt.Errorf("needle type should be match haystack, needle:%s, haystack:%s", nt.Kind().String(), hv.Index(i).Kind().String())
		}
		if hv.Index(i).String() == nv.String() {
			return true, nil
		}
	}
	return false, nil
}
