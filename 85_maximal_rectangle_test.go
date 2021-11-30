package leetcode

import (
	"reflect"
	"testing"
)

/*
	create a cumulative "histogram" from each row and then calc the largest rectangle from the histogram
	0 1 1
	1 1 1
	1 0 0

	histos:
	0 1 1
	1 2 2
	2 0 0

	largest rectangle (looking left to right):
	0 2 1
	3 4 2
	2 0 0
*/

func maximalRectangle(matrix [][]byte) int {
	var total int
	if len(matrix) == 0 {
		return total
	}
	// max histogram
	rowMap := make([]int, len(matrix[0]))
	for i := range matrix {
		// create each row's histogram
		// if a row is zero, hist is zero, otherwise hist is
		// previous hist + 1
		for j := range matrix[i] {
			if matrix[i][j] != 1 && matrix[i][j] != 49 { // difference in int(byte) values
				rowMap[j] = 0
			} else {
				rowMap[j] += 1
			}
		}

		// calc largest rectangle from histogram
		rowMax := largestRectFromHistogram(rowMap)
		if rowMax > total {
			total = rowMax
		}
	}
	return total
}

func largestRectFromHistogram(histo []int) int {
	var max int

	for i := range histo {
		if histo[i] == 0 {
			continue
		}
		if int(histo[i]) > max {
			max = int(histo[i])
		}
		currentMax := histo[i] // currentMax is the "height"
	INNER:
		for j := i + 1; j < len(histo); j++ {
			if histo[j] == 0 {
				break INNER
			}
			if histo[j] < currentMax {
				currentMax = histo[j]
			}
			length := j - i + 1
			area := int(currentMax) * length
			if area > max {
				max = area
			}
		}
	}

	return max
}

/* testing */

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected int
	}{
		{
			input:    [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}},
			expected: 6,
		},
		// {
		// 	input:    [][]byte{},
		// 	expected: 0,
		// },
		// {
		// 	input:    [][]byte{{0}},
		// 	expected: 0,
		// },
		// {
		// 	input:    [][]byte{{1}},
		// 	expected: 1,
		// },
	}
	for i, test := range tests {
		res := maximalRectangle(test.input)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
