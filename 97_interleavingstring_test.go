package leetcode

import (
	"reflect"
	"testing"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for k := range dp {
		dp[k] = make([]bool, len(s2)+1)
	}

	dp[len(s1)][len(s2)] = true // bottom right corner - made it

	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp[i]) - 1; j >= 0; j-- {
			if i < len(s1) && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}
			if j < len(s2) && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}
	return dp[0][0]
}

// type stringPos struct {
// 	s1   int
// 	s2   int
// 	same int
// 	s3   int
// }

// next is brute force way. Call:
/*
	s := &stringPos{
		0, 0, 0, 0,
	}
	return s.next(s1, s2, s3)
*/
// func (s *stringPos) next(s1 string, s2 string, s3 string) bool {
// 	if s.s3 == len(s3) {
// 		if s.s3 == len(s1)+len(s2) {
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	if s.s1 < len(s1) && s.s2 < len(s2) && s3[s.s3] == s1[s.s1] && s3[s.s3] == s2[s.s2] { // fan out
// 		leftS := *s
// 		rightS := *s
// 		leftS.s1++
// 		rightS.s2++
// 		leftS.s3++
// 		rightS.s3++
// 		return leftS.next(s1, s2, s3) || rightS.next(s1, s2, s3)
//
// 	} else if s.s1 < len(s1) && s3[s.s3] == s1[s.s1] {
// 		s.s3++
// 		s.s1++
// 		return s.next(s1, s2, s3)
//
// 	} else if s.s2 < len(s2) && s3[s.s3] == s2[s.s2] {
//
// 		s.s3++
// 		s.s2++
// 		return s.next(s1, s2, s3)
//
// 	}
// 	return false
// }

/* testing */

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1, s2, s3 string
		expected   bool
	}{
		{
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
		},
		{
			s1:       "aabc",
			s2:       "dbbc",
			s3:       "aadbbcbc",
			expected: true,
		},
		{
			s1:       "abc",
			s2:       "abd",
			s3:       "abd",
			expected: false,
		},
		{
			s1:       "abc",
			s2:       "def",
			s3:       "adbecf",
			expected: true,
		},
		{
			s1:       "a",
			s2:       "b",
			s3:       "a",
			expected: false,
		},
		{
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
		},
		{
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbbaccc",
			expected: false,
		},
	}
	for _, test := range tests {
		res := isInterleave(test.s1, test.s2, test.s3)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
