package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var name1 string

func init() {
	flag.StringVar(&name1, "name1", "dujinyang", "Name to Print") //这个参数定义不会生效,因为在下面flag.CommandLine被重新定义了.
	// flag.Usage = func() {
	// 	fmt.Printf("Usage of Hello World:\n")
	// 	flag.PrintDefaults()
	// }
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "dujinyang", "Name to Print") //这个参数定义会生效,因为flag.CommandLine在之后没有重新定义过.

}
func main1() {
	// flag.StringVar(&name, "name", "dujinyang", "Name to Print")
	flag.Parse()
	fmt.Printf("Hello %v\n", name)

}
