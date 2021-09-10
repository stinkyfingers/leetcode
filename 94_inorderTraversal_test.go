package leetcode

import (
	"reflect"
	"testing"
)

func inorderTraversal(root *TreeNode) []int {
	var out []int
	inorder(root, func(i int) {
		out = append(out, i)
	})
	return out
}

func inorder(node *TreeNode, f func(i int)) {
	if node == nil {
		return
	}
	if node.Left != nil {
		inorder(node.Left, f)
	}
	f(node.Val)
	if node.Right != nil {
		inorder(node.Right, f)
	}
}

/* testing */

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected []int
	}{
		{
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			expected: []int{1, 3, 2},
		},
		{
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			expected: []int{2, 1},
		},
		{
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: []int{1, 2},
		},
		{
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
		// {
		// 	root:     nil,
		// 	expected: []int{},
		// },
	}
	for i, test := range tests {
		res := inorderTraversal(test.root)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
