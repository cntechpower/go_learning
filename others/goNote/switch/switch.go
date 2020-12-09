package main

import (
	"fmt"
)

func main() {
	a := "test"
	switch a {
	case "test":
		fmt.Println("got test")
		fallthrough //run next case without matching
	case "test1":
		fmt.Println("got test1")
		fallthrough
	default:
		fmt.Println("got default")
		//fallthrough //cannot fallthrough final case in switch
	}
}
