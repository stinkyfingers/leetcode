package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/submissions/

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) < len(words[0])*len(words) {
		return nil
	}
	var indices []int
	wordLength := len(words[0])

	wordMap := make(map[string]int)
	for _, word := range words {
		if _, ok := wordMap[word]; !ok {
			wordMap[word] = 1
		} else {
			wordMap[word]++
		}
	}

	for i := 0; i < len(s)-(len(words)*wordLength)+1; i++ {
		start := i
		end := i + len(words)*wordLength
		slice := s[start:end]

		sMap := make(map[string]int)

		for j := 0; j < len(slice)-wordLength+1; j += wordLength {
			thisWord := slice[j : j+wordLength]
			if _, ok := sMap[thisWord]; !ok {
				sMap[thisWord] = 1
			} else {
				sMap[thisWord]++
			}
		}
		if substringHasWords(sMap, wordMap) {
			indices = append(indices, i)
		}
	}

	return indices
}

func substringHasWords(sMap, wordMap map[string]int) bool {
	for word, count := range wordMap {
		if val, ok := sMap[word]; !ok || val < count {
			return false
		}
	}
	return true
}

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s     string
		words []string
		exp   []int
	}{
		{
			s:     "aab",
			words: []string{"a", "b"},
			exp:   []int{1},
		},
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			exp:   []int{0, 9},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			exp:   nil,
		},
		{
			s:     "barfoofoobarthefoobarman",
			words: []string{"foo", "bar", "the"},
			exp:   []int{6, 9, 12},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			exp:   []int{8},
		},
		{
			s:     "lingmindraboofooowingdingbarrwingmonkeypoundcake",
			words: []string{"fooo", "barr", "wing", "ding", "wing"},
			exp:   []int{13},
		},
		{
			s:     "aadbbcc",
			words: []string{"bb", "cc"},
			exp:   []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			res := findSubstring(tt.s, tt.words)
			require.Equal(t, tt.exp, res)
		})
	}
}
