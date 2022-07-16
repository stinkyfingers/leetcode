package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func subsets(nums []int) [][]int {
	output := [][]int{}
	f := func(arr []int) {
		output = append(output, arr)
	}
	
	getSubsets([]int{}, nums, f)
	return output
}

func getSubsets(current, nums []int, f func([]int)) {
	f(current)
	for i, num := range nums {
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		newCurrent = append(newCurrent, num)
		getSubsets(newCurrent, nums[i+1:], f)
	}
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		desc string
		nums []int
		expected [][]int
	}{
		{
			desc: "1",
			nums: []int{1,2,3},
			expected: [][]int{{},{1},{2},{3},{1,2},{1,3},{2,3},{1,2,3}},
		},
		{
			desc: "2",
			nums: []int{0},
			expected: [][]int{{},{0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			res := subsets(tt.nums)
			require.ElementsMatch(t, tt.expected, res)
		})
	}
}
