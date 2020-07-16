package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	binlogDir := "/opt/mysql/log/binlog/3306/mysql-bin"
	fmt.Println(filepath.Dir(binlogDir))
	fmt.Println(filepath.Join("/var/log", "test.pid"))
}
