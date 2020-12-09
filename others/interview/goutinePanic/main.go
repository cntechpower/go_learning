package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {

		for range time.NewTicker(time.Second).C {
			defer func() {
				recover()
			}()
			proc()
			fmt.Println(time.Now().Format(time.RFC3339))
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
