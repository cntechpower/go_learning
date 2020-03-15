package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("20060102_150405"))
	fmt.Println(filepath.Base("snapshot/history/mysql-taxb4w[task-1]_20200312_131844.log"))
}
