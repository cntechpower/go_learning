package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/hpcloud/tail"
)

func startWatch(filePath, subString string) (error, chan struct{}) {
	t, err := tail.TailFile(filePath, tail.Config{
		MustExist: true,
		Follow:    true,
	})
	if err != nil {
		return err, nil
	}
	closeChan := make(chan struct{})
	go func() {
		for {
			select {
			case line := <-t.Lines:
				if strings.Contains(line.Text, subString) {
					fmt.Println(line.Text)
				}
			case <-closeChan:
				t.Cleanup()
				log.Println("stopped")
				return

			}
		}
	}()
	return nil, closeChan
}

func main() {
	err, closeChan := startWatch("/tmp/testfile", "blabla")
	if err != nil {
		log.Fatal(err)
	}
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGUSR2 /*graceful-shutdown*/)
	select {
	case <-quitChan:
		log.Printf("Server Existing")
		close(closeChan)
		time.Sleep(1 * time.Second)
	}
}
