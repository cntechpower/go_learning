package slice_test

import "testing"

func TestSlice(t *testing.T) {
	var slice1 []int
	t.Log(len(slice1), cap(slice1))
	slice1 = append(slice1, 1)
	t.Log(len(slice1), cap(slice1))
	var slice2 = []int{1, 2, 3, 4}
	t.Log(len(slice2), cap(slice2))
	var slice3 = make([]int, 3, 5)
	t.Log(len(slice3), cap(slice3))
	// t.Log(slice3[0], slice3[1], slice3[2], slice3[3]) // slice3[4] is not initialized
	slice3 = append(slice3, 1)
	t.Log(slice3[0], slice3[1], slice3[2], slice3[3])
}

func TestSLiceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i) //地址发生了变化,会把原本的数据拷贝到新的位置
		t.Log(len(s), cap(s))
	}
}

func TestSliceShart(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "test"
	t.Log(Q2, len(Q2), cap(Q2))
}
