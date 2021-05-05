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
