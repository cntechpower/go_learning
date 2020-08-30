package main

import (
	"fmt"
	"strconv"
)

//Definition for singly-linked list.
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++

	}
	tail.Next = head
	newTail := head
	for i := 0; i < length-(k%length)-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	fmt.Println(rotateRight(l1, 2))
	fmt.Println(rotateRight(l2, 4))
}
