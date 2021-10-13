package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// spiral matrix, given n x n dimensions
// 1 <= n <= 20
/*
n = 3
1 2 3
8 9 4
7 6 5

n = 4
1  2  3  4
12 13 14 5
11 16 15 6
10  9  8 7
*/

func generateMatrix(n int) [][]int {
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
	}
	total := int(math.Pow(float64(n), 2.0))
	x, y := 0, 0 // position
	start := 0
	end := n - 1
	dir := "r"

	for i := 1; i <= total; i++ {
		out[x][y] = i
		switch dir {
		case "r":
			y++
			if y == end {
				dir = "d"
			}
		case "d":
			x++
			if x == end {
				dir = "l"
				end--
			}
		case "l":
			y--
			if y == start {
				dir = "u"
				start++
			}
		case "u":
			x--
			if x == start {
				dir = "r"
			}
		}
	}

	return out
}

/* testing */

func TestSpiralMatrix(t *testing.T) {
	tests := []struct {
		n        int
		expected [][]int
	}{
		{
			n:        3,
			expected: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			n:        1,
			expected: [][]int{{1}},
		},
		{
			n:        4,
			expected: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
		},
	}
	for i, test := range tests {
		res := generateMatrix(test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
