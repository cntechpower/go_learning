package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Input URL: ")
	url, err := inputReader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")
	checkError(err)
	res, err := http.Get(url)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
