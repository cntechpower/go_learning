package main

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//380ms, 5.3MB
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := res
OUTER:
	for {
		min := math.MaxInt8
		minIdx := -1
		nilCount := 0
		for idx, node := range lists {
			if node == nil {
				nilCount++
				if nilCount == len(lists) {
					break OUTER
				}
				continue
			}
			if node.Val < min {
				min = node.Val
				minIdx = idx
			}
		}
		curr.Next = lists[minIdx]
		curr = curr.Next
		lists[minIdx] = lists[minIdx].Next
	}
	return res.Next
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

// 120ms, 5.3MB
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var res *ListNode
	for i := 0; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

// 8ms, 5.3MB
func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func main() {

}
