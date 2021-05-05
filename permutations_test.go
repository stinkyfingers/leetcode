package leetcode

import (
	"fmt"
	"testing"
)

func permutations(digits []int) [][]int {
	var res [][]int
	permute(digits, []int{}, func(answer []int) {
		res = append(res, answer)
	})
	return res
}

func permute(remaining, current []int, cb func([]int)) {
	newCurrent := current

	if len(remaining) == 0 {
		cb(newCurrent)
		return
	}
	for i, num := range remaining {
		newRemaining := append(remaining[:i:i], remaining[i+1:]...) // also slice capacity
		newCurrent = append(newCurrent, num)
		permute(newRemaining, newCurrent, cb)
		newRemaining = remaining
		newCurrent = current
	}
}

/* testing */

func TestPermute(t *testing.T) {
	tests := []struct {
		digits []int
	}{
		{
			digits: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		res := permutations(test.digits)
		fmt.Println(res)
	}
}
