package main

import "fmt"

func test1() {
	var numbers1 []int
	printSlice(numbers1)
	if numbers1 == nil {
		fmt.Printf("切片是空的\n")
	}
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func test2() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)
	fmt.Println("numbers==", numbers)
	fmt.Println("numbers[:3]==", numbers[:3])
	numbers2 := make([]int, 0, 5)
	printSlice(numbers2)
	numbers3 := numbers[:2]
	printSlice(numbers3)
}
func main() {
	var numbers4 []int
	printSlice(numbers4)
	numbers4 = append(numbers4, 0)
	printSlice(numbers4)
	numbers4 = append(numbers4, 1)
	printSlice(numbers4)
	numbers4 = append(numbers4, 2, 3, 4)
	printSlice(numbers4)
	numbers5 := make([]int, len(numbers4), (cap(numbers4))*2)
	copy(numbers5, numbers4)
	printSlice(numbers5)
}
