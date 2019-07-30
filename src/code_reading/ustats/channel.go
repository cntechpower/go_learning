/*
package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)

	go func() {
		for {
			if d, ok := <-data; ok {
				fmt.Println(d)
			} else {
				fmt.Println("Channel Closed")
			}
			time.Sleep(100 * time.Millisecond)
		}

	}()

	data <- 1 //发送要放在接收协程跑起来后面，因为发送后会阻塞等待接收
	data <- 2
	data <- 3
	close(data)
	//time.Sleep(time.Second)
}

*/

package main

import "fmt"
import "time"
import "os"

const (
	MAX_REQUEST_NUM = 10
	CMD_USER_POS    = 1
)

var (
	save chan bool
	quit chan bool
	req  chan *Request
)

type Request struct {
	CmdID int16
	Data  interface{}
}

type UserPos struct {
	X int16
	Y int16
}

func main() {
	newReq := Request{
		CmdID: CMD_USER_POS,
		Data: UserPos{
			X: 10,
			Y: 20,
		},
	}
	go handler()

	req <- &newReq

	time.Sleep(2000 * time.Millisecond)

	save <- true
	close(req)

	<-quit

}

func handler() {
	for {
		select {
		case <-save:
			saveGame()
		case r, ok := <-req:
			if ok {
				onReq(r)
			} else {
				fmt.Println("req chan closed.")
				os.Exit(0)
			}
		}
	}

}
func init() {
	req = make(chan *Request, MAX_REQUEST_NUM)
	save = make(chan bool)
	quit = make(chan bool)
}

func saveGame() {
	fmt.Printf("Do Something With Save Game.\n")
	quit <- true
}

func onReq(r *Request) {
	pos := r.Data.(UserPos)
	fmt.Println(r.CmdID, pos)
}
