package leetcode

import (
	"reflect"
	"testing"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	var out [][]int
	if root == nil {
		return out
	}
	fw := true
	current := []*TreeNode{root}
	for len(current) > 0 {
		var level []int
		var next []*TreeNode
		for i := range current {
			if current[i] == nil {
				continue
			}
			if fw {
				level = append(level, current[i].Val)
			} else {
				level = append([]int{current[i].Val}, level...)
			}

			if current[i].Left != nil {
				next = append(next, current[i].Left)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			}

		}
		out = append(out, level)
		current = next
		fw = !fw
	}
	return out
}

/* testing */

func TestZigzagLevelOrder(t *testing.T) {
	var empty [][]int
	tests := []struct {
		root     *TreeNode
		expected [][]int
	}{
		{
			//[3,9,20,null,null,15,7]
			root: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7}}},
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			expected: empty,
		},
		{
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3,
					Right: &TreeNode{Val: 5}},
			},
			expected: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for i, test := range tests {
		res := zigzagLevelOrder(test.root)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("test %d: got %v, expected %v", i, res, test.expected)
		}
	}
}
