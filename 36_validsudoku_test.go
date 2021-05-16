package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func isValidSudoku(board [][]byte) bool {
	horiz := make(map[int]map[byte]struct{})
	vert := make(map[int]map[byte]struct{})
	cluster := make(map[string]map[byte]struct{})

	for i, row := range board {
		if horiz[i] == nil {
			horiz[i] = make(map[byte]struct{})
		}
		for j, val := range row {
			if val == 46 {
				continue
			}
			clusterName := fmt.Sprintf("%d-%d", i/3, j/3)
			if vert[j] == nil {
				vert[j] = make(map[byte]struct{})
			}
			if cluster[clusterName] == nil {
				cluster[clusterName] = make(map[byte]struct{})
			}

			if _, ok := horiz[i][val]; ok {
				return false
			}
			if _, ok := vert[j][val]; ok {
				return false
			}
			if _, ok := cluster[clusterName][val]; ok {
				return false
			}
			horiz[i][val] = struct{}{}
			vert[j][val] = struct{}{}
			cluster[clusterName][val] = struct{}{}
		}
	}
	return true
}

/* testing */

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		board    [][]byte
		expected bool
	}{
		{
			board: [][]byte{
				{53, 51, 46, 46, 55, 46, 46, 46, 46},
				{54, 46, 46, 49, 57, 53, 46, 46, 46},
				{46, 57, 56, 46, 46, 46, 46, 54, 46},
				{56, 46, 46, 46, 54, 46, 46, 46, 51},
				{52, 46, 46, 56, 46, 51, 46, 46, 49},
				{55, 46, 46, 46, 50, 46, 46, 46, 54},
				{46, 54, 46, 46, 46, 46, 50, 56, 46},
				{46, 46, 46, 52, 49, 57, 46, 46, 53},
				{46, 46, 46, 46, 56, 46, 46, 55, 57},
			},
			expected: true,
		},
		{
			board: [][]byte{
				{56, 51, 46, 46, 55, 46, 46, 46, 46},
				{54, 46, 46, 49, 57, 53, 46, 46, 46},
				{46, 57, 56, 46, 46, 46, 46, 54, 46},
				{56, 46, 46, 46, 54, 46, 46, 46, 51},
				{52, 46, 46, 56, 46, 51, 46, 46, 49},
				{55, 46, 46, 46, 50, 46, 46, 46, 54},
				{46, 54, 46, 46, 46, 46, 50, 56, 46},
				{46, 46, 46, 52, 49, 57, 46, 46, 53},
				{46, 46, 46, 46, 56, 46, 46, 55, 57},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		res := isValidSudoku(test.board)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
