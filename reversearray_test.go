package leetcode

import (
	"reflect"
	"testing"
)

func reverseArraySwap(arr []int) []int {
	return reverseFunc(arr, 0)
}

func reverseFunc(arr []int, pos int) []int {
	if pos >= len(arr)/2 {
		return arr
	}
	arr[0+pos], arr[len(arr)-1-pos] = arr[len(arr)-1-pos], arr[0+pos]
	reverseFunc(arr, pos+1)
	return arr
}

func reverseArrayAppend(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	first := arr[len(arr)-1]
	arr = append([]int{first}, reverseArrayAppend(arr[:len(arr)-1])...)
	return arr
}

func TestReverseArray(t *testing.T) {
	tests := []struct {
		x   []int
		exp []int
	}{
		{
			x:   []int{1, 2, 3},
			exp: []int{3, 2, 1},
		},
		{
			x:   []int{1, 2, 3, 4, 5},
			exp: []int{5, 4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		res := reverseArrayAppend(test.x)
		if !reflect.DeepEqual(res, test.exp) {
			t.Errorf("expected %d, got %d", test.exp, res)
		}

		res = reverseArraySwap(test.x)
		if !reflect.DeepEqual(res, test.exp) {
			t.Errorf("expected %d, got %d", test.exp, res)
		}
	}
}
