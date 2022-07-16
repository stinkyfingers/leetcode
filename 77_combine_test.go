package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func combine(n int, k int) [][]int {
	output := [][]int{}
	f := func(arr []int) {
		output = append(output, arr)
	}
	
	start := []int{}
	for i := 1; i <= n; i++ {
		start = append(start, i)
	}
	
	getCombination([]int{}, start, k, f)	
	return output
}

func getCombination(current []int, remaining []int, k int, f func([]int)) {
	if k == 1 {
		for _, num := range remaining {
			newCurrent := make([]int, len(current))
			copy(newCurrent, current)
			f(append(newCurrent, num))
		}
		return
	}
	for i := 0; i < len(remaining) - k + 1; i++ {
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		newCurrent = append(newCurrent, remaining[i])
		getCombination(newCurrent, remaining[i+1:], k - 1, f)
	}
}

func TestCombine(t *testing.T) {
	tests := []struct {
		desc string
		n, k int
		expected [][]int
	}{
		{
			desc: "1",
			n: 4,
			k: 2,
			expected: [][]int{{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}},
		},
		{
			desc: "2",
			n: 4,
			k: 3,
			expected: [][]int{{1,2,3},{1,2,4},{1,3,4},{2,3,4}},
		},
		{
			desc: "3",
			n: 1,
			k: 1,
			expected: [][]int{{1}},
		},
		{
			desc: "4",
			n: 3,
			k: 1,
			expected: [][]int{{1},{2},{3}},
		},
		{
			desc: "5",
			n: 3,
			k: 3,
			expected: [][]int{{1,2,3}},
		},
		{
			desc: "6",
			n: 5,
			k: 4,
			expected: [][]int{{1,2,3,4},{1,2,3,5},{1,2,4,5},{1,3,4,5},{2,3,4,5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			res := combine(tt.n, tt.k)
			require.Equal(t, tt.expected, res)
		})
	}
}
