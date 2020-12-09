package funcAsParameter

import (
	"fmt"
	"testing"
)

type filterInt func(int) bool

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func filter(s []int, f filterInt) []int {
	var result []int
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func TestFuncParameter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(filter(s, isEven))
	fmt.Println(filter(s, isOdd))

}
