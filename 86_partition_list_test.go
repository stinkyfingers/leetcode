package leetcode

import (
	"reflect"
	"testing"
)

func partition(head *ListNode, x int) *ListNode {
	var lo *ListNode
	var hi *ListNode
	var newHead, hiHead *ListNode
	current := head
	for current != nil {
		if current.Val < x {
			if lo == nil {
				lo = &ListNode{Val: current.Val}
				newHead = lo
			} else {
				lo.Next = &ListNode{Val: current.Val}
				lo = lo.Next
			}
			lo.Next = nil
		} else {
			if hi == nil {
				hi = &ListNode{Val: current.Val}
				hiHead = hi
			} else {
				hi.Next = &ListNode{Val: current.Val}
				hi = hi.Next
			}
			hi.Next = nil
		}

		current = current.Next
	}
	if lo == nil { // every node higher than n
		return hiHead
	}
	lo.Next = hiHead
	return newHead
}

/* testing */

func TestPartitionList(t *testing.T) {
	tests := []struct {
		head     *ListNode
		n        int
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
			n:        3,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
			n:        1,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
			n:        6,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
		},
		{
			head:     &ListNode{Val: 1},
			n:        2,
			expected: &ListNode{Val: 1},
		},
		{
			head:     &ListNode{},
			n:        2,
			expected: &ListNode{},
		},
		{
			head:     nil,
			n:        2,
			expected: nil,
		},
	}
	for _, test := range tests {
		res := partition(test.head, test.n)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
