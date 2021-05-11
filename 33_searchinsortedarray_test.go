package leetcode

import (
	"testing"
)

func search(nums []int, target int) int {
	hi := len(nums) - 1
	lo := 0

	for {
		mid := (hi-lo)/2 + lo
		if nums[mid] == target {
			return mid
		}
		if nums[hi] == target {
			return hi
		}
		if nums[lo] == target {
			return lo
		}

		if hi <= lo {
			break
		}

		if (nums[lo] < nums[mid] && target > nums[lo] && target < nums[mid]) || (nums[lo] > nums[mid] && (target >= nums[lo] || target <= nums[mid])) {
			hi = mid
			continue
		}
		lo = mid + 1
	}

	return -1
}

/* testing */

func TestSearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			nums:     []int{1, 3},
			target:   3,
			expected: 1,
		},
		{
			nums:     []int{5, 1, 3},
			target:   5,
			expected: 0,
		},
		{
			nums:     []int{4, 5, 6, 7, 8, 1, 2, 3},
			target:   8,
			expected: 4,
		},
	}
	for _, test := range tests {
		res := search(test.nums, test.target)
		if res != test.expected {
			t.Errorf("got %d, expected %d", res, test.expected)
		}
	}
}
