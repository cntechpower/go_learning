package b

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	test1()
	test2()
	test3()
	test4()
}

func test1() {
	var out [][]int
	for _, i := range [][]int{{1, 1}, {2, 2}, {3, 3}} {
		out = append(out, i[1:2])
	}
	fmt.Println(out)
}

func test2() {
	var out [][]int
	for _, i := range [][2]int{{1, 1}, {2, 2}, {3, 3}} {
		out = append(out, i[1:2])
		//fmt.Println(out)
	}
	fmt.Println(out)
}

func test3() {
	var out [][]int
	for _, i := range [][2]int{{1, 1}, {2, 2}, {3, 3}} {
		tmpI := i
		out = append(out, tmpI[1:2])
	}
	fmt.Println(out)
}

func test4() {
	var out [][]int
	for _, i := range [][2]int{{1, 1}, {2, 2}, {3, 3}} {
		tmpI := i[1:2]
		//fmt.Printf("%p\n", tmpI)
		out = append(out, tmpI)
	}
	//for _, v := range out {
	//	fmt.Printf("%p\n", v)
	//}
	fmt.Println(out)

}
