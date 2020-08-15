package date

import (
	"fmt"
	"testing"
	"time"
)

func TestS2Stamp(t *testing.T) {
	fmt.Println(S2Stamp("2020-08-01", YMD))
}

func TestStamp2S(t *testing.T) {
	fmt.Println(Stamp2S(YMD,1596240000))
}

func TestToday0Clock(t *testing.T) {
	fmt.Println(Get0Clock(0))
	fmt.Println(Get0Clock(1))
	fmt.Println(Get0Clock(-1))
}

func TestTime2S(t *testing.T) {
	fmt.Println(Time2S(HM, time.Now()))
}
