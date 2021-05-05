package leetcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	check := head         // node n from end
	checkMinusOne := head // node n+1 from end
	i := 0
	for current != nil {
		if i >= n {
			checkMinusOne = check
			check = check.Next
		}
		current = current.Next
		i++
	}
	// if length == 1 and we're removing that node, return nil
	if i == 1 && n == 1 {
		return nil
	}

	if check.Next == nil {
		// remove last node
		checkMinusOne.Next = nil
	} else {
		// remove i-nth node
		*check = *check.Next
	}

	return head
}

/* testing */

func TestRemoveNthNode(t *testing.T) {
	tests := []struct {
		head     *ListNode
		n        int
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			n:        2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			head:     &ListNode{Val: 1},
			n:        1,
			expected: nil,
		},
		{
			head:     &ListNode{Val: 1},
			n:        2,
			expected: &ListNode{Val: 1},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			n:        1,
			expected: &ListNode{Val: 1},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			n:        1,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}
	for _, test := range tests {
		res := removeNthFromEnd(test.head, test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %s, expected %s", res.String(), test.expected.String())
		}
	}
}
