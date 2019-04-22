package convert

import (
	"testing"
)

func Test_ConvStruct2Map(t *testing.T) {
	source := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "lily",
		Age:  30,
	}
	t.Log(ConvStruct2Map(&source))
}
