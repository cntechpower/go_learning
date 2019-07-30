package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("String", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}
func TestInterfaceType(t *testing.T) {
	DoSomething("123")
	DoSomething(1)
}

//倾向于使用小的接口定义,很多接口只包含一个方法
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

//较大的接口可以由多个小接口组合而成
type ReadWriter interface {
	Reader
	Writer
}

//只依赖与必要功能的最小接口
//func StoreData(Reader Reader) error
