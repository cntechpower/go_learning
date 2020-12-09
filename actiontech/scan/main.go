package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Press Enter to Continue, q/Q to Exit")
	choice := ""
	fmt.Scanln(&choice)
	if strings.ToUpper(choice) == "Q" {
		os.Exit(0)
	} else {
		time.Sleep(5 * time.Second)
	}
}
