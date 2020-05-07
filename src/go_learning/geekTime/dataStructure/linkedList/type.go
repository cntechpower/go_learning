package main

import (
	"fmt"
	"strings"
)

type SingleLinkedListNode struct {
	next  *SingleLinkedListNode
	value interface{}
}

type SingleLinkedList struct {
	head *SingleLinkedListNode
	len  int64
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{head: &SingleLinkedListNode{
		next:  nil,
		value: "THIS_IS_FAKE_HEAD",
	}, len: 0}
}

func (l *SingleLinkedList) findPrevNode(node *SingleLinkedListNode) *SingleLinkedListNode {
	var prevNode *SingleLinkedListNode
	for n := l.head; n != nil; n = n.next {
		if n.next == node {
			prevNode = n
		}
	}
	return prevNode
}

func (l *SingleLinkedList) InsertAfter(node *SingleLinkedListNode, value interface{}) bool {
	if node == nil {
		return false
	}
	newNode := &SingleLinkedListNode{
		next:  node.next,
		value: value,
	}
	node.next = newNode
	l.len++
	return true
}

func (l *SingleLinkedList) InsertBefore(node *SingleLinkedListNode, value interface{}) bool {
	if node == nil || node == l.head {
		return false
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return false
	}
	newNode := &SingleLinkedListNode{
		next:  node,
		value: value,
	}
	prevNode.next = newNode
	l.len++
	return true
}

func (l *SingleLinkedList) InsertHead(value interface{}) bool {
	return l.InsertAfter(l.head, value)
}

func (l *SingleLinkedList) InsertTail(value interface{}) bool {
	lastNode := l.findPrevNode(nil)
	return l.InsertAfter(lastNode, value)
}

func (l *SingleLinkedList) Delete(node *SingleLinkedListNode) bool {
	if node == nil || node == l.head {
		return false
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return false
	}
	prevNode.next = node.next
	node = nil
	l.len--
	return true
}

func (l *SingleLinkedList) DeleteHead() bool {
	return l.Delete(l.head.next)
}

func (l *SingleLinkedList) DeleteTail() bool {
	return l.Delete(l.findPrevNode(nil))
}

func (l *SingleLinkedList) String() string {
	builder := strings.Builder{}
	builder.WriteString("{")
	for n := l.head.next; n != nil; n = n.next {
		builder.WriteString(fmt.Sprintf("%v", n.value))
		if n.next != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}
