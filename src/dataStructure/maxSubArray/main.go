package main

import "fmt"

func maxSubArraySum1(a []int) (int, int, int) {
	var thisSum, maxSum int
	var begin, end int
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			thisSum = 0
			for k := i; k <= j; k++ {
				thisSum += a[k]
				if thisSum > maxSum {
					maxSum = thisSum
					begin = i
					end = k
				}
			}
		}

	}
	return maxSum, begin, end
}

func maxSubArraySum2(a []int) (int, int, int) {
	var thisSum, maxSum int
	var begin, end int
	for i := 0; i < len(a); i++ {
		thisSum = 0
		for j := i; j < len(a); j++ {
			thisSum += a[j]
			if thisSum > maxSum {
				maxSum = thisSum
				begin = i
				end = j
			}
		}
	}
	return maxSum, begin, end
}

func maxSubArraySum3(a []int) (int, int, int) {
	if len(a) == 1 {
		return a[0], 0, 0
	}
	mid := len(a) / 2
	left, _, _ := maxSubArraySum3(a[:mid])
	right, _, _ := maxSubArraySum3(a[mid:])
	if left+right > left {
		return left + right, 0, 0
	}
	return left, 0, 0
}

//https://blog.csdn.net/samjustin1/article/details/52043173
func main() {
	a := []int{1, 10, 20, -24, 10}
	fmt.Println(maxSubArraySum1(a))
	fmt.Println(maxSubArraySum2(a))
	fmt.Println(maxSubArraySum3(a))
}
