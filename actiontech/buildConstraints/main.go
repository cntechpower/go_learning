package main

import "fmt"

func main() {
	fmt.Println(test())
}

/*
➜  buildConstraints git:(master) GOPATH=`pwd` go build -tags rel .
➜  buildConstraints git:(master) ./buildConstraints
i am rel
➜  buildConstraints git:(master) GOPATH=`pwd` go build -tags test .
➜  buildConstraints git:(master) ./buildConstraints
i am test
*/
