package main

import (
	"container/list"
	"fmt"
)

//Stack test
type Stack struct {
	list *list.List
}

func newStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (s *Stack) push(v interface{}) {
	s.list.PushBack(v)
}

func (s *Stack) pop() interface{} {
	if s.len() == 0 {
		return nil
	}
	back := s.list.Back()
	if back != nil {
		s.list.Remove(back)
	}
	return back.Value

}
func (s *Stack) isEmpty() bool {
	if s.list.Len() == 0 {
		return true
	}
	return false
}

func (s *Stack) len() int {
	return s.list.Len()
}

func main() {
	s := newStack()
	fmt.Printf("Stack isEmpty: %v, Length is %v\n", s.isEmpty(), s.len())
	for i := 1; i < 10; i++ {
		s.push(i)
		fmt.Printf("Push %d to Stack,Stack isEmpty: %v, Length is %v\n", i, s.isEmpty(), s.len())
	}
	for i := 1; i < 10; i++ {
		num := s.pop()
		fmt.Printf("Pop %v from Stack,Stack isEmpty: %v, Length is %v\n", num, s.isEmpty(), s.len())
	}

	s1 := newStack()
	str1 := "{})}"
	fmt.Printf("Test Char is %v\n", str1)
	for k, v := range str1 {
		fmt.Printf("No %v character is %c\n", k, v)
		char := fmt.Sprintf("%c", v)
		switch char {
		case "{":
			s1.push(char)
		case "(":
			s1.push(char)
		case "}":
			charOut := s1.pop()
			if charOut == nil {
				fmt.Printf("Short Of }\n")
				break
			}
			if charOut != "{" {
				fmt.Printf("pop is %v, Exceped is {\n", charOut)
			}
		case ")":
			charOut := s1.pop()
			if charOut == nil {
				fmt.Printf("Short Of (\n")
				break
			}
			if charOut != "(" {
				fmt.Printf("pop is %v, Exceped is (\n", charOut)
			}
		}
	}
}
