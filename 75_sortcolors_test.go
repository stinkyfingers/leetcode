package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// good explaination here: https://medium.com/nerd-for-tech/75-sort-colors-15768309bf2f

func sortColors(nums []int)  {
	var ones, twos int
	threes := len(nums)-1
	for {
		if twos > threes {
			break
		}
		switch nums[twos] {
		case 0:
			nums[ones], nums[twos] = nums[twos], nums[ones]
			ones+=1
			twos+=1
		case 1:
			twos+=1
		case 2:
			nums[threes], nums[twos] = nums[twos],nums[threes]
			threes-=1
		}
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		desc string
		nums []int 
		expected []int
	}{
		{
			desc:     "1",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			desc:     "2",
			nums:     []int{2, 2, 2, 2, 0},
			expected: []int{0, 2, 2, 2, 2},
		},
		{
			desc:     "3",
			nums:     []int{0},
			expected: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sortColors(tt.nums)
			require.Equal(t, tt.expected, tt.nums)
		})
	}
}
