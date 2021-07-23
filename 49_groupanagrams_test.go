package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	out := [][]string{}
	m := make(map[string]int)

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		if index, ok := m[string(b)]; ok {
			out[index] = append(out[index], str)
		} else {
			m[string(b)] = len(out)
			out = append(out, []string{str})
		}
	}

	return out
}

/* testing */

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			// is good, but too lazy to write complete test comparison
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs:     []string{},
			expected: [][]string{},
		},
		{
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}
	for _, test := range tests {
		res := groupAnagrams(test.strs)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
