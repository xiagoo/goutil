package convert

import (
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
