package main

import (
	"compress/gzip"
	"database/sql"
	"fmt"
	"io"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

func saveFileToDb1(name, fileContent string) error {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}
	defer db.Close()

	//ensure table exists
	_, err = db.Query("create table if not exists test.t1 (name varchar(100),content longblob)")
	if err != nil {
		return err
	}

	_, err = db.Query("insert into test.t1 values(?,?)", name, fileContent)
	if err != nil {
		return err
	}
	return nil
}

func loadFileFromDb1(name, filePath string) error {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}
	defer db.Close()
	res, err := db.Query("select content from test.t1 where name = ?", name)
	if err != nil {
		return err
	}
	defer res.Close()
	file, err := os.Create(fmt.Sprintf("%v/%v", filePath, name))
	if err != nil {
		return err
	}
	defer file.Close()
	var out string
	for res.Next() {
		if err := res.Scan(&out); err != nil {
			return err
		}
	}

	fmt.Println(out)
	return nil

}
func gzipFile(srcFileName, dstTarBallPathName string) error {
	src, err := os.Open(srcFileName)
	if err != nil {
		return err
	}
	defer src.Close()
	dstFile, err := os.Create(dstTarBallPathName)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	srcReader, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer srcReader.Close()
	dstWriter := gzip.NewWriter(dstFile)
	defer dstWriter.Close()
	if _, err := io.Copy(dstWriter, srcReader); err != nil {
		return err
	}
	return nil
}

func main() {
	//file, err := ioutil.ReadFile("/tmp/adobegc.log")
	//if err != nil {
	//	panic(err)
	//}
	//if err := saveFileToDb1("1", string(file)); err != nil {
	//	panic(err)
	//}
	//if err := loadFileFromDb1("1", "/tmp"); err != nil {
	//	panic(err)
	//}
	if err := gzipFile("/tmp/adobegc.log", "/tmp/adobe.tar.gz"); err != nil {
		panic(err)
	}

}
