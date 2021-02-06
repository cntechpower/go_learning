package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"k8s.io/apimachinery/pkg/util/rand"
)

const sql4InsertCdKey = `
INSERT INTO rewards_cdkey (cdkey_name, cdkey_content, unique_id)
VALUES %s;`

const sql4SendCdkey = `
UPDATE
  rewards_cdkey FORCE INDEX(ix_cdkey_mid)
SET
  mid = 1930204063,
  activity_id = 100,
  unique_id = ?,
  is_used = 1
WHERE
  cdkey_name = ?
  AND mid = 0
  AND activity_id = 0
  AND is_used = 0
LIMIT
  1;`


//go test -count 1 -v .
func TestCdkey(t *testing.T) {
	dsn := "admin:admin@(10.0.0.2:3306)/test?autocommit=1&multiStatements=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("truncate table rewards_cdkey")
	params := strings.Builder{}
	values := make([]interface{}, 0)
	first := true
	u := 1
	for i := 0; i < 500000; i++ {
		uniqIDStr := strconv.Itoa(-u)
		if first {
			params.WriteString("(?,?,?)")
			values = append(values, "拜年纪红包封面", rand.String(15), uniqIDStr)
			first = false
		} else {
			params.WriteString(",(?,?,?)")
			values = append(values, "拜年纪红包封面", rand.String(15), uniqIDStr)
		}
		if i%1000 == 0 {
			_, err = db.Exec(fmt.Sprintf(sql4InsertCdKey, params.String()), values...)
			if err != nil {
				panic(err)
			}
			params = strings.Builder{}
			values = make([]interface{}, 0)
			first = true
		}
		u += 1
	}

	for i := 0; i < 500000; i++ {
		uniqIDStr := strconv.Itoa(-u)
		if first {
			params.WriteString("(?,?,?)")
			values = append(values, "干杯红包封面", rand.String(15), uniqIDStr)
			first = false
		} else {
			params.WriteString(",(?,?,?)")
			values = append(values, "干杯红包封面", rand.String(15), uniqIDStr)
		}
		if i%1000 == 0 {
			_, err = db.Exec(fmt.Sprintf(sql4InsertCdKey, params.String()), values...)
			if err != nil {
				panic(err)
			}
			params = strings.Builder{}
			values = make([]interface{}, 0)
			first = true
		}
		u += 1
	}

	for i := 0; i < 200000; i++ {
		ts := time.Now()
		_, err = db.Exec(sql4SendCdkey, rand.String(10), "拜年纪红包封面")
		if err != nil {
			panic(err)
		}
		spend := time.Since(ts).Milliseconds()
		if spend >= 20 {
			t.Logf("ms spend: %v", spend)
		}
	}

	for i := 0; i < 200000; i++ {
		ts := time.Now()
		_, err = db.Exec(sql4SendCdkey, rand.String(10), "干杯红包封面")
		if err != nil {
			panic(err)
		}
		spend := time.Since(ts).Milliseconds()
		if spend >= 20 {
			t.Logf("ms spend: %v", spend)
		}
	}

}
