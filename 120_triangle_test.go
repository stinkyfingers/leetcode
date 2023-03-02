package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
2
3 4
6 5 7
4 1 8 3

2
3 4
6 7 3
4 5 4 5
2-4-3-4
*/

// bottom up
// for each node in the next-to-last row, assign it that value plus the lesser of the last rows's i or i+1 values
func minimumTotal(triangle [][]int) int {
	for len(triangle) > 1 {
		row := triangle[len(triangle)-1]
		prev := triangle[len(triangle)-2]
		triangle = triangle[:len(triangle)-1]

		for i := 0; i < len(row)-1; i++ {
			lower := row[i]
			if row[i+1] < row[i] {
				lower = row[i+1]
			}
			prev[i] = lower + prev[i]
		}

		triangle[len(triangle)-1] = prev
	}
	return triangle[0][0]
}

/* testing */

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		triangle [][]int
		expected int
	}{
		{
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expected: 11,
		},
		{
			triangle: [][]int{{2}, {3, 4}, {6, 7, 3}, {4, 5, 4, 5}},
			expected: 13,
		},
		{
			triangle: [][]int{{-1}},
			expected: -1,
		},
		{
			triangle: [][]int{{-1}, {-2, -3}},
			expected: -4,
		},
	}
	for _, test := range tests {
		res := minimumTotal(test.triangle)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
