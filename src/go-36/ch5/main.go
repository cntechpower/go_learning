package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "dujinyang", "Name for Print")
}

func hello(name string) (s string, err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}
	s = fmt.Sprintf("hello, %s", name)
	return

}

func introduce() string {
	return "introduce"
}

func main() {
	flag.Parse()
	s, err := hello(name)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(s, introduce())
}
