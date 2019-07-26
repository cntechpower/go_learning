package main

import "fmt"

func main() {
	var count int
	source := []int{1, 7, 6, 5, 4, 3, 2, 8}
	fmt.Printf("Source is: %v\n", source)
	sourceLen := len(source)
	fmt.Printf("Source Length is %v\n", sourceLen)
	for j := 0; j < sourceLen; j++ {
		//for i := 0; i < sourceLen-1; i++ { //最后j个元素已经有序, 不需要再次循环.此种方式需要循环56次
		for i := 0; i < sourceLen-j-1; i++ { //此种方式只需要循环28次
			if source[i] > source[i+1] {
				source[i+1], source[i] = source[i], source[i+1]
			}
			count++
			fmt.Printf("Afer %d Swap, Source is %v\n", count, source)
		}
	}
}
