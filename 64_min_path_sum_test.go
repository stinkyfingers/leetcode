package leetcode

import (
	"reflect"
	"testing"
)

/*
	The "sums" 2d array cells each contain the total shortest route to that cell.
	We can populate the top and left rows by adding the current grid cell value to the
	previous top or left, respectively. We calculate the remainder of the "sums" cells by adding
	the current "grid" cell value to the lesser of the upper or left "sums" cell value. The bottom
	right "sums" cell is the shortest.
*/

func minPathSum(grid [][]int) int {
	// check for no grid
	if len(grid) == 0 {
		return 0
	}

	// init array
	sums := make([][]int, len(grid))
	for i := range sums {
		sums[i] = make([]int, len(grid[0]))
	}

	// populate top row
	for i := range grid[0] {
		prev := 0
		if i > 0 {
			prev = sums[0][i-1]
		}
		sums[0][i] = grid[0][i] + prev
	}

	// populate left column
	for i := range grid {
		prev := 0
		if i > 0 {
			prev = sums[i-1][0]
		}
		sums[i][0] = grid[i][0] + prev
	}

	// dynamically program - enter sum value in each sum[][] cell
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			smaller := sums[i-1][j]
			if sums[i][j-1] < smaller {
				smaller = sums[i][j-1]
			}
			sums[i][j] = smaller + grid[i][j]
		}
	}

	// bottom right = answer
	return sums[len(grid)-1][len(grid[0])-1]
}

/* testing */

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		{
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
		{
			grid:     [][]int{{1}, {1}},
			expected: 2,
		},
		{
			grid:     [][]int{{1}},
			expected: 1,
		},
	}
	for _, test := range tests {
		res := minPathSum(test.grid)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
