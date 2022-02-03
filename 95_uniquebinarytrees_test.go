package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func generateTrees(n int) []*TreeNode {
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	var res []*TreeNode
	if start == end {
		return []*TreeNode{{Val: start}}
	}
	if start > end {
		res = append(res, nil)
	}

	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, j := range left {
			for _, k := range right {
				root := &TreeNode{Val: i, Left: j, Right: k}
				res = append(res, root)
			}
		}
	}
	return res
}

/* testing */

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		nums     int
		expected []*TreeNode
	}{
		{
			nums: 3,
			expected: []*TreeNode{
				{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
				{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
				{Val: 2, Right: &TreeNode{Val: 3}, Left: &TreeNode{Val: 1}},
				{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
				{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
			},
		},
		// {
		// 	nums: 2,
		// 	expected: []*TreeNode{
		// 		{Val: 1, Right: &TreeNode{Val: 2}},
		// 		{Val: 2, Left: &TreeNode{Val: 1}},
		// 	},
		// },
	}
	for _, test := range tests {
		res := generateTrees(test.nums)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
			for _, r := range res {
				fmt.Println(r.String())
			}
		}
	}
}
