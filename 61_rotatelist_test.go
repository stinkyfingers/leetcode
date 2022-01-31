package leetcode

import (
	"reflect"
	"testing"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || length(head) == 1 { // no "list" to rotate; return head
		return head
	}
	if k%length(head) == 0 { // no rotation; return head
		return head
	}
	endPos := length(head) - (k % length(head)) - 1 // new Ending Node index - modulus for multiple complete rotations
	oldHead := head                                 // first node - will become Next of last node
	var newHead *ListNode
	current := head
	for i := 0; i < length(head); i++ {
		next := current.Next
		if i == endPos { // found our new end
			current.Next = nil
			newHead = next // new head
			break
		}
		current = next
	}
	last(newHead).Next = oldHead // assign old head to end of new head
	return newHead
}

// length returns a ListNode's length
func length(n *ListNode) int {
	if n.Next == nil {
		return 1
	}
	return 1 + length(n.Next)
}

// last returns a ListNode's last node
func last(n *ListNode) *ListNode {
	if n.Next == nil {
		return n
	}
	return last(n.Next)
}

/* testing */

func TestRotateRight(t *testing.T) {
	tests := []struct {
		head     *ListNode
		k        int
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:        2,
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			head:     &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			k:        4,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			k: 0,
		},
		{
			head:     &ListNode{Val: 1},
			k:        2,
			expected: &ListNode{Val: 1},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			k:        0,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			k:        1,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			k:        2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
	}
	for _, test := range tests {
		res := rotateRight(test.head, test.k)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
