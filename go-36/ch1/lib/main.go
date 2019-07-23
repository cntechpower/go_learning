package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "default", "Default Print Name")
}

func main() {
	flag.Parse()
	hello(name)
}
