package leetcode

import (
	"reflect"
	"testing"
)

func searchRange(nums []int, target int) []int {
	start := beginning(nums, target, []int{0, len(nums) - 1})
	end := end(nums, target, []int{0, len(nums) - 1})
	return []int{start, end}
}

func beginning(nums []int, target int, indices []int) int {
	if indices[1] < indices[0] {
		return -1
	}
	midIndex := len(nums[indices[0]:indices[1]])/2 + indices[0]
	if nums[midIndex] < target {
		return beginning(nums, target, []int{midIndex + 1, indices[1]})
	}
	if nums[midIndex] > target {
		return beginning(nums, target, []int{indices[0], midIndex - 1})
	}
	if nums[midIndex] == target {
		if midIndex == 0 || nums[midIndex-1] < target {
			return midIndex
		}
		return beginning(nums, target, []int{indices[0], midIndex - 1})
	}
	return indices[1]
}

func end(nums []int, target int, indices []int) int {
	if indices[1] < indices[0] {
		return -1
	}
	midIndex := len(nums[indices[0]:indices[1]])/2 + indices[0]
	if nums[midIndex] < target {
		return end(nums, target, []int{midIndex + 1, indices[1]})
	}
	if nums[midIndex] > target {
		return end(nums, target, []int{indices[0], midIndex - 1})
	}
	if nums[midIndex] == target {
		if midIndex == len(nums)-1 || nums[midIndex+1] > target {
			return midIndex
		}
		return end(nums, target, []int{midIndex + 1, indices[1]})
	}
	return indices[0]
}

/* testing */

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			target:   0,
			expected: []int{0, 14},
		},
	}
	for _, test := range tests {
		res := searchRange(test.nums, test.target)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
