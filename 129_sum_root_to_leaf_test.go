package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
		 4
		9  0
	   5 1

495, 491, 40
*/
var answer int

func sumNumbers(root *TreeNode) int {
	answer = 0
	calc(root, 0)
	return answer
}

func calc(node *TreeNode, val int) {
	if node == nil {
		return
	}
	val = val*10 + node.Val
	if node.Left == nil && node.Right == nil {
		answer += val
		return
	}
	calc(node.Left, val)
	calc(node.Right, val)
}

/* testing */

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 25,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: 123,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 1026,
		},
		/*
			    8
			   3 5
			    9
			   9 5
			8399,8395,85
		*/
		{
			root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			expected: 16879,
		},
	}
	for _, test := range tests {
		res := sumNumbers(test.root)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
