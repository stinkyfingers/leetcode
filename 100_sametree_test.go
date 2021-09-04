package leetcode

import (
	"reflect"
	"testing"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil || p == nil && q != nil || p.Val != q.Val {
		return false
	}
	if ok := isSameTree(p.Left, q.Left); !ok {
		return false
	}
	if ok := isSameTree(p.Right, q.Right); !ok {
		return false
	}
	return true
}

/* testing */

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p, q     *TreeNode
		expected bool
	}{
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: true,
		},
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:        &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: false,
		},
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			expected: false,
		},
		{
			p:        &TreeNode{},
			q:        &TreeNode{},
			expected: true,
		},
		{
			p:        &TreeNode{Val: 1},
			q:        &TreeNode{},
			expected: false,
		},
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 2}},
			q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2}},
			expected: false,
		},
	}
	for _, test := range tests {
		res := isSameTree(test.p, test.q)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
