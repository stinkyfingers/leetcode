package edb

import (
	"reflect"
	"testing"
)

func TestRotate2DSlice(t *testing.T) {
	tests := []struct {
		desc     string
		image    [][]int
		expected [][]int
	}{
		{
			desc: "J",
			image: [][]int{
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0},
				{1, 0, 0, 0, 1, 0},
				{0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			desc: "J sideways",
			image: [][]int{
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 1},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
			},
		},
		{
			desc: "2x2",
			image: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			desc:     "nil",
			image:    [][]int{},
			expected: [][]int{},
		},
		{
			desc: "single row",
			image: [][]int{
				{0, 0, 0, 0, 1, 0},
			},
			expected: [][]int{
				{0}, {0}, {0}, {0}, {1}, {0},
			},
		},
		{
			desc: "single column",
			image: [][]int{
				{0},
				{0},
				{0},
				{0},
				{1},
				{0},
			},
			expected: [][]int{
				{0, 1, 0, 0, 0, 0},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			res := rotate2DSlice(tc.image)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, res)
			}
		})
	}
}
