package leetcode

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	Recursive tree
*/

func combinationSum2(candidates []int, target int) [][]int {
	var output [][]int
	labelMap := make(map[string]struct{})
	sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
	findSums([]int{}, candidates, target, func(arr []int) {
		if _, ok := labelMap[fmt.Sprintf("%v", arr)]; ok {
			return
		}
		labelMap[fmt.Sprintf("%v", arr)] = struct{}{}
		output = append(output, arr)
	})

	return output
}

func findSums(current []int, candidates []int, target int, f func(arr []int)) {
	if sumInts(current) == target {
		f(current)
		return
	}
	if sumInts(current) > target {
		return
	}

	prev := -1

	for i, candidate := range candidates {
		if candidate == prev { // skip if previous candidate was same value (must be sorted)
			continue
		}
		newCurrent := []int{}
		newCurrent = append(newCurrent, current...)
		newCurrent = append(newCurrent, candidate)
		findSums(newCurrent, candidates[i+1:], target, f)
		prev = candidate
	}
}

func sumInts(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

/* testing */

func TestCombinationSum2(t *testing.T) {
	var none [][]int
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			// 10 7 6 5 2 1 1
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected:   [][]int{{7, 1}, {6, 2}, {6, 1, 1}, {5, 2, 1}},
		},
		{
			// 10 7 6 5 2 1 1
			candidates: []int{5, 2, 2, 1, 1},
			target:     7,
			expected:   [][]int{{5, 2}, {5, 1, 1}},
		},
		{
			candidates: []int{2, 5, 2, 1},
			target:     5,
			expected:   [][]int{{5}, {2, 2, 1}},
		},
		{
			candidates: []int{1, 2},
			target:     3,
			expected:   [][]int{{2, 1}},
		},
		{
			candidates: []int{1},
			target:     3,
			expected:   none,
		},
		{
			// 5 2 2 2 1
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected:   [][]int{{5}, {2, 2, 1}},
		},
		{
			candidates: []int{4, 2, 5, 2, 5, 3, 1, 5, 2, 2},
			target:     9,
			expected:   [][]int{{5, 4}, {5, 3, 1}, {5, 2, 2}, {4, 3, 2}, {4, 2, 2, 1}, {3, 2, 2, 2}, {2, 2, 2, 2, 1}},
		},
		{
			candidates: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target:     27,
			expected:   none,
		},
	}
	for _, test := range tests {
		res := combinationSum2(test.candidates, test.target)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
