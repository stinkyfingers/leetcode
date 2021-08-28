package leetcode

import (
	"reflect"
	"testing"
)

func canJump(nums []int) bool {
	can := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 || i == len(nums)-1 {
			continue
		}
		if !surpass(nums[:i]) {
			return false
		}
	}
	return can
}

func surpass(nums []int) bool {
	steps := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > steps {
			return true
		}
		steps++
	}
	return false
}

/* testing */

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			nums:     []int{3},
			expected: true,
		},
		{
			nums:     []int{0},
			expected: true,
		},
		{
			nums:     []int{0, 1},
			expected: false,
		},
		{
			nums:     []int{1, 0},
			expected: true,
		},
	}
	for _, test := range tests {
		res := canJump(test.nums)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
