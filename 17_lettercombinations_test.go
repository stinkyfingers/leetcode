package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	var combos []string

	if digits == "" { // empty case defined specifically in challenge
		return combos
	}

	permuteLetters(digits, "", len(digits), func(s string) {
		combos = append(combos, s)
	})
	return combos
}

func phoneLetters(digits string) string {
	var letters string
	for i := range digits {
		switch string(digits[i]) {
		case "2":
			letters += "abc"
		case "3":
			letters += "def"
		case "4":
			letters += "ghi"
		case "5":
			letters += "jkl"
		case "6":
			letters += "mno"
		case "7":
			letters += "pqrs"
		case "8":
			letters += "tuv"
		case "9":
			letters += "wxyz"
		}
	}
	return letters
}

func permuteLetters(remaining, current string, length int, cb func(string)) {
	if len(current) == length {
		cb(current)
		return
	}
	if len(remaining) == 0 {
		return
	}
	for i := range remaining {
		letters := phoneLetters(string(remaining[i]))
		for j := range letters {
			newRemaining := remaining[i+1:]
			permuteLetters(newRemaining, current+string(letters[j]), length, cb)
		}
	}
}

/* testing */

func compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
OUTER:
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				if j < len(b)-1 {
					b = append(b[:j], b[j+1:]...)
				} else {
					b = b[:j]
				}
				continue OUTER
			}
		}
		return false
	}
	return true
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		exp    []string
	}{
		{
			digits: "23",
			exp:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		// {
		// 	digits: "222",
		// 	exp:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		// },
		{
			digits: "",
			exp:    []string{},
		},
		{
			digits: "2",
			exp:    []string{"a", "b", "c"},
		},
	}
	for _, test := range tests {
		res := letterCombinations(test.digits)
		if !compare(res, test.exp) {
			t.Errorf("expected %v, got %v", test.exp, res)
		}
	}
}
