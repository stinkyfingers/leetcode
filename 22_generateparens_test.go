package leetcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var res []string
	parens(n, n, "", func(s string) {
		res = append(res, s)
	})

	return res
}

func parens(open, close int, current string, cb func(string)) {
	if open == 0 && close == 0 {
		cb(current)
		return
	}

	if close == open {
		// must be open
		parens(open-1, close, current+"(", cb)
	} else {
		// and/or close
		if open > 0 {
			parens(open-1, close, current+"(", cb)
		}
		if close > 0 {
			parens(open, close-1, current+")", cb)
		}
	}
}

/* testing */

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:        1,
			expected: []string{"()"},
		},
		{
			n:        2,
			expected: []string{"(())", "()()"},
		},
	}
	for _, test := range tests {
		res := generateParenthesis(test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %s, expected %s", res, test.expected)
		}
	}
}
