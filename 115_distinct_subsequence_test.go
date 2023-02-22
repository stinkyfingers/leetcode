package leetcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
dynamic prog:
- - a b b b
- 1 1 1 1 1
a 0 1 1 1 1
b 0 0 1 2 3*
*/

func numDistinct(s, t string) int {
	var dp [][]int // [t][s]
	for i := 0; i < len(t)+1; i++ {
		dp = append(dp, make([]int, len(s)+1))
	}
	for j := range s {
		dp[0][j] = 1
	}
	for i := range t {
		for j := range s {
			if string(t[i]) == string(s[j]) {
				// match: cell = sum of cell to upper left + cell to left
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			} else {
				// no match: cell = cell to left
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}
	return dp[len(t)][len(s)]
}

// NOTE: recursion works, but is "too slow."
type distinctNode struct {
	sPos int
	tPos int
}

func numDistinctRecursive(s string, t string) int {
	root := &distinctNode{
		sPos: 0,
		tPos: 0,
	}

	var total int
	callback := func() {
		total++
	}
	root.branchRecursive(cleanT(s, t), t, callback)
	return total
}

func (d *distinctNode) branchRecursive(s, t string, callback func()) {
	if d.tPos == len(t) {
		callback()
		return
	}
	if d.sPos == len(s) {
		return
	}
	for i := d.sPos; i < len(s); i++ {
		if s[i] != t[d.tPos] {
			continue
		}
		next := &distinctNode{
			sPos: i + 1,
			tPos: d.tPos + 1,
		}
		next.branchRecursive(s, t, callback)
	}
}

func cleanT(s, t string) string {
	m := make(map[int32]struct{})
	for _, val := range t {
		m[val] = struct{}{}
	}
	var out strings.Builder
	for _, val := range s {
		if _, ok := m[val]; ok {
			out.WriteRune(val)
		}
	}
	return out.String()
}

/* testing */

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			s:        "abbb",
			t:        "ab",
			expected: 3,
		},
		{
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
		{
			s:        "aabdbaabeeadcbbdedacbbeecbabebaeeecaeabaedadcbdbcdaabebdadbbaeabdadeaabbabbecebbebcaddaacccebeaeedababedeacdeaaaeeaecbe",
			t:        "bddabdcae",
			expected: 10582116,
		},
		{
			s:        "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc",
			t:        "bcddceeeebecbc",
			expected: 700531452,
		},
	}
	for _, test := range tests {
		res := numDistinct(test.s, test.t)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
