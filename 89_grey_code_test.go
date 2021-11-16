package leetcode

import (
	"math"
	"reflect"
	"testing"
)

// grayCode sets {0,1} then iterates over output in reverse adding the next highest bit,
// appending the result:
/*
	00, 01 =>
	11, 10 =>
	110, 111, 101, 100 =>
	1100, 1101, 1111, 1110, 1011, 1010, 1001, 1000 etc...

*/
func grayCode(n int) []int {
	output := []int{0, 1}
	total := int(math.Pow(2.0, float64(n)))
	current := 1
	for {
		if len(output) >= total {
			break
		}
		current = current << 1
		for i := len(output) - 1; i >= 0; i-- {
			output = append(output, output[i]+current)
		}
	}

	return output
}

/* testing */

func TestGrayCode(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{
			input:    2,
			expected: []int{0, 1, 3, 2},
		},
		{
			input:    3,
			expected: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
		{
			input:    1,
			expected: []int{0, 1},
		},
	}
	for i, test := range tests {
		res := grayCode(test.input)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
