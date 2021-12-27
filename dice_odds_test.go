package leetcode

import (
	"math"
	"reflect"
	"testing"
)

/*
  calc the odds of rolling a number (1-6), regardless of the number of rolls. Useful in board games in determining what
  the odds of landing on a space are, allowing for multiple turns to arrive at the space. E.g.

  1: 1 in 6
  2: 7 in 36 (1 in 6 of rolling a "2", plus 1 in 36 of rolling a "1" then another "1")
  3: 49 in 216 (1 in 6 of rolling a "3", plus 1 in 36 of rolling a "1" then "2", plus 1 in 36 of rolling a "2" then "1", plus a 1 in 216 of rolling a "1" then "1" then "1")

  1\(6^n) n: number of rolls

  NOTE: only takes one die into account
*/

func diceOdds(num int) float64 {
	var total float64

	calcOdds(num, 1, func(inc float64) {
		total += inc
	})

	return total
}

func calcOdds(num int, rounds int, f func(total float64)) {
	for i := 1; i <= 6; i++ {
		if i > num {
			break
		}

		if i == num {
			denom := float64(math.Pow(6.0, float64(rounds)))
			f(1 / denom)
			break
		}
		if i < num {
			calcOdds(i, rounds+1, f)
		}
	}
}

/* testing */

func TestDiceOdds(t *testing.T) {
	tests := []struct {
		num      int
		expected float64
	}{
		{
			num:      1,
			expected: 1 / float64(6),
		},
		{
			num:      2,
			expected: 1/float64(6) + 1/float64(36), // .19444...
		},
		{
			num:      3,
			expected: 1/float64(6) + 2/float64(36) + 1/float64(216), // .22685185185...
		},
		{
			num:      4,
			expected: 1/float64(6) + 3/float64(36) + 3/float64(216) + 1/float64(216*6),
		},
	}
	for _, test := range tests {
		res := diceOdds(test.num)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
