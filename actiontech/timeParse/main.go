package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	time1, err := time.Parse("060102 15:04:05", "170323 15:29:40")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time1.Format(time.RFC3339))
	type s struct {
		ok *bool
	}
	s1 := &s{ok: nil}

	if *s1.ok {
		fmt.Println("ok")
	}

}
