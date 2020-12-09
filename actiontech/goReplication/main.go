package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/siddontang/go-log/log"

	uuid "github.com/satori/go.uuid"

	"github.com/siddontang/go-mysql/replication"
)

type GtidInfo struct {
	gtid string
	sql  []string
}

func parseBinlog() {
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
				gtidInfos = append(gtidInfos, gtidInfo)
			}
			//clear current gtid and set gtid to next
			gtidInfo = &GtidInfo{
				gtid: uuidInEvent.String() + ":" + strconv.FormatInt(gtidEvent.GNO, 10),
				sql:  nil,
			}
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

	for _, gtid := range gtidInfos {
		fmt.Printf("GTID is %v, sql is %v\n", gtid.gtid, gtid.sql)
	}

}

func matchMysqlBinlogOutput() error {
	//Transaction is split by GTID event in mysqlbinlog output, below comment is a GTID event sample
	//#190827 16:26:19 server id 1174285442  end_log_pos 2784191 CRC32 0x14a911aa 	GTID	last_committed=6242	sequence_number=6243	rbr_only=yes	original_committed_timestamp=0	immediate_commit_timestamp=0	transaction_length=0
	startOfTransactions := regexp.MustCompile("#.*GTID[[:space:]]*last_committed.*transaction_length=[[:digit:]]+")

	//c95e95e8-c872-11e9-ac89-0242ac148602:39815
	gtidFormat := regexp.MustCompile("[[:graph:]]{36}:[[:digit:]]+")

	//   ### UPDATE `universe`.`u_delay`
	//   ### WHERE
	//   ###   @1='ustats'
	//   ###   @2='2019-08-27 16:26:17'
	//   ###   @3=39712
	//   ### SET
	//   ###   @1='ustats'
	//   ###   @2='2019-08-27 16:26:18'
	//   ###   @3=39713
	//   # at 2783649
	//   #190827 16:26:18 server id 1174285442  end_log_pos 2783680 CRC32 0x8e9d69a3     Xid = 1559957
	//   COMMIT/*!*/
	startOfSql := regexp.MustCompile("(?s)#{3}.*COMMIT")
	startOfEvents := regexp.MustCompile("# at [[:digit:]]+")
	space := regexp.MustCompile("\\s+")
	cmd := exec.Command("cat", "./mysqlbinlog.out")
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	trxs := startOfTransactions.Split(string(output), 100)
	//fmt.Println(trxs[1])
	for _, trx := range trxs {
		gtid := gtidFormat.FindStringSubmatch(trx)
		//match full trx data in sqlsInTrx[0]
		sqlsInTrx := startOfSql.FindStringSubmatch(trx)
		if sqlsInTrx != nil {
			sqls := startOfEvents.Split(sqlsInTrx[0], 100)
			//sqls[len(sqls)-1] is always be COMMIT, so drop it
			for _, sql := range sqls[:len(sqls)-1] {
				fmt.Printf("----------------------------------------------------------------------------------------")
				//fmt.Printf("\n%v, \n%v\n", gtid[0], strings.Replace(sql, "#", "", 100))
				sqlNew := space.ReplaceAllString(strings.Replace(sql, "#", "", 100), " ")
				fmt.Printf("\n%v, \n%v\n", gtid[0], sqlNew)
			}
		}
	}
	return nil

}

func readLineBinlog() error {
	cmd := exec.Command("cat", "./trxMultiSQL.out")
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	type GtidContent struct {
		gtid string
		sql  []string
	}
	gtids := make([]*GtidContent, 0)
	currentGtid := new(GtidContent)
	currentSql := ""
	binlogReader := bufio.NewReader(bytes.NewReader(output))
	for {
		line, _, err := binlogReader.ReadLine()
		if err != nil {
			break
		}
		typeOfLine, value := binlogLineProcessor(string(line))
		switch typeOfLine {
		case "GTID":
			if currentGtid.gtid != "" {
				gtids = append(gtids, currentGtid)
				currentGtid = new(GtidContent)
			}
			currentGtid.gtid = value
		case "SQL":
			currentSql += value
		case "COMMIT":
			//currentGtid.sql = append(currentGtid.sql, strings.Replace(currentSql, "  ", "", -1))
			//currentSql = ""
			if currentGtid.gtid != "" && currentSql != "" {
				currentGtid.sql = append(currentGtid.sql, currentSql)
				currentSql = ""
			}
		case "POS":
			if currentGtid.gtid != "" && currentSql != "" {
				currentGtid.sql = append(currentGtid.sql, currentSql)
				currentSql = ""
			}
		}

	}

	//append last currentGtid to gtids
	if currentGtid.gtid != "" {
		gtids = append(gtids, currentGtid)
	}
	for _, v := range gtids {
		fmt.Printf("GTID is %v\n", v.gtid)
		for _, s := range v.sql {
			fmt.Printf("SQL is %v\n", s)
		}
	}
	return nil

}

func binlogLineProcessor(line string) (typeOfLine, value string) {
	//match GTID event:
	//line eg: SET @@SESSION.GTID_NEXT= 'c95e95e8-c872-11e9-ac89-0242ac148602:39815'/*!*/;
	gtidFormat := regexp.MustCompile("[[:graph:]]{36}:[[:digit:]]+")
	if strings.HasPrefix(line, "SET @@SESSION.GTID_NEXT= ") {
		value = gtidFormat.FindString(line)
		return "GTID", value
	}
	if strings.HasPrefix(line, "###") {
		value = strings.TrimPrefix(line, "###")
		return "SQL", value
	}
	if strings.HasPrefix(line, "COMMIT") {
		return "COMMIT", ""
	}
	if strings.HasPrefix(line, "# at") {
		return "POS", ""
	}
	return "", ""
}

func main() {
	//parseBinlog()
	if err := readLineBinlog(); err != nil {
		log.Fatal(err.Error())
	}

}
