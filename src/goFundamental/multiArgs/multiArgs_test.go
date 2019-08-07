package multiArgs

import (
	"fmt"
	"reflect"
	"testing"
)

func myFunc(arg ...int) {
	fmt.Println(reflect.TypeOf(arg)) //变参的类型为slice
	for _, v := range arg {
		fmt.Println(v)
	}
}

func TestMultiArgs(t *testing.T) {
	myFunc(1, 2, 3, 4)

}
