package leetcode

import (
	"reflect"
	"testing"
)

func maxDepth(root *TreeNode) int {
	var max int
	runDepth(root, 0, func(i int) {
		if i > max {
			max = i
		}
	})
	return max
}

func runDepth(root *TreeNode, cur int, f func(i int)) {
	if root == nil {
		f(cur)
		return
	}
	runDepth(root.Left, cur+1, f)
	runDepth(root.Right, cur+1, f)
}

/* testing */

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		p        *TreeNode
		expected int
	}{
		{
			p:        &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			expected: 3,
		},
		{
			p:        &TreeNode{Val: 3},
			expected: 1,
		},
		{
			p:        &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 15, Left: &TreeNode{Val: 7}}}},
			expected: 4,
		},
	}
	for _, test := range tests {
		res := maxDepth(test.p)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
