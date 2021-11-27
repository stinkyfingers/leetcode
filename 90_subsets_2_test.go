package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

/*
	1 -> 2 -> 2
	  1, 12, 122
	2 -> 2
	  2, 22
	2 (skip)
*/

func subsetsWithDup(nums []int) [][]int {
	out := [][]int{}
	f := func(set []int) {
		out = append(out, set)
	}
	sort.Ints(nums)
	subsetDig(nums, []int{}, f)
	return out
}

func subsetDig(nums, current []int, f func(set []int)) {
	f(current)
	if len(nums) == 0 {
		return
	}

	for i := range nums {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// append this num
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		newCurrent = append(newCurrent, nums[i])

		// move num array forward
		newNums := nums[i+1:]

		subsetDig(newNums, newCurrent, f)
	}
}

/* testing */

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 2},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			input:    []int{1, 2},
			expected: [][]int{{}, {1}, {1, 2}, {2}},
		},
		{
			input:    []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			input:    []int{2, 1, 2},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}
	for i, test := range tests {
		res := subsetsWithDup(test.input)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
