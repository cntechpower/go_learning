package main

import (
	"net/http"
	"time"
)
import _ "net/http/pprof"

type t struct {
	s string
}

var strChan chan string

func init() {
	strChan = make(chan string, 1)
}

func testProducer() {
	for range time.NewTicker(time.Microsecond).C {
		strChan <- time.Now().Format(time.RFC3339)
	}
}

func flush(list []*t) []*t {
	return list[0:0:0]
}

func main() {
	list := make([]*t, 0)
	go http.ListenAndServe("0.0.0.0:6060", nil)
	go testProducer()
	commitTicker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-commitTicker.C:
			list = flush(list)
		case s, ok := <-strChan:
			if !ok {
				panic("")
			}
			list = append(list, &t{s: s})
		}
	}
}
