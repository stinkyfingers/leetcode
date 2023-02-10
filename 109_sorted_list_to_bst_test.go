package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var arr []int
	current := head
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	root := &TreeNode{}
	buildBranch(root, arr)

	return root
}

func buildBranch(head *TreeNode, arr []int) {
	if len(arr) == 0 {
		return
	}
	mid := len(arr) / 2
	lower := arr[:mid]
	higher := arr[mid+1:]
	head.Val = arr[mid]
	if len(lower) > 0 {
		head.Left = &TreeNode{}
		buildBranch(head.Left, lower)
	}
	if len(higher) > 0 {
		head.Right = &TreeNode{}
		buildBranch(head.Right, higher)
	}
}

/* testing */

func TestSortedListToBST(t *testing.T) {
	tests := []struct {
		p        *ListNode
		expected *TreeNode
	}{
		{
			p: &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}},
			expected: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		//{
		//	//-11 -10 -5 -3 0 3 5 10 11
		//	p: &ListNode{
		//		Val: -11, Next: &ListNode{
		//			Val: -10, Next: &ListNode{
		//				Val: -5, Next: &ListNode{
		//					Val: -3, Next: &ListNode{
		//						Val: 0, Next: &ListNode{
		//							Val: 3, Next: &ListNode{
		//								Val: 5, Next: &ListNode{
		//									Val: 10, Next: &ListNode{
		//										Val: 11,
		//									}}}}}}}}},
		//	expected: &TreeNode{
		//		Val: 0,
		//		Left: &TreeNode{
		//			Val: -5,
		//			Left: &TreeNode{
		//				Val: -10,
		//				Left: &TreeNode{
		//					Val: -11,
		//				},
		//			},
		//			Right: &TreeNode{
		//				Val: -3,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val: 10,
		//			Left: &TreeNode{
		//				Val: 5,
		//				Left: &TreeNode{
		//					Val: 3,
		//				},
		//			},
		//			Right: &TreeNode{
		//				Val: 11,
		//			},
		//		},
		//	},
		//},
	}
	for _, test := range tests {
		res := sortedListToBST(test.p)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
