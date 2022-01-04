package leetcode

import (
	"reflect"
	"testing"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	current := []*TreeNode{root}
	depth := 0
OUTER:
	for len(current) != 0 {
		depth++

		newCurrent := []*TreeNode{}
		for _, n := range current {
			if n.Left == nil && n.Right == nil {
				break OUTER
			}
			if n.Left != nil {
				newCurrent = append(newCurrent, n.Left)
			}

			if n.Right != nil {
				newCurrent = append(newCurrent, n.Right)
			}
		}
		current = newCurrent
	}

	return depth
}

/* testing */

func TestMinDepth(t *testing.T) {
	tests := []struct {
		n        *TreeNode
		expected int
	}{
		{
			n:        &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			expected: 2,
		},
		{
			n:        &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}}}},
			expected: 4,
		},
		{
			n:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}},
			expected: 2,
		},
		{
			expected: 0,
		},
	}
	for _, test := range tests {
		res := minDepth(test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
