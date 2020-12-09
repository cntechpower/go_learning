package main

import "fmt"

func isPalindrome(l *SingleLinkedList) bool {
	if l.len == 0 || l.head.next == nil {
		return true
	}
	slow := l.head
	fast := l.head
	var prev *SingleLinkedListNode
	var tmp *SingleLinkedListNode
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		tmp = slow.next
		slow.next = prev
		prev = slow
		slow = tmp
	}
	if fast != nil {
		slow = slow.next
	}
	for prev != nil && slow != nil && prev.value == slow.value {
		prev = prev.next
		slow = slow.next
	}

	isPalindrome := prev == nil && slow == nil
	slow = l.head
	fast = l.head
	prev = nil
	tmp = nil
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		tmp = slow.next
		slow.next = prev
		prev = slow
		slow = tmp
	}
	return isPalindrome

}

func main() {
	l := NewSingleLinkedList()
	l.InsertHead("1")
	l.InsertHead("2")
	l.InsertHead("3")
	fmt.Println(l) //{3, 2, 1}
	l.InsertTail("4")
	fmt.Println(l) //{3, 2, 1, 4}
	l.DeleteHead()
	fmt.Println(l) //{2, 1, 4}
	l.DeleteTail()
	fmt.Println(l) //{2, 1}
	fmt.Println(l.FindNodeByIndex(1))

	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewSingleLinkedList()
		for _, c := range str1 {
			l.InsertTail(string(c))
		}
		fmt.Printf("checking %v\n", l)
		fmt.Printf("checked isPalindrome for %v -- %v\n", l, isPalindrome(l))
	}
}
