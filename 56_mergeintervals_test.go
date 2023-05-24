package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104

func mergeX(intervals [][]int) [][]int { // TODO name
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	m := make(map[int][]int)
	var output [][]int
	for _, interval := range intervals {
		a := interval[0]
		b := interval[1]
		if existing, ok := m[a]; ok {
			currentTail := existing[1]
			if b > currentTail {
				existing[1] = b
				for i := currentTail; i <= b; i++ {
					m[i] = existing
				}
			}
		} else if existing, ok := m[b]; ok {
			currentHead := existing[0]
			if a < currentHead {
				existing[0] = a
				for i := a; i <= currentHead; i++ {
					m[i] = existing
				}
			}
		} else {
			newInterval := []int{a, b}
			output = append(output, newInterval)
			for i := a; i <= b; i++ {
				m[i] = newInterval
			}
		}
	}
	return output
}

/* testing */

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		intervals, expected [][]int
	}{
		//{
		//	intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		//	expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		//},
		//{
		//	intervals: [][]int{{1, 4}, {4, 5}},
		//	expected:  [][]int{{1, 5}},
		//},
		//{
		//	intervals: [][]int{{1, 4}, {1, 4}},
		//	expected:  [][]int{{1, 4}},
		//},
		//{
		//	intervals: [][]int{{1, 4}, {0, 4}},
		//	expected:  [][]int{{0, 4}},
		//},
		//{
		//	intervals: [][]int{{1, 4}, {2, 3}},
		//	expected:  [][]int{{1, 4}},
		//},
		//{
		//	intervals: [][]int{{1, 4}},
		//	expected:  [][]int{{1, 4}},
		//},
		//{
		//	intervals: [][]int{{1, 4}, {0, 5}},
		//	expected:  [][]int{{0, 5}},
		//},
		{
			intervals: [][]int{{4, 5}, {1, 4}, {0, 1}},
			expected:  [][]int{{0, 5}},
		},
	}
	for i, test := range tests {
		res := mergeX(test.intervals)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
