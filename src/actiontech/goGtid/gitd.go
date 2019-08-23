package main

import (
	"fmt"

	gtid "github.com/ikarishinjieva/go-gtid"
)

func main() {
	masterGtid := "56ddb5ca-c30b-11e9-8c02-0242ac148601:1-201856"
	slaveGtid := "56ddb5ca-c30b-11e9-8c02-0242ac148601:1-201860"
	if ok, _ := gtid.GtidEqual(masterGtid, slaveGtid); ok {
		fmt.Println("GTID Equal")
	} else {
		fmt.Println("GTID Not Equal")
	}
	sub1, _ := gtid.GtidSub(masterGtid, slaveGtid)
	fmt.Println(sub1)
	sub2, _ := gtid.GtidSub(slaveGtid, masterGtid)
	fmt.Println(sub2)
}
