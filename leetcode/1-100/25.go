package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = myReverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head
	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return tail, head
}
