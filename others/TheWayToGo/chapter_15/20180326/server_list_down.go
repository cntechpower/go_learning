package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	mapUsers = make(map[string]int)
	fmt.Println("Starting Server...")
	listener, err := net.Listen("tcp", "localhost:50000")
	checkerror(err)
	for {
		conn, err := listener.Accept()
		checkerror(err)
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		checkerror(err)
		input := string(buf)
		if strings.Contains(input, "list") {
			listUsers()
		}
		if strings.Contains(input, "shutdown") {
			fmt.Println("Server shutting down...")
			os.Exit(0)
		}
		name_index := strings.Index(input, "says")
		clientName := input[0 : name_index-1]
		mapUsers[string(clientName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}
func checkerror(error error) {
	if error != nil {
		panic("Error: " + error.Error())
	}
}
func listUsers() {
	fmt.Println("--------------------------------")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------")
}
