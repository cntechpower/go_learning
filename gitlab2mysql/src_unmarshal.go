package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//type DictionaryType map[string]interface{}
type issueData struct {
	Id          int    `json:"id"`
	Title       string `json:"title""`
	Description string `json:"description"`
}

func main() {
	var json_data issueData
	json_src_read, err := ioutil.ReadFile("issue.src")
	if err != nil {
		panic(err)
	}
	//json_src_string := string(json_src_read)
	//fmt.Printf("Json is :\n %s", json_src_string)
	err1 := json.Unmarshal(json_src_read, &json_data)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("%v\n", json_data)

}
