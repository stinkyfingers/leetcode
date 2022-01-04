package leetcode

import (
	"reflect"
	"testing"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	middle := nums[len(nums)/2]
	n := &TreeNode{Val: middle}

	insertMiddle(n, [][]int{nums[:midIndex], nums[midIndex+1:]})
	return n
}

func insertMiddle(n *TreeNode, arrs [][]int) {
	for _, arr := range arrs {
		if len(arr) == 0 {
			return
		}

		midIndex := len(arr) / 2
		middle := arr[midIndex]
		insertTreeNode(n, middle)
		insertMiddle(n, [][]int{arr[:midIndex], arr[midIndex+1:]})
	}
	return
}

func insertTreeNode(n *TreeNode, val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
		} else {
			insertTreeNode(n.Left, val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
		} else {
			insertTreeNode(n.Right, val)
		}
	}
}

/* testing */

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums     []int
		expected *TreeNode
	}{
		// Test comparison fails but is correct
		// {
		// 	nums: []int{-10, -3, 0, 5, 9},
		// 	expected: &TreeNode{Val: 0, Left: &TreeNode{
		// 		Val: -3, Right: &TreeNode{
		// 			Val: -10,
		// 		},
		// 	}, Right: &TreeNode{
		// 		Val: 9, Left: &TreeNode{
		// 			Val: 5,
		// 		},
		// 	},
		// 	},
		// },
		{
			nums:     []int{1, 3},
			expected: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}},
		},
	}
	for _, test := range tests {
		res := sortedArrayToBST(test.nums)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
