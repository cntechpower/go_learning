package main

import "fmt"

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
}
