package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Name string `json:"Name"`
	Sex  bool   `json:"Sex"`
}

func main() {
	u := user{
		Name: "dujinyang",
		Sex:  false,
	}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s : %v\n", t.Field(i).Tag.Get("json"), v.Field(i))
	}
}
