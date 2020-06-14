package testUtil

import (
	"fmt"
	"math/rand"
	"time"
)

func NewRandomList(length int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	}
	return res
}
func IsSorted(list []int, isAsc bool) bool {
	for i := 1; i < len(list); i++ {
		if (isAsc && list[i-1] > list[i]) || (!isAsc && list[i-1] < list[i]) {
			return false
		}
	}
	return true
}

func Print(list []int) {
	fmt.Println("-------------------------------------")
	for _, p := range list {
		fmt.Println(p)
	}
	fmt.Println("-------------------------------------")
}
