package error_handle_test

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("n should bigger than 2")
var LessThanZeroError = errors.New("n should bigger than 0")

func getFibonacci(n int) ([]int, error) {
	if n <= 0 {
		return nil, LessThanZeroError
	}
	if n <= 2 {
		return nil, LessThanTwoError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

func TestFib(t *testing.T) {
	if fibList, err := getFibonacci(10); err == LessThanTwoError {
		t.Logf("LessThanTwoError Error: %v", err)
	} else if err == LessThanZeroError {
		t.Logf("LessThanZeroError Error: %v", err)
	} else {
		t.Log(fibList)
	}

	//Method 2
	fibList, err := getFibonacci(1)
	if err == LessThanTwoError {
		t.Logf("LessThanTwoError Error: %v", err)
	} else if err == LessThanZeroError {
		t.Logf("LessThanZeroError Error: %v", err)
	} else {
		t.Log(fibList)
	}
}
