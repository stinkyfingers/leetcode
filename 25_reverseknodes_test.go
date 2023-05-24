package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/
// UNSOLVED...so far
func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead, newCurrent *ListNode
	m := make(map[int]*ListNode)
	current := head
	for {
		for i := k - 1; i >= 0; i-- {
			m[i] = current
			// TODO - if no current.Next
			current = current.Next
		}
		for i := 0; i < k; i++ {
			newCurrent = m[i]
			if i < k-1 {
				newCurrent.Next = m[i]
			}
		}
	}
	return newHead
}

func reverseKGroupOLD(head *ListNode, k int) *ListNode {
	var newHead *ListNode

	list := []*ListNode{}

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	revList := []*ListNode{}
	iteration := 1
	for {
		// if fewer than k left
		if k*iteration > len(list) {
			revList = append(revList, list[(k-1)*iteration:]...)
			break
		}

		// else
		for i := (k - 1) * iteration; i >= (iteration-1)*k; i-- {
			revList = append(revList, list[i])
		}
		iteration++
	}

	fmt.Println(revList)

	return newHead
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
