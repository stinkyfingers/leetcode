package leetcode

import (
	"reflect"
	"testing"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var arrays [][]int
	cur := []int{root.Val}
	calcSum(root, cur, root.Val, targetSum, func(arr []int) {
		arrays = append(arrays, arr)
	})
	return arrays
}

func calcSum(root *TreeNode, cur []int, val int, targetSum int, f func(arr []int)) {
	if val == targetSum && root.Left == nil && root.Right == nil {
		f(cur)
		return
	}

	if root.Left != nil {
		a := make([]int, len(cur)+1)
		copy(a, cur)
		a[len(a)-1] = root.Left.Val
		calcSum(root.Left, a, val+root.Left.Val, targetSum, f)
	}

	if root.Right != nil {
		a := make([]int, len(cur)+1)
		copy(a, cur)
		a[len(a)-1] = root.Right.Val
		calcSum(root.Right, a, val+root.Right.Val, targetSum, f)
	}
}

/* testing */

func TestPathSum2(t *testing.T) {
	tests := []struct {
		p        *TreeNode
		target   int
		expected [][]int
	}{
		{
			p: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			target:   22,
			expected: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			p: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 8,
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			target:   22,
			expected: [][]int{{5, 8, 4, 5}},
		},
		{
			target:   0,
			expected: nil,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			target:   1,
			expected: nil,
		},
		{
			p: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: -3,
				},
			},
			target:   -5,
			expected: [][]int{{-2, -3}},
		},
	}
	for _, test := range tests {
		res := pathSum(test.p, test.target)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
