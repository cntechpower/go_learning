package main

import (
	"fmt"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	res := ""
	for l1 := l; l1 != nil; l1 = l1.Next {
		fmt.Println(l1.Val)
	}
	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	h := head
	nodeToDelPrev := &ListNode{}
	nodeToDel := head
	currIdx := 0
	begin := head
	for begin != nil {
		begin = begin.Next
		currIdx++
		if currIdx > n {
			nodeToDelPrev = nodeToDel
			nodeToDel = nodeToDel.Next
		}
	}
	if nodeToDel == head {
		return head.Next
	}
	nodeToDelPrev.Next = nodeToDel.Next
	return h
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	slow := dummyHead
	quick := dummyHead
	for quick.Next != nil {
		quick = quick.Next
		if n == 0 {
			slow = slow.Next
		} else {
			n--
		}
	}
	fmt.Printf("慢指针: %v 快指针: %v\n", slow.Val, quick.Val)
	delNode := slow.Next
	slow.Next = delNode.Next
	return dummyHead.Next
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
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	fmt.Println(removeNthFromEnd(l1, 2))
	fmt.Println(removeNthFromEnd(l2, 2))
	fmt.Println(removeNthFromEnd2(l1, 2))
	fmt.Println(removeNthFromEnd2(l2, 2))
}
