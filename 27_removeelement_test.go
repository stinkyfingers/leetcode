package leetcode

import (
	"reflect"
	"testing"
)

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != val {
			continue
		}
		nums = append(nums[:i], nums[i+1:]...)
	}
	return len(nums)
}

/* testing */

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		n        int
		expected []int
	}{
		{
			arr:      []int{3, 2, 2, 3},
			val:      3,
			n:        2,
			expected: []int{2, 2},
		},
	}
	for _, test := range tests {
		res := removeElement(test.arr, test.val)
		if !reflect.DeepEqual(test.arr[:test.n], test.expected[:test.n]) {
			t.Errorf("got %v, expected %v", test.arr, test.expected)
		}
		if res != test.n {
			t.Errorf("got %d, expected %d", res, test.n)
		}
	}
}
