package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op } //add value to map
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{} //init value is false
	mySet[1] = true
	for i := 0; i < 10; i++ {
		if mySet[i] {
			t.Logf("%d is exist", i) //判断mySet[i]是否存在
		} else {
			t.Logf("%d is not exist", i)
		}
	}
	delete(mySet, 1)
	for i := 0; i < 10; i++ {
		if mySet[i] {
			t.Logf("%d is exist", i)
		} else {
			t.Logf("%d is not exist", i)
		}
	}
}
