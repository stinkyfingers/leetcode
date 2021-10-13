package leetcode

import (
	"reflect"
	"testing"
)

// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9
// digits does not contain any leading 0's.

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
		if i == 0 {
			return append([]int{1}, digits...)
		}
	}
	return digits
}

/* testing */

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits, expected []int
	}{
		{
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			digits:   []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			digits:   []int{0},
			expected: []int{1},
		},
		{
			digits:   []int{3, 9},
			expected: []int{4, 0},
		},
		{
			digits:   []int{9},
			expected: []int{1, 0},
		},
	}
	for i, test := range tests {
		res := plusOne(test.digits)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
