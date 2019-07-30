package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5, 6} //数组,不可以进行append
	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 := []int{1, 2, 3, 4, 5, 6} //切片,长度可变,可以进行append
	slice2 := array1[0:]
	slice3 := append(slice1, 7)
	for k, v := range slice2 {
		fmt.Printf("Key %v: %v\n", k, v)
	}
	for k, v := range slice3 {
		fmt.Printf("Key %v: %v\n", k, v)
	}
	slice4 := array2[3:6]
	fmt.Printf("slice4 Len is %v,Cap is %v\n", len(slice4), cap(slice4)) //len=3, cap=5. cap等于[3:]
	slice5 := append(slice4, 9)
	fmt.Println(slice5)
	fmt.Println(array2)
	fmt.Printf("slice5 Len is %v,Cap is %v\n", len(slice5), cap(slice5))
	slice6 := append(slice5, 10)
	fmt.Println(slice6)
	fmt.Println(array2)
	fmt.Printf("slice6 Len is %v,Cap is %v\n", len(slice6), cap(slice6))
	slice7 := append(slice6, 11)
	fmt.Println(slice7)
	fmt.Println(array2)
	fmt.Printf("slice7 Len is %v,Cap is %v\n", len(slice7), cap(slice7))

}
