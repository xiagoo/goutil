package slice

import (
	"testing"
)

func Test_InArray(t *testing.T) {
	needle := 1
	haystack := []int{1, 2, 3}
	flag, err := InArray(needle, haystack)
	if err != nil {
		t.Error(err)
	}
	t.Log(flag)

	needleFloat := 1.0
	haystackFloat := []float32{1.0, 3.0}
	flag, err = InArray(needleFloat, haystackFloat)
	if err != nil {
		t.Error(err)
	}
	t.Log(flag)

	needleString := "1.0"
	haystackString := []string{"1.0", "3.0"}
	flag, err = InArray(needleString, haystackString)
	if err != nil {
		t.Error(err)
	}
	t.Log(flag)
}

func Test_ArrayMerge(t *testing.T) {
	sliceString := []string{"1", "2", "3"}
	othersString := []string{"1", "2", "3"}
	s, err := ArrayMerge(sliceString, othersString)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)

	sliceInt := []int{1, 2, 3}
	othersInt := []int{3, 4, 5}
	s, err = ArrayMerge(sliceInt, othersInt)
	if err != nil {
		t.Error(err)
	}

	t.Log(s.([]int))
}
