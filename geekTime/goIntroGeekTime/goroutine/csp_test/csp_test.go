package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(50 * time.Millisecond)
	return "service is done"
}

func otherTask() {
	fmt.Println("Working on otherTask")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask is done")
}

func asyncService() chan string {
	// retCh := make(chan string) //没有buffer, retCh<-ret会被阻塞到client接收
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("asyncService returned")
		retCh <- ret
		fmt.Println("asyncService exited")
	}()
	return retCh
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func TestServiceAsync(t *testing.T) {
	//fmt.Println(service())
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
}
