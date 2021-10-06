package leetcode

import (
	"reflect"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := 0; i < m; i++ {
		// Insert nums2[j]
		if len(nums2) > j && nums1[i] > nums2[j] {
			// fmt.Println(nums1[:i], nums2[j], nums1[i:m+n-j-1])
			nums1 = append(nums1[:i], append([]int{nums2[j]}, nums1[i:m+n-j-1]...)...)
			j++
			m++
		}
	}

	if j < n {
		for i := range nums1[m:] {
			nums1[m+i] = nums2[j]
			j++
		}
	}
}

/* testing */

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		m, n         int
		expected     []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{-1, 0, 1, 0, 0},
			m:        3,
			nums2:    []int{2, 5},
			n:        2,
			expected: []int{-1, 0, 1, 2, 5},
		},
		{
			nums1:    []int{-1, 0, 1, 0, 0},
			m:        3,
			nums2:    []int{2, 5},
			n:        2,
			expected: []int{-1, 0, 1, 2, 5},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}
	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.nums1, test.expected) {
			t.Errorf("got %v, expected %v", test.nums1, test.expected)
		}
	}
}
