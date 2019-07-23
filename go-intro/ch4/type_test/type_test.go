package type_test

import "testing"

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{"0", "dujinyang", 20}
	e1 := Employee{Name: "jinyang", Age: 30}
	e2 := new(Employee)
	t.Log(e)
	t.Log(e1)
	t.Log(*e2)
	e.Name = "Name1"
	e1.Name = "Name2"
	e2.Name = "Name3"
	t.Log(e)
	t.Log(e1)
	t.Log(*e2)
	t.Logf("e is %T", e1)
	t.Logf("e2 is %T", e2)
}
