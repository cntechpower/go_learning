package main

import (
	"fmt"
	"reflect"
)

func binarySearch(src []int, target int) (idx int, err error) {
	if src == nil {
		return -1, fmt.Errorf("source is empty")
	}
	left := 0
	right := len(src) - 1
	mid := (left + right) / 2
	count := 0
	fmt.Printf("Init state : left: %v, right: %v, mid: %v \n", left, right, mid)
	for left <= right {
		if target == src[mid] {
			return mid, nil
		} else if target > src[mid] {
			fmt.Printf("Move left from %v to %v\n", left, mid)
			left = mid + 1
		} else {
			fmt.Printf("Move right from %v to %v\n", right, mid)
			right = mid - 1
		}
		count++
		fmt.Printf("After %v Exec, left: %v, right: %v, mid: %v \n", count, left, right, mid)
		mid = (left + right) / 2

	}
	return -1, fmt.Errorf("%v not found", target)

}

func transToInterfaceSlice(arg interface{}) (out []interface{}, ok bool) {
	if reflect.ValueOf(arg).Kind() != reflect.Slice {
		return nil, false
	}
	val := reflect.ValueOf(arg)
	for i := 0; i < val.Len(); i++ {
		out = append(out, val.Index(i).Interface())
	}
	return out, true
}

func main() {
	srcArray := make([]int, 0)
	for i := 1; i < 100001; i++ {
		srcArray = append(srcArray, i)
	}
	target := 99999
	// val, ok := transToInterfaceSlice(srcArray)
	// if !ok {
	// 	fmt.Println("transToInterface error")
	// }
	idx, err := binarySearch(srcArray, target)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	} else {
		fmt.Printf("found %v, index is %v, index value is %v\n", target, idx, srcArray[idx])
	}

}
