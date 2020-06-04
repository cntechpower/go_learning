package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	res := ""
	for l1 := l; l1 != nil; l1 = l1.Next {
		res += strconv.Itoa(l1.Val)
	}
	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		} else {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		curr.Next = l1
		curr = curr.Next
		l1 = l1.Next
	}
	if l2 != nil {
		curr.Next = l2
		curr = curr.Next
		l2 = l2.Next
	}
	return head.Next

}
func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(l1, l2))
}
