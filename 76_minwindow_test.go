package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func minWindow(s string, t string) string {
	bestWindow := ""
	start := 0
	end := 0
	tMap := make(map[rune]int)
	for _, r := range t {
		if _, ok := tMap[r]; !ok {
			tMap[r] = 1
		} else {
			tMap[r]++
		}
	}
	sMap := make(map[rune]int)

	for end <= len(s) {
		if containsSubstring(sMap, tMap) {
			if len(s[start:end]) < len(bestWindow) || bestWindow == "" {
				bestWindow = s[start:end]
			}
			sMap[rune(s[start])]--
			start++

		} else {
			if end >= len(s) {
				break
			}
			if _, ok := sMap[rune(s[end])]; !ok {
				sMap[rune(s[end])] = 1
			} else {
				sMap[rune(s[end])]++
			}
			end++
		}
	}

	return bestWindow
}

// returns true if s contains t
func containsSubstring(s, t map[rune]int) bool {
	for k, v := range t {
		if quant, ok := s[k]; !ok || quant < v {
			return false
		}
	}
	return true
}

// returns true if s contains t
func containsSubstringX(s string, t string) bool {
	temp := make(map[rune]int)
	for _, r := range s {
		if _, ok := temp[r]; !ok {
			temp[r] = 1
		} else {
			temp[r]++
		}
	}

	for _, r := range t {
		if quantity, ok := temp[r]; !ok || quantity < 1 {
			return false
		}
		temp[r]--
	}
	return true
}

// map[0:1 3:0 9:0 12:1]
func TestMinWindow(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		exp string
	}{
		{
			s:   "ADOBECODEBANC",
			t:   "ABC",
			exp: "BANC",
		},
		{
			s:   "ABBBABC",
			t:   "ABC",
			exp: "ABC",
		},
		{
			s:   "a",
			t:   "a",
			exp: "a",
		},
		{
			s:   "a",
			t:   "aa",
			exp: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			res := minWindow(tt.s, tt.t)
			require.Equal(t, tt.exp, res)
		})
	}
}

func TestContainsSubstring(t *testing.T) {
	tests := []struct {
		desc string
		s    map[rune]int
		t    map[rune]int
		exp  bool
	}{
		{
			desc: "1",
			s:    map[rune]int{'A': 1, 'B': 1, 'C': 1},
			t:    map[rune]int{'A': 1, 'B': 1, 'C': 1},
			exp:  true,
		},
		{
			desc: "2",
			s:    map[rune]int{'A': 2, 'B': 1},
			t:    map[rune]int{'A': 1, 'B': 1, 'C': 1},
			exp:  false,
		},
		{
			desc: "3",
			s:    map[rune]int{'A': 2, 'B': 1, 'C': 1},
			t:    map[rune]int{'A': 2, 'B': 1},
			exp:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			res := containsSubstring(tt.s, tt.t)
			require.Equal(t, tt.exp, res)
		})
	}
}

func TestContainsSubstringX(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		exp bool
	}{
		{
			s:   "ABC",
			t:   "ABC",
			exp: true,
		},
		{
			s:   "ABC",
			t:   "ABA",
			exp: false,
		},
		{
			s:   "ABA",
			t:   "BAA",
			exp: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			res := containsSubstringX(tt.s, tt.t)
			require.Equal(t, tt.exp, res)
		})
	}
}
