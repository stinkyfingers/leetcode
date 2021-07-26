package leetcode

import (
	"reflect"
	"testing"
)

// // NOTE: recursive trie - correct, but time limit exceeded
// // no need to examine every path
// func jump(nums []int) int {
// 	lo := len(nums) - 1
// 	cb := func(num int) {
// 		if num < lo {
// 			lo = num
// 		}
// 	}
// 	hop(nums, 1, cb)
// 	return lo
// }
//
// func hop(nums []int, hops int, cb func(num int)) {
// 	rightBound := nums[0] + 1
// 	if rightBound > len(nums) {
// 		rightBound = len(nums)
// 	}
// 	for i, _ := range nums[1:rightBound] {
// 		// fmt.Println("cur", nums[0], i, value, nums[i+1:], "hops", hops)
// 		if len(nums[i+1:]) == 1 {
// 			cb(hops)
// 			return
// 		}
//
// 		hop(nums[i+1:], hops+1, cb)
// 	}
// }

// biggest in sub-arry
func jump(nums []int) int {
	hops := 0
	if len(nums) < 2 {
		return hops
	}
	index := 0
	for {
		if index >= len(nums)-1 {
			break
		}
		rightIndex := index + 1 + nums[index]
		if rightIndex > len(nums) {
			rightIndex = len(nums)
		}
		index = getFarthest(nums, index+1, rightIndex)
		hops++
	}
	return hops
}

func getFarthest(nums []int, i, j int) int {
	index := 0
	dist := 0
	for n := i; n < j; n++ {
		if n == len(nums)-1 {
			return n // at end
		}
		if nums[n]+n >= dist {
			dist = nums[n] + n
			index = n
		}
	}
	return index
}

/* testing */

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			nums:     []int{2, 3, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 1, 1, 1, 1, 1},
			expected: 5,
		},
		{
			nums:     []int{0},
			expected: 0,
		},
		{
			nums:     []int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5},
			expected: 5,
		},
	}
	for _, test := range tests {
		res := jump(test.nums)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
