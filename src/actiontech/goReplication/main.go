package main

import (
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"github.com/siddontang/go-mysql/replication"
)

type GtidInfo struct {
	gtid string
	sql  []string
}

func main() {
	gtidInfos := make([]*GtidInfo, 0)
	gtidInfo := new(GtidInfo)
	binlogParser := replication.NewBinlogParser()
	err := binlogParser.ParseFile("./mysql-bin.000003", 0, func(ev *replication.BinlogEvent) error {
		switch ev.Event.(type) {
		case *replication.GTIDEvent:
			gtidEvent := ev.Event.(*replication.GTIDEvent)
			uuidInEvent, _ := uuid.FromBytes(gtidEvent.SID)

			//append current gtid to list
			if gtidInfo.gtid != "" {
				//fmt.Printf("append gtid: %v, sql: %v to list.\n", gtidInfo.gtid, gtidInfo.sql)
				gtidInfos = append(gtidInfos, gtidInfo)
			}
			//clear current gtid and set gtid to next
			gtidInfo = &GtidInfo{
				gtid: uuidInEvent.String() + ":" + strconv.FormatInt(gtidEvent.GNO, 10),
				sql:  nil,
			}
		//case *replication.RowsEvent:
		//	rowsEvent := ev.Event.(*replication.RowsEvent)
		//	table := string(rowsEvent.Table.Table)
		//	fmt.Println(table)
		//	rowsEvent.Dump(os.Stdout)
		case *replication.RowsQueryEvent:
			rowsQueryEvent := ev.Event.(*replication.RowsQueryEvent)
			if gtidInfo.gtid == "" {
				return fmt.Errorf("found RowQueryEvent without gtid")
			}
			gtidInfo.sql = append(gtidInfo.sql, string(rowsQueryEvent.Query))

		default:
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("parse binlog error %v", err)
	}
	//json, _ := json.Marshal(gtidInfos)
	//fmt.Println(string(json))

	for _, gtid := range gtidInfos {
		fmt.Printf("GTID is %v, sql is %v\n", gtid.gtid, gtid.sql)
	}

}
