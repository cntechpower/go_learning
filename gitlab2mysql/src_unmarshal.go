package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"time"
)

type issueStruct struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	var json_data []issueStruct
	json_src_read, err := ioutil.ReadFile("issue.src")
	if err != nil {
		panic(err)
	}
	//json_src_string := string(json_src_read)
	//fmt.Printf("Json is :\n %s", json_src_string)
	err = json.Unmarshal(json_src_read, &json_data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", json_data)

}
