package main

func main() {}

//package main
//
//import (
//	"archive/tar"
//	"compress/gzip"
//	"database/sql"
//	"fmt"
//	"io"
//	"os"
//	"time"
//)
//import _ "github.com/go-sql-driver/mysql"
//
//func saveFileToDb1(name, fileContent string) error {
//	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	//ensure table exists
//	_, err = db.Query("create table if not exists test.t1 (name varchar(100),content longblob)")
//	if err != nil {
//		return err
//	}
//
//	_, err = db.Query("insert into test.t1 values(?,?)", name, fileContent)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func loadFileFromDb1(name, filePath string) error {
//	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//	res, err := db.Query("select content from test.t1 where name = ?", name)
//	if err != nil {
//		return err
//	}
//	defer res.Close()
//	file, err := os.Create(fmt.Sprintf("%v/%v", filePath, name))
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	var out string
//	for res.Next() {
//		if err := res.Scan(&out); err != nil {
//			return err
//		}
//	}
//
//	fmt.Println(out)
//	return nil
//
//}
//func gzipFile(srcFileName, dstTarBallPathName string) error {
//	src, err := os.Open(srcFileName)
//	if err != nil {
//		return err
//	}
//	defer src.Close()
//	dstFile, err := os.Create(dstTarBallPathName)
//	if err != nil {
//		return err
//	}
//	defer dstFile.Close()
//	srcReader, err := gzip.NewReader(src)
//	if err != nil {
//		return err
//	}
//	defer srcReader.Close()
//	dstWriter := gzip.NewWriter(dstFile)
//	dstWriter.Name = srcFileName
//	dstWriter.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)
//	defer dstWriter.Close()
//	if _, err := io.Copy(dstWriter, srcReader); err != nil {
//		return err
//	}
//	return nil
//}
//
//func gzipFile2(srcFileName, dstTarBallPathName string) error {
//	src, err := os.Open(srcFileName)
//	if err != nil {
//		return err
//	}
//	defer src.Close()
//	srcInfo, err := src.Stat()
//	if err != nil {
//		return err
//	}
//	fmt.Println(srcInfo.Name())
//	headerInfo, err := tar.FileInfoHeader(srcInfo, "")
//	if err != nil {
//		return err
//	}
//	headerInfo.Name = srcFileName
//	dstFile, err := os.Create(dstTarBallPathName)
//	if err != nil {
//		return err
//	}
//	defer dstFile.Close()
//	dstGzipWriter := gzip.NewWriter(dstFile)
//	defer dstGzipWriter.Close()
//	//dstTarWriter := tar.NewWriter(dstGzipWriter)
//	//dstTarWriter.Close()
//	//if err := dstTarWriter.WriteHeader(headerInfo); err != nil {
//	//	return err
//	//}
//	dstGzipWriter.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)
//	defer dstGzipWriter.Close()
//	if _, err := io.Copy(dstGzipWriter, src); err != nil {
//		return err
//	}
//	return nil
//}
//
//func gUnzipFile(TarBallPathName string) error {
//	tarBall, err := os.Open(TarBallPathName)
//	if err != nil {
//		return err
//	}
//	defer tarBall.Close()
//	gzipReader := gzip.NewReader(tarBall)
//}
//
//func main() {
//	//file, err := ioutil.ReadFile("/tmp/adobegc.log")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//if err := saveFileToDb1("1", string(file)); err != nil {
//	//	panic(err)
//	//}
//	//if err := loadFileFromDb1("1", "/tmp"); err != nil {
//	//	panic(err)
//	//}
//	if err := gzipFile2("/tmp/adobegc.log", "/tmp/adobe.tar.gz"); err != nil {
//		panic(err)
//	}
//
//}
