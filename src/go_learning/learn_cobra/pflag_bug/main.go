package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var boolArg bool

func main() {
	pflag.BoolVar(&boolArg, "b", false, "test bool arg")
	pflag.Parse()
	fmt.Println(boolArg)
}
