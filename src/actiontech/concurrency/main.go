package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkSchemaExist(ip, port, user, pass, schema string) (bool, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/?autocommit=1&multiStatements=true&loc=Local", user, pass, ip, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return false, err
	}
	defer db.Close()
	res, err := db.Query("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", schema)
	if err != nil {
		return false, err
	}
	defer res.Close()
	for res.Next() {
		return true, nil
	}
	return false, nil
}

func main() {
	//var wg sync.WaitGroup
	//maxRunningGoroutine := make(chan struct{}, 5) //hard code default concurrency limit: 5
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		maxRunningGoroutine <- struct{}{}
	//		defer func() { <-maxRunningGoroutine }()
	//		fmt.Println(i)
	//		time.Sleep(2 * time.Second)
	//	}(i)
	//}
	//wg.Wait()
	exist, err := checkSchemaExist("127.0.0.1", "3306", "root", "root", "test")
	fmt.Printf("isExist: %v, error: %v", exist, err)
}
