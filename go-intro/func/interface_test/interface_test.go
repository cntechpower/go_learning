package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProGrammer struct {
}

func (g *GoProGrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	// var p Programmer
	// t.Logf("%T", p)
	p := new(GoProGrammer)
	t.Logf("%T", p)
	t.Log(p.WriteHelloWorld())
}
