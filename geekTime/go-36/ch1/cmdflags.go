package main

import (
	"flag"
	"fmt"
	"os"
)

var name2 string

// func init() {
// 	cmdLine := flag.NewFlagSet("question", flag.ExitOnError)
// 	cmdLine.StringVar(&name, "name", "dujinyang", "Name to Print")
// }
func main() {
	cmdLine := flag.NewFlagSet("question", flag.ExitOnError)
	cmdLine.StringVar(&name2, "name2", "dujinyang", "Name to Print")
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello : %v\n", name2)
}
