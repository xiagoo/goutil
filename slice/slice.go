package slice

import (
	"fmt"
	"reflect"

	"github.com/xiagoo/goutil/convert"
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

func ArrayMerge(slice interface{}, others ...interface{}) (interface{}, error) {
	st := reflect.TypeOf(slice)
	sv := reflect.ValueOf(slice)
	ot := reflect.TypeOf(others)
	ov := reflect.ValueOf(others)
	if ot.Kind() == reflect.Ptr {
		ot = ot.Elem()
		ov = ov.Elem()
	}
	if ot.Kind() != reflect.Slice {
		return nil, fmt.Errorf("others type should be slice, others:%s", ot.Kind().String())
	}
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		sv = sv.Elem()
	}
	if sv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("slice type should be slice, slice:%s", st.Kind().String())
	}

	if st.Kind() != ot.Kind() {
		return nil, fmt.Errorf("slice type and others type should be same, slice:%s, others:%s", ot.Kind().String(), st.Kind().String())
	}

	for i := 0; i < ov.Len(); i++ {
		sv = reflect.AppendSlice(
			sv,
			ov.Index(i).Elem(),
		)
	}
	rv := reflect.New(sv.Type()).Elem()
	tmp := make(map[string]reflect.Value)
	for i := 0; i < sv.Len(); i++ {
		tmp[convert.ConvIntf2Str(sv.Index(i))] = sv.Index(i)
	}
	for _, v := range tmp {
		rv = reflect.Append(rv, v)
	}
	return rv, nil
}

func RemoveDuplicate(a []string) []string {
	ret := []string{}
	for i := 0; i < len(a); i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}
