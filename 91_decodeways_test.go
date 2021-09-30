package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

// func numDecodings(s string) int {
// 	m := make(map[string]int)
// 	return recNumDecodings(s, m)
// }
// func recNumDecodings(s string, m map[string]int) int {
// 	var count int
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	if len(s) == 1 {
// 		if valid(s) {
// 			return 1
// 		}
// 		return 0
// 	}
// 	if len(s) == 2 {
// 		total := 0
// 		if valid(string(s[0])) && valid(string(s[1])) {
// 			total += 1
// 		}
// 		if valid2(s) {
// 			total += 1
// 		}
// 		return total
// 	}
//
// 	if valid(string(s[0])) {
// 		count += numDecodings(s[1:])
// 	}
//
// 	if val, ok := m[s[0:2]]; ok {
// 		return val
// 	}
// 	if valid2(s[0:2]) {
// 		count += numDecodings(s[2:])
// 	}
//
// 	return count
// }
//
// func valid(s string) bool {
// 	a, err := strconv.Atoi(s)
// 	if err != nil || a < 1 {
// 		return false
// 	}
// 	return true
// }
//
// func valid2(s string) bool {
// 	if string(s[0]) == "0" {
// 		return false
// 	}
// 	a, err := strconv.Atoi(s)
// 	if err != nil || a < 1 || a > 26 {
// 		return false
// 	}
// 	return true
// }

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1 //? init one - uses one-based array, zero position is inited
	dp[1] = 1 // first char - one or zero
	if s[0] == '0' {
		dp[1] = 0
	}

	for i := 2; i <= len(s); i++ {
		oneDigit, _ := strconv.Atoi(s[i-1 : i])
		twoDigit, _ := strconv.Atoi(s[i-2 : i])
		if oneDigit >= 1 {
			dp[i] += dp[i-1]
		}
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

/* testing */

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "12",
			expected: 2,
		},
		{
			s:        "226",
			expected: 3,
		},
		{
			s:        "0",
			expected: 0,
		},
		{
			s:        "20",
			expected: 1,
		},
		{
			s:        "06",
			expected: 0,
		},
		{
			s:        "010",
			expected: 0,
		},
	}
	for _, test := range tests {
		res := numDecodings(test.s)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
