package map_test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	// m1 = append(m1, "5:5") //first argument to append must be slice; have map[int]int
	m2 := map[int]int{}
	t.Logf("len m2 is %d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3 is %d", len(m3))
}

func TestAccessMapNotExist(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //Init Value is 0
	m1[2] = 0
	t.Log(m1[2]) //m1[2] is 0,无法区分不存在和0值.
	if v, ok := m1[1]; !ok {
		t.Log("Key 1 is Not Exist")
	} else {
		t.Logf("Key 1 is Exist,value is %d", v)
	}

}

func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m1 {
		t.Logf("Key %d : %d", k, v)
	}
}
