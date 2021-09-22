package leetcode

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var builder strings.Builder
	for l != nil {
		fmt.Fprintf(&builder, "%d", l.Val)
		l = l.Next
	}
	return builder.String()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	b := strings.Builder{}
	current := []*TreeNode{t}
	for len(current) > 0 {
		var next []*TreeNode
		for _, node := range current {
			// fmt.Println(node == nil)
			if node == nil {
				b.WriteString("nil,")
			} else {
				b.WriteString(fmt.Sprintf("%d,", node.Val))
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
	}
	return strings.TrimRight(b.String(), ",")
}
