package leetcode

import (
	"reflect"
	"testing"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	current := []*TreeNode{root}
	for len(current) > 0 {
		if !mirror(current) {
			return false
		}
		var next []*TreeNode
		for i := range current {
			if current[i] == nil {
				continue
			}
			if current[i].Left != nil {
				next = append(next, current[i].Left)
			} else {
				next = append(next, nil)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			} else {
				next = append(next, nil)
			}
		}
		current = next
	}
	return true
}

func mirror(nodes []*TreeNode) bool {
	// not symmetric
	for i := 0; i < len(nodes)/2; i++ {
		opp := nodes[len(nodes)-i-1]
		if opp == nil && nodes[i] == nil {
			continue
		}
		if opp == nil && nodes[i] != nil || nodes[i] == nil && opp != nil {
			return false
		}

		if nodes[i].Val != opp.Val {
			return false
		}
	}
	return true
}

/* testing */

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected bool
	}{
		{
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3}}},
			expected: true,
		},
		{
			expected: true,
		},
		{
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4}}},
			expected: false,
		},
		{
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3}}},
			expected: false,
		},
		{
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 3},
				}},
			expected: true,
		},
	}
	for i, test := range tests {
		res := isSymmetric(test.root)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("%d: got %v, expected %v", i, res, test.expected)
		}
	}
}
