package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Username: ")
	clientName, _ := inputReader.ReadString('\n')
	//string.Trim:cut string by given field char
	trimmedClient := strings.Trim(clientName, "\n")
	for {
		fmt.Printf("send> ")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "exit" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		if err != nil {
			fmt.Println("Error sending", err.Error())
			return
		}
	}
}
