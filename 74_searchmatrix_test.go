package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Note: probably faster to use binary search (twice) but using indices is still faster than O(mn)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	var r []int
	for _, row := range matrix {
		if len(row) < 1 {
			return false
		}
		if row[0] > target {
			break 
		}
		r = row
	}
	
	for _, val := range r {
		if val == target {
			return true 
		}
		if val > target {
			return false 
		}
	}
	return false
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		desc string
		matrix [][]int
		target int 
		expected bool
	}{
		{
			desc: "not found 1",
			matrix: [][]int{{1,1}},
			target: 0,
			expected: false,
		},
		{
			desc: "not found 3",
			matrix: [][]int{{1,1}},
			target: 2,
			expected: false,
		},
		{
			desc: "not found 2",
			matrix: [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
			target: 13,
			expected: false,
		},
		{
			desc: "1",
			matrix: [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
			target: 3,
			expected: true,
		},
		{
			desc: "one row",
			matrix: [][]int{{1,3}},
			target: 1,
			expected: true,
		},
		{
			desc: "one row 2 ",
			matrix: [][]int{{1,3,5,7}},
			target: 3,
			expected: true,
		},
		{
			desc: "two rows",
			matrix: [][]int{{1},{2}},
			target: 2,
			expected: true,
		},
		{
			desc: "two cols",
			matrix: [][]int{{1,2},{3,5}},
			target: 5,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ok := searchMatrix(tt.matrix, tt.target)
			require.Equal(t, tt.expected, ok)
		})
	}
}
