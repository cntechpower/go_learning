package oop_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s, Name:%s, Age:%d", e.Id, e.Name, e.Age)
}

// func (e Employee) String() string { //对象的地址不同,整个结构被复制了一份,不推荐这种方式
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s, Name:%s, Age:%d", e.Id, e.Name, e.Age)
// }

func TestCreateEmployee(t *testing.T) {
	e1 := Employee{"1", "test1", 20}
	e2 := Employee{Name: "test2", Age: 21}
	e3 := new(Employee)
	e1.Id = "2"
	e2.Age = 22
	t.Log(e1.String())
	t.Log(e2)
	t.Log(e3)
	t.Logf("%T", e1)
	t.Logf("%T", e2)
	t.Logf("%T", e3)

}

func TestStructOper(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Logf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
