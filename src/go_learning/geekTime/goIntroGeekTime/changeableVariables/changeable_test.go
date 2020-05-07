package main

import (
	"testing"
)

func changeableSum(ops ...int) int {
	sum := 0
	for _, v := range ops {
		sum += v
	}
	return sum
}

func TestChangeableVariables(T *testing.T) {
	variableArray := []int{0, 1, 2, 3, 4, 5}
	T.Logf("variableArray is %v\n", variableArray)
	//T.Logf("variableArray... is %v\n", variableArray...)
	T.Logf("Got Sum of variableArray: %v\n", changeableSum(variableArray...))
}
