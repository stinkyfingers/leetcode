package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func setZeroes(matrix [][]int)  {
	change := make(map[string]struct{})
	for i := range matrix {
		for j := range matrix[i] {
			if _, ok := change[key(i,j)]; ok {
				matrix[i][j] = 0 // is in change map - set value to zero
			} else {
				if matrix[i][j] != 0 {
					continue // not zero, continue
				}
				// set change row
				for k := 0; k < len(matrix[i]); k++ {
					if matrix[i][k] == 0 {
						continue // do not set if already zero
					}
					change[key(i, k)] = struct{}{}
					matrix[i][k] = 0
				}
				// set change col
				for k := 0; k < len(matrix); k++ {
					if matrix[k][j] == 0 {
						continue // do not set if already zero
					}
					change[key(k, j)] = struct{}{}
					matrix[k][j] = 0
				}
				
			}
		}
	}
}

func key(a, b int) string {
	return fmt.Sprintf("%d:%d", a, b)
}

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		desc string
		matrix, expected [][]int
	}{
		{
			desc: "1",
			matrix: [][]int{{1,1,1},{1,0,1},{1,1,1}},
			expected: [][]int{{1,0,1},{0,0,0},{1,0,1}},
		},
		{
			desc: "2",
			matrix: [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
			expected: [][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			setZeroes(tt.matrix)
			require.Equal(t, tt.expected,tt.matrix)
		})
	}
}
