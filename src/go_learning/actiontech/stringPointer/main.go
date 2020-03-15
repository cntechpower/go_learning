package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "./tmp"
	absPath, err := filepath.Abs(path)
	fmt.Printf("path: %v , error: %v\n", absPath, err)
}
