package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/*
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1
*/

func generate(numRows int) [][]int {
	var output [][]int
	for i := 0; i < numRows; i++ {
		rowLen := i + 1
		row := make([]int, rowLen)
		for j := 0; j < rowLen; j++ {
			if j == 0 || j == rowLen-1 {
				row[j] = 1
			} else {
				row[j] = output[i-1][j] + output[i-1][j-1]
			}
		}
		output = append(output, row)
	}
	return output
}

/* testing */

func TestGeneratePascal(t *testing.T) {
	tests := []struct {
		numRows  int
		expected [][]int
	}{
		{
			numRows:  5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			numRows:  1,
			expected: [][]int{{1}},
		},
	}
	for _, test := range tests {
		res := generate(test.numRows)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
