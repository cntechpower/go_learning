package main

import (
	"fmt"
	"syscall"
)

/*
type Rlimit struct {
	Cur uint64 //Soft limit
	Max uint64 //Hard limit (ceiling for rlim_cur)
}
//http://man7.org/linux/man-pages/man2/getrlimit.2.html
*/

func main() {
	name := syscall.RLIMIT_NOFILE
	var rlimit syscall.Rlimit
	err := syscall.Getrlimit(name, &rlimit)
	if nil != err {
		panic(err)
	}
	fmt.Println(rlimit.Cur)
	fmt.Println(rlimit.Max)
	rlimit.Max = 65535
	rlimit.Cur = 65535
	syscall.Setrlimit(name, &rlimit)
	fmt.Println(rlimit.Cur)
	fmt.Println(rlimit.Max)
}
