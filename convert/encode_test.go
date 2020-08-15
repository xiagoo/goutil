package convert

import (
	"fmt"
	"testing"
)

func TestToURLEncode(t *testing.T) {
	fmt.Println(ToURLEncode("http://www.baidu.com?key=1"))
}

func TestToURLDecode(t *testing.T) {
	fmt.Println(ToURLDecode("http%3A%2F%2Fwww.baidu.com%3Fkey%3D1"))
}

func TestUnicode2S(t *testing.T) {
	fmt.Println(Unicode2S(`\u5bb6\u65cf`))
}

func TestS2Unicode(t *testing.T) {
	fmt.Println(S2Unicode("家族"))
}
