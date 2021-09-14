package leetcode

import (
	"reflect"
	"testing"
)

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	current := []*TreeNode{root}
	for len(current) > 0 {
		var next []*TreeNode
		var level []int
		for i := range current {
			if current[i] == nil {
				continue
			}
			level = append(level, current[i].Val)
			if current[i].Left != nil {
				next = append(next, current[i].Left)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			}
		}
		current = next
		res = append(res, level)
	}
	return res
}

/* testing */

func TestLevelOrder(t *testing.T) {
	var empty [][]int
	tests := []struct {
		root     *TreeNode
		expected [][]int
	}{
		{
			root: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7}}},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			expected: empty,
		},
		{
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
	}
	for i, test := range tests {
		res := levelOrder(test.root)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
