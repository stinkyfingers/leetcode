package leetcode

import (
	"reflect"
	"testing"
)

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) == 0 {
		return false
	}
	arr := []byte(word)
	starts := getStartingPoints(board, arr[0])
	for _, start := range starts {
		if ok := trace(board, [][2]int{start}, arr[1:]); ok {
			return true
		}
	}

	return false
}

func getStartingPoints(board [][]byte, letter byte) [][2]int {
	var out [][2]int
	for i := range board {
		for j := range board[i] {
			if board[i][j] == letter {
				out = append(out, [2]int{i, j})
			}
		}
	}
	return out
}

func trace(board [][]byte, path [][2]int, arr []byte) bool {
	if len(arr) == 0 {
		return true
	}
	currentCoord := path[len(path)-1]
	x := currentCoord[0]
	y := currentCoord[1]

	// left
	if y > 0 && !visited(path, [2]int{x, y - 1}) && board[x][y-1] == arr[0] {
		newPath := append(path, [2]int{x, y - 1})
		if ok := trace(board, newPath, arr[1:]); ok {
			return true
		}
	}
	// up
	if x > 0 && !visited(path, [2]int{x - 1, y}) && board[x-1][y] == arr[0] {
		newPath := append(path, [2]int{x - 1, y})
		if ok := trace(board, newPath, arr[1:]); ok {
			return true
		}
	}
	// right
	if y < len(board[0])-1 && !visited(path, [2]int{x, y + 1}) && board[x][y+1] == arr[0] {
		newPath := append(path, [2]int{x, y + 1})
		if ok := trace(board, newPath, arr[1:]); ok {
			return true
		}
	}
	// down
	if x < len(board)-1 &&
		!visited(path, [2]int{x + 1, y}) &&
		board[x+1][y] == arr[0] {
		newPath := append(path, [2]int{x + 1, y})
		if ok := trace(board, newPath, arr[1:]); ok {
			return true
		}
	}
	return false
}

func visited(path [][2]int, coord [2]int) bool {
	for _, p := range path {
		if coord == p {
			return true
		}
	}
	return false
}

/* testing */

func TestWordSearch(t *testing.T) {
	// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board:    [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 69}},
			word:     "ABCCED",
			expected: true,
		},
		{
			board:    [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 69}},
			word:     "SEE", // 83 69 69
			expected: true,
		},
		{
			board:    [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 69}},
			word:     "ABCB", // 65 66 67 66
			expected: false,
		},
		{
			board:    nil,
			word:     "A",
			expected: false,
		},
		{
			board:    nil,
			word:     "",
			expected: true,
		},
	}
	for _, test := range tests {
		res := exist(test.board, test.word)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
