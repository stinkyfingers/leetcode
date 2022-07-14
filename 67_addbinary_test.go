package leetcode

import (
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/require"
)

func pad (s string, i int) string {
	l := len(s)
	for i > l {
		s = fmt.Sprintf("0%s", s)
		l = len(s)
	} 
	return s
}

func addBinary(a string, b string) string {
	var total string 
	a = pad(a, len(b))
	b = pad(b, len(a))
	
	var carry bool 
	for i := len(a) - 1; i >= 0; i-- {
		var valA, valB uint8
		if i >= len(a) {
			valA = 0
		} else {
			valA = a[i] - 48
		}
		if i >= len(b) {
			valB = 0
		} else {
			valB = b[i] - 48
		}

		digit := "0"
		switch valA + valB {
		case 0:
			if carry {
				digit = "1"
			}
			carry = false
		case 1:
			if carry {
				carry = true
			} else {
				digit = "1"
				carry = false
			}
		case 2:
			if carry {
				digit = "1"
			}
			carry = true
		}
		total = fmt.Sprintf("%s%s", digit, total)
	}
	if carry {
		total = fmt.Sprintf("1%s", total)
	}
	return total
}

func TestAddBinary(t *testing.T) {
	tests := []struct {
		desc, a, b, expected string
	}{
		{
			a: "11",
			b: "1",
			expected: "100",
		},
		{
			a: "1010",
			b: "1011",
			expected: "10101",
		},
		{
			a: "",
			b: "1",
			expected: "1",
		},
		{
			a: "111",
			b: "111",
			expected: "1110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			res := addBinary(tt.a, tt.b)
			require.Equal(t, tt.expected, res)
		})
	}
}
