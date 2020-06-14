package testUtil

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const DefaultListLength = 10

func NewRandomList(length int) []int {
	res := make([]int, length)
RETRY:
	for i := 0; i < length; i++ {
		res[i] = rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	}
	// ensure list is not sorted.
	if checkSorted(res, true, true) || checkSorted(res, false, true) {
		goto RETRY
	}
	return res
}
func checkSorted(list []int, isAsc bool, silence bool) bool {
	for i := 1; i < len(list); i++ {
		if (isAsc && list[i-1] > list[i]) || (!isAsc && list[i-1] < list[i]) {
			if !silence {
				printArray(list, i-1, i)
				fmt.Println("Failed")
			}
			return false
		}
	}
	fmt.Println("OK")
	return true
}

func printArray(list []int, wrongIdx1, wrongIdx2 int) {
	fmt.Println("-------------------------------------")
	for idx, p := range list {
		s := strconv.Itoa(p)
		if idx == wrongIdx1 || idx == wrongIdx2 {
			s += " <--"
		}
		fmt.Println(s)
	}
	fmt.Println("-------------------------------------")
}

func Wrapper(sortFunc func([]int, bool)) {
	l1 := NewRandomList(DefaultListLength)
	sortFunc(l1, true)
	checkSorted(l1, true, false)
	l2 := NewRandomList(DefaultListLength)
	sortFunc(l2, false)
	checkSorted(l2, false, false)

}
