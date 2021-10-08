package leetcode

import (
	"reflect"
	"testing"
)

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if length > 0 {
				break
			}
			continue
		}
		length++
	}
	return length
}

/* testing */

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		{
			word:     "hello World",
			expected: 5,
		},
		{
			word:     "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			word:     "luffy is still joyboy",
			expected: 6,
		},
	}
	for _, test := range tests {
		res := lengthOfLastWord(test.word)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
