package leetcode

import (
	"reflect"
	"testing"
)

func removeDuplicates80(nums []int) int {
	m := 0
	for i := range nums {
		nums[m] = nums[i]
		if i < 2 {
			m++
			continue
		}

		if nums[i] != nums[m-1] || nums[i] != nums[m-2] {
			m++
		}
	}
	return m
}

/* testing */

func Test80RemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums          []int
		expected      int
		expectedSlice []int
	}{
		{
			nums:          []int{1, 1, 1, 2, 2, 3},
			expected:      5,
			expectedSlice: []int{1, 1, 2, 2, 3, 3},
		},
		{
			nums:          []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected:      7,
			expectedSlice: []int{0, 0, 1, 1, 2, 3, 3, 3, 3},
		},
		{
			nums:          []int{1, 1, 1, 2, 2, 3},
			expected:      5,
			expectedSlice: []int{1, 1, 2, 2, 3, 3},
		},
		{
			nums:          []int{1, 1, 1},
			expected:      2,
			expectedSlice: []int{1, 1, 1},
		},
		{
			nums:          []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			expected:      6,
			expectedSlice: []int{1, 1, 2, 2, 3, 3, 3, 3, 3},
		},
	}
	for _, test := range tests {
		res := removeDuplicates80(test.nums)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
		for i := range test.expectedSlice {
			if i > test.expected {
				break
			}
			if test.expectedSlice[i] != test.nums[i] {
				t.Errorf("error at position %d: %d, %d", i, test.expectedSlice[i], test.nums[i])
			}
		}
	}
}
