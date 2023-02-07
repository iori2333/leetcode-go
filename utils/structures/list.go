package structures

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func NewList(sequence []int) *ListNode {
	head := NewListNode(0)
	node := head
	for _, val := range sequence {
		node.Next = NewListNode(val)
		node = node.Next
	}
	return head.Next
}

func (head *ListNode) Sequence() []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func (head *ListNode) String() string {
	seq := head.Sequence()
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range seq {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(v))
	}
	sb.WriteString("]")
	return sb.String()
}
