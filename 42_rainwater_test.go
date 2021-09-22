package leetcode

import (
	"reflect"
	"testing"
)

func trap(height []int) int {
	var total int
	right := 0
	left := 0
	start, end := 0, len(height)-1
	for start < end {
		lh := height[start]
		rh := height[end]
		if lh < rh {
			if lh > left {
				left = lh
			} else {
				total += left - lh
			}
			start++
		} else {
			if rh > right {
				right = rh
			} else {
				total += right - rh
			}
			end--
		}
	}
	return total
}

/* testing */

func TestTrap(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			height:   []int{1, 0, 1},
			expected: 1,
		},
		{
			height:   []int{1, 0},
			expected: 0,
		},
		{
			height:   []int{1},
			expected: 0,
		},
		{
			height:   []int{0, 1, 1},
			expected: 0,
		},
	}
	for _, test := range tests {
		res := trap(test.height)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
