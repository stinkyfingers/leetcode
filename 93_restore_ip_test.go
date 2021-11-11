package leetcode

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// uses recursion; uses ip struct (shallow copy) to store current ip value and which segment
// we're on. When we've completed segment 4 and there is no s string left, we add the value to the output
/*
	255
	| | \
	2 25 255
	..and so on
*/

type ip struct {
	val     string
	segment int
}

func restoreIpAddresses(s string) []string {
	var out []string
	f := func(val string) {
		out = append(out, val)
	}
	i := ip{
		segment: 1,
	}
	evalIps(s, i, f)

	return out
}

func evalIps(s string, i ip, f func(val string)) {
	if i.segment >= 5 {
		if len(s) == 0 {
			f(i.val)
		}
		return
	}

	maxIndex := 3
	if len(s) < 3 {
		maxIndex = len(s)
	}

	for j := 0; j < maxIndex; j++ {
		val := s[0 : j+1]
		if !validIp(val) {
			continue
		}

		newIp := i
		if newIp.segment > 1 {
			newIp.val = fmt.Sprintf("%s.%s", newIp.val, val)
		} else {
			newIp.val = val
		}
		newIp.segment++

		evalIps(s[j+1:], newIp, f)
	}

}

func validIp(s string) bool {
	if s == "" || (s[0] == '0' && len(s) > 1) {
		return false
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if val > 255 {
		return false
	}
	return true
}

/* testing */

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		s        string
		expected []string
	}{
		{
			s: "255",
		},
		{
			s:        "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s:        "0000",
			expected: []string{"0.0.0.0"},
		},
		{
			s:        "1111",
			expected: []string{"1.1.1.1"},
		},
		{
			s:        "010010",
			expected: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			s:        "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for i, test := range tests {
		res := restoreIpAddresses(test.s)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
