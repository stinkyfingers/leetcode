package leetcode

import (
	"reflect"
	"testing"
)

/*
	Dynamic Programming Version
*/

// dyanmic programming
// top row and left column: each cell has only 1 way to get there
// other cells: add together upper and left values
func uniquePaths(m int, n int) int {
	if n < 1 {
		return 0
	}
	// initialize with 1 for first row and column
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := range matrix {
		matrix[i][0] = 1
	}
	for i := range matrix[0] {
		matrix[0][i] = 1
	}
	// calc value of other cells
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[m-1][n-1]
}

/* MATH VERSION */
// math - paths = m+n-2! / m-1! / n-1!
func uniquePathsMATH(m int, n int) int {
	return factorial(m+n-2) / factorial(m-1) / factorial(n-1)
}

func factorial(x int) int {
	total := x
	for i := x - 1; i > 1; i-- {
		total *= i
	}
	return total
}

/* testing */

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n     int
		expected int
	}{
		{
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			m:        2,
			n:        3,
			expected: 3,
		},
		{
			m:        2,
			n:        4,
			expected: 4,
		},
		{
			m:        51,
			n:        9,
			expected: 1916797311,
		},
		{
			m:        10,
			n:        10,
			expected: 48620,
		},
		{
			m:        1,
			n:        0,
			expected: 0,
		},
		{
			m:        1,
			n:        1,
			expected: 1,
		},
	}
	for _, test := range tests {
		res := uniquePaths(test.m, test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
