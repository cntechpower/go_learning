package utils

import "fmt"

func init() {
	fmt.Println("Called utils's init once")
}

func init() {
	fmt.Println("Called utils's init twice")
}

//LowerCase test
func LowerCase() string {
	return "Called LowerCase"
}
