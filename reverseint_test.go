package leetcode

import (
	"math"
	"testing"
)

func reverse(x int) int {
	res := 0

	var digits []int
	for place := 10; place < 100000000000; place *= 10 {
		digit := x % place
		digits = append([]int{digit / (place / 10)}, digits...)
		x = x - digit
		if x == 0 {
			break
		}
	}
	if x != 0 {
		return 0
	}
	for i, digit := range digits {
		res += digit * int(math.Pow(10.0, float64(i)))
		if res >= math.MaxInt32 || res <= math.MinInt32 {
			return 0
		}
	}
	return res
}

func TestReverse(t *testing.T) {
	tests := []struct {
		x   int
		exp int
	}{
		{
			x:   123,
			exp: 321,
		},
		{
			x:   -123,
			exp: -321,
		},
		{
			x:   0,
			exp: 0,
		},
		{
			x:   120,
			exp: 21,
		},
		{
			x:   201,
			exp: 102,
		},
		{
			x:   7463847412,
			exp: 0,
		},
		{
			x:   -8463847412,
			exp: 0,
		},
		{
			x:   -2147483412,
			exp: -2143847412,
		},
	}
	for _, test := range tests {
		res := reverse(test.x)
		if res != test.exp {
			t.Errorf("expected %d, got %d", test.exp, res)
		}
	}
}
