package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"github.com/siddontang/go-log/log"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

func main() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "10.186.62.3",
		Port:     33306,
		User:     "admin",
		Password: "admin",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	gtidSet, err := mysql.ParseMysqlGTIDSet("35d5c7fa-c7d9-11e9-91a8-cb2d7ff653bb:1-5")
	if err != nil {
		log.Fatal("parse mysql gtid set %v error: %v", err)
	}

	streamer, _ := syncer.StartSyncGTID(gtidSet)
	for {
		ev, _ := streamer.GetEvent(context.Background())
		ev.Dump(os.Stdout)
		switch ev.Event.(type) {
		case *replication.GTIDEvent:
			gtidEvent := ev.Event.(*replication.GTIDEvent)
			uuidInEvent, _ := uuid.FromBytes(gtidEvent.SID)
			gtid := uuidInEvent.String() + ":" + strconv.FormatInt(gtidEvent.GNO, 10)
			fmt.Printf("Got GTID :%v\n", gtid)
		default:
			fmt.Printf("Other Event\n")

		}
	}

}
