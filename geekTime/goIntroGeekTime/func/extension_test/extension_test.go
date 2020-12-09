package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}
func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(name)
}

type Dog struct {
	*Pet //匿名类型,类似继承.
}

func (d *Dog) Speak() {
	fmt.Println("wang~~~ wang")
}

func (d *Dog) SpeakTo(name string) {
	d.Speak()
	fmt.Println(name)
}

func TestExtend(t *testing.T) {
	// d := new(Dog)
	d := Dog{}
	d.Speak()
	d.SpeakTo("dujinyang")
}
