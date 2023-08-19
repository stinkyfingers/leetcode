package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Explanation: Notice that an 'O' should not be flipped if:
- It is on the border, or
- It is adjacent to an 'O' that should not be flipped.
The bottom 'O' is on the border, so it is not flipped.
The other three 'O' form a surrounded region, so they are flipped.
*/

func solve(board [][]byte) {
	if len(board) < 2 || len(board[0]) < 2 {
		return
	}
	ohs := make(map[[2]int]bool)
	rim := [][2]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 88 {
				continue
			}
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1 {
				rim = append(rim, [2]int{i, j})
			}
		}
	}
	snake(rim, board, ohs)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !ohs[[2]int{i, j}] {
				board[i][j] = 88
			}
		}
	}
}

func snake(coords [][2]int, board [][]byte, ohs map[[2]int]bool) {
	for len(coords) > 0 {
		coord := coords[0]
		coords = coords[1:]
		ohs[coord] = true
		i, j := coord[0], coord[1]
		if i-1 > 0 && !ohs[[2]int{i - 1, j}] && board[i-1][j] == 79 {
			ohs[[2]int{i, j}] = true
			coords = append(coords, [2]int{i - 1, j})
		}
		if j-1 > 0 && !ohs[[2]int{i, j - 1}] && board[i][j-1] == 79 {
			ohs[[2]int{i, j}] = true
			coords = append(coords, [2]int{i, j - 1})
		}
		if i+1 < len(board) && !ohs[[2]int{i + 1, j}] && board[i+1][j] == 79 {
			ohs[[2]int{i, j}] = true
			coords = append(coords, [2]int{i + 1, j})
		}
		if j+1 < len(board[0])-1 && !ohs[[2]int{i, j + 1}] && board[i][j+1] == 79 {
			ohs[[2]int{i, j}] = true
			coords = append(coords, [2]int{i, j + 1})
		}
	}
}

/* testing */

func Test_surroundedRegionss(t *testing.T) {
	tests := []struct {
		board    [][]byte
		expected [][]byte
	}{
		{
			board:    [][]byte{{88, 88, 88, 88}, {88, 79, 79, 88}, {88, 88, 79, 88}, {88, 79, 88, 88}},
			expected: [][]byte{{88, 88, 88, 88}, {88, 88, 88, 88}, {88, 88, 88, 88}, {88, 79, 88, 88}},
		},
		{
			board:    [][]byte{{88}},
			expected: [][]byte{{88}},
		},
		{
			board:    [][]byte{{88, 88, 88, 88}, {88, 79, 79, 88}, {88, 79, 88, 88}, {88, 79, 88, 88}},
			expected: [][]byte{{88, 88, 88, 88}, {88, 79, 79, 88}, {88, 79, 88, 88}, {88, 79, 88, 88}},
		},
	}
	for _, test := range tests {
		solve(test.board)
		fmt.Println(test.board, "...", test.expected)
		if !reflect.DeepEqual(test.board, test.expected) {
			t.Errorf("got %v, expected %v", test.board, test.expected)
		}
	}
}
