package convert

import (
	"fmt"
	"testing"
)

func Test_ConvStruct2Map(t *testing.T) {
	source := struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age"`
	}{
		Name: "lily",
		Age:  30,
	}
	t.Log(ConvStruct2Map(&source))
}

func Test_ConvIntf2Str(t *testing.T) {
	t.Log(ConvIntf2Str("test"))
	t.Log(ConvIntf2Str(1))
}

func Test_ConvMapIntf2MapStr(t *testing.T) {
	t.Log(ConvMapIntf2MapStr(map[string]interface{}{
		"name": "lily",
		"age":  20,
	}))
}

func Test_ConvSlice2StrSlice(t *testing.T) {
	fmt.Printf("%#v\n", ConvSlice2StrSlice([]string{"1", "3", "4"}))
	fmt.Printf("%#v\n", ConvSlice2StrSlice([]int{1, 3, 4}))
	fmt.Printf("%#v\n", ConvSlice2StrSlice([]float64{1.1, 3.0, 4.0}))
}
