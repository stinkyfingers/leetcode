package leetcode

import (
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	current := nums[0]
	back := 0

	for i := range nums {
		if i == 0 {
			continue
		}

		if nums[i] == current {
			back++
			continue
		}
		current = nums[i]
		nums[i-back] = nums[i]
	}

	return len(nums) - back
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "normal",
			nums:     []int{1, 2, 2, 3, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "empty",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "1",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "normal",
			nums:     []int{2, 2, 2, 2, 2, 2},
			expected: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := removeDuplicates(tt.nums)
			if l != len(tt.expected) {
				t.Error("wrong len ", l)
			}
			for i := 0; i < len(tt.expected); i++ {
				if tt.expected[i] != tt.nums[i] {
					t.Errorf("at %d: %d != %d", i, tt.expected[i], tt.nums[i])
				}
			}
			t.Log(tt.nums, l)
		})
	}
}
