package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var protect uint

func init() {
	flag.UintVar(&protect, "protect", 1, "weather to use mutex to protect")
}
func main() {
	flag.Parse()
	var mu sync.Mutex
	var buffer bytes.Buffer

	const (
		maxGoroutine = 5
		maxBolck     = 10
		maxNumber    = 20
	)
	sign := make(chan struct{}, maxGoroutine)
	for i := 0; i < maxGoroutine; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j < maxBolck; j++ {
				header := fmt.Sprintf("\n[id: %d, iteration: %d]", id, j)
				data := fmt.Sprintf(" %d", id*j)
				if protect > 0 {
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("write header %d error: %s\n", id, err)
				}
				for k := 0; k < maxNumber; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("write data %d error: %s\n", id, err)
					}
				}
				if protect > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}
	for i := 0; i < maxGoroutine; i++ {
		<-sign
	}
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("ReadAll error: %s\n", err)
	}
	log.Printf("The Contents:\n%s", data)
}
