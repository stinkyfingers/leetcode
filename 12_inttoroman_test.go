package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

// https://leetcode.com/problems/integer-to-roman/

var numsMap = map[int]string{
	0:    "",
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	builder := strings.Builder{}

	for i := 1000; i > 0; i /= 10 {
		digits := num / i
		num = num % i

		// handle 9
		if digits == 9 {
			builder.WriteString(numsMap[i])
			builder.WriteString(numsMap[i*10])
			continue
		}

		quantity5 := digits / 5 * i * 5 // # 5 group
		remainder5 := digits % 5        // left over ones groups
		builder.WriteString(numsMap[quantity5])
		if remainder5 > 3 {
			builder.WriteString(numsMap[i])
			builder.WriteString(numsMap[i*5])
		} else {
			for j := 0; j < remainder5; j++ {
				builder.WriteString(numsMap[i])
			}
		}
	}

	return builder.String()
}

/* testing */

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{
			num:      3,
			expected: "III",
		},
		{
			num:      4,
			expected: "IV",
		},
		{
			num:      9,
			expected: "IX",
		},
		{
			num:      58,
			expected: "LVIII",
		},
		{
			num:      1994,
			expected: "MCMXCIV",
		},
		{
			num:      60,
			expected: "LX",
		},
	}
	for _, test := range tests {
		res := intToRoman(test.num)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %s, expected %s", res, test.expected)
		}
	}
}
