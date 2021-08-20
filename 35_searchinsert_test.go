package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	hi := len(nums)
	lo := 0
	return searchInsertRange(nums, hi, lo, target)
}

func searchInsertRange(nums []int, hi, lo, target int) int {
	mid := len(nums[lo:hi])/2 + lo

	if nums[mid] == target {
		return mid
	}
	if len(nums[lo:hi]) == 1 {
		if nums[lo] > target {
			return lo
		} else {
			return lo + 1
		}
	}
	if nums[mid] > target {
		return searchInsertRange(nums, mid, lo, target)
	}
	return searchInsertRange(nums, hi, mid, target)
}

// map[0:1 3:0 9:0 12:1]
func TestSearchInsert(t *testing.T) {
	tests := []struct {
		desc   string
		nums   []int
		target int
		exp    int
	}{
		{
			desc:   "found",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			exp:    2,
		},
		{
			desc:   "last",
			nums:   []int{1, 3, 5, 6},
			target: 6,
			exp:    3,
		},
		{
			desc:   "first",
			nums:   []int{1, 3, 5, 6},
			target: 1,
			exp:    0,
		},
		{
			desc:   "short",
			nums:   []int{1, 3},
			target: 3,
			exp:    1,
		},
		{
			desc:   "empty",
			nums:   []int{},
			target: 3,
			exp:    0,
		},
		{
			desc:   "missing",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			exp:    1,
		},
		{
			desc:   "missing end",
			nums:   []int{1, 3, 5, 6},
			target: 8,
			exp:    4,
		},
		{
			desc:   "missing beginning",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			exp:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			res := searchInsert(tt.nums, tt.target)
			require.Equal(t, tt.exp, res)
		})
	}
}
