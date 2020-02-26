package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/hpcloud/tail"

	"github.com/radovskyb/watcher"

	"github.com/fsnotify/fsnotify"
)

func startWatcher() {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only notify rename and move events.
	w.FilterOps(watcher.Rename, watcher.Move)

	// Only files that match the regular expression during file listings
	// will be watched.
	r := regexp.MustCompile("^abc$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add("/tmp/testfile"); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	// Trigger 2 events after watcher started.
	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Create, nil)
		w.TriggerEvent(watcher.Remove, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func startWatch(filePath string) (error, chan error, chan struct{}) {
	watcher, err := fsnotify.NewWatcher()
	closeChan := make(chan struct{})
	errChan := make(chan error)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					errChan <- fmt.Errorf("get watcher events error")
					close(errChan)
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("modified file: %v,content: %v\n", event.Name, event.String())
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					errChan <- fmt.Errorf("error occured but get message error")
					close(errChan)
				}
				log.Println("error:", err)
			case <-closeChan:
				watcher.Close()
				log.Println("stopped")
				return
			}
		}
	}()
	err = watcher.Add(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return nil, errChan, closeChan
}

func startTail(filePath string) (error, chan struct{}) {
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
				fmt.Println(line.Text)
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
	//err, errChan, closeChan := startWatch("/tmp/testfile")
	//go startWatcher()
	//if err != nil {
	//	log.Fatal(err)
	//}
	err, closeChan := startTail("/tmp/testfile")
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
		//case err := <-errChan:
		//	log.Printf("watcher error %v ,will close it", err)
		//	close(closeChan)
	}
}
