package main

import "fmt"

// func modifyArray(a *[3]string) *[3]string { //这种方式会修改a的值
// 	a[1] = "X"
// 	return a
// }

// func main() {
// 	array1 := [3]string{"a", "b", "c"}
// 	array2 := modifyArray(&array1)
// 	fmt.Println(array1)
// 	fmt.Println(*array2)

// }

// func modifyArray(a [3]string) [3]string { //这种方式不会修改a的值
// 	a[1] = "X"
// 	return a
// }
// func main() {
// 	array1 := [3]string{"a", "b", "c"}
// 	array2 := modifyArray(array1)
// 	fmt.Println(array1)
// 	fmt.Println(array2)
// }

func modifyArray(a [3][]string) [3][]string {
	//a[1][1] = "x" //会影响到a
	a[1] = []string{"x"} //不会影响到a
	return a
}

func main() {
	array1 := [3][]string{[]string{"a", "b", "c"}, []string{"d", "e", "f"}, []string{"g", "h", "i"}}
	array2 := modifyArray(array1)
	fmt.Println(array1)
	fmt.Println(array2)
}
