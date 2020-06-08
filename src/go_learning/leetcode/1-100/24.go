package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	first.Next = swapPairs(second.Next)
	second.Next = first
	return second

}

func main() {

}
