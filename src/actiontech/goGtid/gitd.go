package main

import (
	"fmt"

	gtid "github.com/ikarishinjieva/go-gtid"
)

func main() {
	//slaveGtid := "56ddb5ca-c30b-11e9-8c02-0242ac148601:1-201856,56ddb5ca-c30b-11e9-8c02-0242ac148602:1-10"
	//masterGtid := "56ddb5ca-c30b-11e9-8c02-0242ac148601:1-201860"
	previousGtids := "56ddb5ca-c30b-11e9-8c02-0242ac148601:1-201860,c95e95e8-c872-11e9-ac89-0242ac148602:1-33561"
	gtid5 := "c95e95e8-c872-11e9-ac89-0242ac148602:33560-33574"
	gtid6 := "c95e95e8-c872-11e9-ac89-0242ac148602:33562-33574"

	//inconsistentGtid, _ := gtid.GtidSub(slaveGtid, masterGtid)
	//fmt.Println(inconsistentGtid)
	res, _ := gtid.GtidSub(gtid5, previousGtids)
	fmt.Println(res)
}
