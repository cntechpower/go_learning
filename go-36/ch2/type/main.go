package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	value, ok := interface{}(container).(map[int]string)
	if ok {
		fmt.Printf("OK,The element is %q\n", value[1])
	} else {
		fmt.Printf("Not OK,The element is %q\n", container[1])
	}
}
