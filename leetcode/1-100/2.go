package main

import (
	"fmt"
	"math"
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

func addTwoNumbersWrong(l1 *ListNode, l2 *ListNode) *ListNode {

	//  输入:
	//	[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	//	[5,6,4]
	//	输出
	//	[-2,-4,-3,-5,-7,-7,-4,-5,-8,-6,-3,0,-2,-7,-3,-3,-2,-2,-9]
	//	预期结果
	//	[6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// Overflow
	total := 0
	n := 0
	for ll1 := l1; ll1 != nil; ll1 = ll1.Next {
		total += ll1.Val * int(math.Pow10(n))
		n++
	}
	n = 0
	for ll2 := l2; ll2 != nil; ll2 = ll2.Next {
		total += ll2.Val * int(math.Pow10(n))
		n++

	}
	res := &ListNode{}
	now := res
	for i := total; i != 0; i = i / 10 {
		now.Val = i % 10
		if i/10 != 0 {
			now.Next = &ListNode{}
			now = now.Next
		}
	}
	return res
}

func addTowNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	p := l1
	q := l2
	overflow := 0
	for p != nil || q != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		x := 0
		y := 0
		if p != nil {
			x = p.Val
			p = p.Next
		}
		if q != nil {
			y = q.Val
			q = q.Next
		}
		sum := x + y + overflow
		curr.Val = sum % 10
		overflow = sum / 10
	}
	if overflow > 0 {
		curr.Next = &ListNode{Val: overflow}
	}
	return res.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(addTowNumbers(l1, l2))
}
