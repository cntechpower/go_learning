package polymorphism_test

import "testing"

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func TestPolymorphism(t *testing.T) {
	goPro := new(GoProgrammer)
	javaPro := new(JavaProgrammer)
	t.Log(goPro.WriteHelloWorld())
	t.Log(javaPro.WriteHelloWorld())
}
