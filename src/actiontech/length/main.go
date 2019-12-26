package main

import "fmt"

func main() {
	m := make(map[string]bool, 0)
	fmt.Println(len(m))
	m1 := make(map[string]bool, 1)
	fmt.Println(len(m1))
	s1 := make([]string, 0)
	fmt.Println(len(s1))
	s2 := make([]string, 1)
	fmt.Println(len(s2))
	fmt.Println(".............")
	fmt.Println(len(returnNilMap()))
	fmt.Println(len(returnNilSlice()))
}

func returnNilMap() map[string]bool {
	return nil
}

func returnNilSlice() []string {
	return nil
}
