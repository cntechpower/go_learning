package interfaceTest

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "HelloWorld from GoProgrammer"
}

type CppProgrammer struct {
}

func (p *CppProgrammer) WriteHelloWorld() string {
	return "HelloWorld from CppProgrammer"
}

func TestInterface(t *testing.T) {
	var programmers []Programmer
	programmers = append(programmers, &GoProgrammer{})
	programmers = append(programmers, &CppProgrammer{})
	for _, programmer := range programmers {
		t.Log(programmer.WriteHelloWorld())
	}
}
