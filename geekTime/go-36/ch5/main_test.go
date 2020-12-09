package main

import "testing"

func TestHello(t *testing.T) {
	var name string
	s, err := hello(name)
	if err == nil {
		t.Errorf("err is nil, but it not should be")
	}
	if s != "" {
		t.Errorf("s is not empty, but it not should be")
	}
	name = "test"
	s, err = hello(name)
	if err != nil {
		t.Errorf("err is not nil, but it not should be")
	}
	if s == "" {
		t.Errorf("s is empty, but it not should be")
	}
	t.Logf("TestHello Succeed")
}

func TestIntroduce(t *testing.T) {
	expected := "introduce"
	s := introduce()
	if s != expected {
		//t.Errorf("s is not as expected") //Not Exit, 等同于t.Log+t.Fail
		t.Fatalf("s is not as expected") //Exit,等同于 t.Log+t.FailNow
		//t.FailNow()                      //Manual Exit
	}
	t.Logf("TestIntroduce Succeed")
}
