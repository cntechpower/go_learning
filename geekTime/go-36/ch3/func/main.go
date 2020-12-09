package main

import "fmt"

type printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesCount int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p printer
	p = printToStd
	p("test")
}
