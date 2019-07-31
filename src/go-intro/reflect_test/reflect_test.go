package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Employee struct {
	id   int
	name string `format:"normal"`
}

func (e *Employee) UpdateId(newId int) {
	e.id = newId
}
func TestTypeAndValue(t *testing.T) {
	var f int = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	//t := reflect.TypeOf(v)

}

func TestInvokeByName(t *testing.T) {
	e := &Employee{1, "test1"}
	t.Logf("Name---Value: %[1]v, Type: %[1]T", reflect.ValueOf(*e).FieldByName("name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("name"); !ok {
		t.Errorf("Failed to get Name Fileld")
	} else {
		t.Log("Tag--format: ", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateId").
		Call([]reflect.Value{reflect.ValueOf(10)})
	t.Log("Updated Age", e)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//	t.Log(a == b)
	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Employee{1, "Mike"}
	c2 := Employee{1, "Mike"}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}
