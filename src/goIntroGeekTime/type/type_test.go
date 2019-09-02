package _type_test

import (
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(b)

}

type Person struct {
	Name string
	Age  int
}

func TestPtr(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	person := new(Person)
	person.Name = "dujinyang"
	person.Age = 20
	t.Logf("%T  %T", a, aPtr)
	reflect.TypeOf(person)
	t.Logf("person format : %T", person)
}
