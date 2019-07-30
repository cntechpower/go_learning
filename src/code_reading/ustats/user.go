package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"strconv"
)

func main() {
	//https://golang.org/pkg/os/user/
	username, _ := user.Current()
	fmt.Println(username.Name)
	if rootUser, err := user.Lookup("root"); err == nil {
		// 将字符串转换为十进制整数，即：ParseInt(s, 10, 0) 的简写）
		//func Atoi(s string) (int, error)
		gids, _ := rootUser.GroupIds()
		for num, gidStr := range gids {
			gName, _ := user.LookupGroupId(gidStr)
			// https://tip.golang.org/pkg/os/user/#LookupGroup
			fmt.Printf("User root group %v is %v\n", num+1, gName.Name)
		}
		//fmt.Println(gids)
		uid, _ := strconv.Atoi(rootUser.Uid)
		//fmt.Printf("0 to Atoi is %v\n", uid)
		//uidParse, _ := strconv.ParseInt("0", 10, 0)
		//fmt.Printf("0 ParseInt is %v\n", uidParse)
		fmt.Printf("root username is %v\n", rootUser.Username)
		fmt.Printf("root uid is %v\n", uid)
		idCmd, _ := exec.Command("id").Output()
		fmt.Printf("output is %v", idCmd)
		/*
			if err := idCmd.Start(); err != nil {
				fmt.Printf("exec cmd error: %v", err)
			} else {
				if output, err := idCmd.Output(); err != nil {
					fmt.Printf("get output error: %v", err)
				} else {
					fmt.Printf("Exec output is %s\n", output)
				}
		*/

	}
	/*
		} else {
			fmt.Printf("Lookup User Root fail, reason is %v\n", err)
		}
	*/
	//uid := os.Getuid()
	//fmt.Println(uid)
	//hostname, _ := os.Hostname()
	//fmt.Println(hostname)
	//fmt.Printf(os.Args[0])

}
