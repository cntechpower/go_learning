package main

import (
	"fmt"
	"os"
)

func main() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/usr/sbin/ip", []string{"ip", "a"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)
}
