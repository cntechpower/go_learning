package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "/v3/mysql/instance/get_inconsistent_gtid_sqls_from_slave?mysql_instance_id=mysql-fhizem&is_dry_run=false&parse_gtid_timeout=15"

	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.Path)
	}
}
