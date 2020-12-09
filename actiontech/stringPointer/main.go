package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
)

func main() {
	path := "./tmp"
	absPath, err := filepath.Abs(path)
	fmt.Printf("path: %v , error: %v\n", absPath, err)
	ip, ipNet, err := net.ParseCIDR("127.16.0.0/16")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip.String())
	fmt.Println(ipNet.IP.String())
}
