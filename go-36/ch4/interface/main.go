package main

import "fmt"

//Pet test
type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

//Dog test
type Dog struct {
	name string
}

//SetName test
func (dog *Dog) SetName(name string) {
	dog.name = name
}

//Name test
func (dog Dog) Name() string {
	return dog.name
}

//Category test
func (dog Dog) Category() string {
	return "dog"
}
func main() {
	dog := Dog{"little dog"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

	var pet Pet = &dog
	fmt.Printf("This pet is a %s,the name is %q\n", pet.Category(), pet.Name()) //To double-quote strings as in Go source, use %q ("little dog")

}
