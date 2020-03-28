package main

import "fmt"

func main() {
	m := make(map[int]int, 0)
	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}
	for k := range m {
		if k == 5 {
			m[100] = 1000
		}
		delete(m, k) //m[100] won't be delete
		fmt.Println(m)
	}
}
