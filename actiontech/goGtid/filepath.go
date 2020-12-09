package main

import (
	"fmt"
	"path/filepath"
)

func main1() {
	d := filepath.Dir("/tmp/test")
	fmt.Println(d)
}
