package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	requestList := []string{"", "hello"}
	for _, v := range requestList {
		response, err := echo(v)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Response is %s\n", response)
		}
	}
}
