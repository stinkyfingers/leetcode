package leetcode

import (
	"reflect"
	"testing"
)

func isValidBST(root *TreeNode) bool {
	return isValidBSTMax(root, nil, nil)
}

func isValidBSTMax(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if max != nil && root.Right.Val >= *max {
			return false
		}

	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		if min != nil && root.Left.Val <= *min {
			return false
		}
	}
	return isValidBSTMax(root.Right, greater(min, &root.Val), max) && isValidBSTMax(root.Left, min, less(max, &root.Val))
}

func greater(a, b *int) *int {
	if a != nil && *a > *b {
		return a
	}
	return b
}
func less(a, b *int) *int {
	if a != nil && *a < *b {
		return a
	}
	return b
}

/* testing */

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected bool
	}{
		{
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			expected: true,
		},
		{
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}}},
			expected: false,
		},
		{
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}},
			expected: true,
		},
		{
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}}},
			expected: false,
		},
		{
			root:     &TreeNode{Val: 2},
			expected: true,
		},
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}},
			expected: true,
		},
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}},
			expected: false,
		},
		{
			root:     &TreeNode{Val: 0, Left: &TreeNode{Val: -1}},
			expected: true,
		},
		{
			root:     &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}},
			expected: true,
		},
		{
			root:     &TreeNode{Val: -2147483648, Right: &TreeNode{Val: 2147483647}},
			expected: true,
		},
	}
	for i, test := range tests {
		res := isValidBST(test.root)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
