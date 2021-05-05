package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/
// UNSOLVED...so far

func reverseKGroup(head *ListNode, k int) *ListNode {
	current := head
	var prev *ListNode
	var end, prevEnd *ListNode
	// var newHead *ListNode
	next := head
	i := 0

	for current != nil {
		fmt.Println(i, current, prev)

		if i == 0 {
			end = current
		}

		if i < k {
			next = current.Next
			current.Next = prev
			prev = current
			current = next
			i++
			continue
		}
		if i == k {
			if prevEnd != nil {
				prevEnd.Next = current
			}
			prevEnd = end

			current = current.Next
			i = 0
		}
	}
	return prev
}

// reverse a singly linked list
func reverseList(head *ListNode, k int) *ListNode {
	current := head
	var prev *ListNode
	next := head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

/* testing */

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		head     *ListNode
		n        int
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			n:        2,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}},
		},
	}
	for _, test := range tests {
		res := reverseKGroup(test.head, test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %s, expected %s", res.String(), test.expected.String())
		}
	}
}
