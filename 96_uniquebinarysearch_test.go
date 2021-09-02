package leetcode

import (
	"reflect"
	"testing"
)

// numTrees calcs catalan number
// 0 => 1
// 1 => 1
// 2 => 5
// https://www.youtube.com/watch?v=YDf982Lb84o
func numTrees(n int) int {
	nums := map[int]int{0: 1, 1: 1}
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			nums[i] += nums[j] * nums[i-j-1] // increase map value at i with previous values up to that point
		}
	}
	return nums[n]
}

/* testing */

func TestNumTrees(t *testing.T) {
	tests := []struct {
		nums     int
		expected int
	}{
		{
			nums:     3,
			expected: 5,
		},
		{
			nums:     1,
			expected: 1,
		},
		{
			nums:     4,
			expected: 14,
		},
		{
			nums:     5,
			expected: 42,
		},
		{
			nums:     0,
			expected: 1,
		},
	}
	for _, test := range tests {
		res := numTrees(test.nums)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
