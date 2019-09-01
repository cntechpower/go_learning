package array_test

import (
	"testing"
)

func TestArray(t *testing.T) {
	var a [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 4, 5}
	t.Log(a[1], a[2])
	t.Log(arr1, arr2)
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}
	for idx, value := range arr1 {
		t.Logf("Index %v Value is %v", idx, value)
	}
}

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s1 := append(s0, 1)
	t.Log(len(s0), cap(s0)) //s0 still 0 0
	t.Log(len(s1), cap(s1)) //s1 1,1
	s2 := make([]int, 3, 5) //cap代表容量,length代表着初始化的数量
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 999)
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

//cap *2
func TestSliceGrow(t *testing.T) {
	//s := []int{}
	s := make([]int, 3, 5)
	t.Logf("Type is %T", s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShare(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Logf("Type is %T", year)
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
}

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("Len m1=%d", len(m1))
	m2 := map[string]int{"test1": 1, "test2": 3}
	t.Log(m2["test2"])
	t.Logf("Len m2=%d", len(m2))
	m3 := map[int]string{1: "dujinyang", 2: "xiehongya"}
	t.Log(m3[1])
	t.Logf("Len m3=%d", len(m3))
	m4 := map[int]int{}
	m4[4] = 16
	t.Logf("Len m4=%d", len(m4))
	m5 := make(map[int]int, 10)
	t.Logf("Len m5=%d", len(m5))
	if v, ok := m1[5]; ok {
		t.Logf("Exist,Var is %d", v)
	} else {
		t.Logf("Not Exist,Var is %d", v)
	}
	for k, v := range m1 {
		t.Logf("Key %v: %v", k, v)
	}
}

func TestMapFunc(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}
