package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func LookupUidGidByUser(username string) (uid int, gids []int, err error) {
	if "" != username {
		u, err := user.Lookup(username)
		if nil != err {
			return 1, []int{}, err
		}
		uid, _ = strconv.Atoi(u.Uid)
		gids = []int{}
		{
			groupIds, err := u.GroupIds()
			if nil != err {
				return 2, []int{}, err
			}
			for _, gidStr := range groupIds {
				gid, _ := strconv.Atoi(gidStr)
				gids = append(gids, gid)
			}
		}
	} else {
		uid = os.Getuid()
		gids = append(gids, os.Getgid())
		if groups, err := os.Getgroups(); nil == err {
			gids = append(gids, groups...)
		}
	}
	return uid, gids, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Username: ")
	username, _ := reader.ReadString('\n')
	uid, gids, err := LookupUidGidByUser(username)
	if err != nil {
		fmt.Println(err)
		os.Exit(uid)
	}
	fmt.Printf("%s Uid is %d ,Gid is %v\n", username, uid, gids)
}
