package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var db_user string
var db_pass string
var master_db_host string
var master_db_port string
var slave_db_host string
var slave_db_port string

var RootCmd = &cobra.Command{Use: "mswitcher"}

func init() {
	RootCmd.PersistentFlags().StringVarP(&db_user, "db_user", "u", "admin", "database user")
	RootCmd.PersistentFlags().StringVarP(&db_pass, "db_pass", "p", "1", "database password")
	RootCmd.PersistentFlags().StringVarP(&master_db_host, "master_db_host", "a", "127.0.0.1", "master database host")
	RootCmd.PersistentFlags().StringVarP(&master_db_port, "master_db_port", "b", "3306", "master database port")
	RootCmd.PersistentFlags().StringVarP(&slave_db_host, "slave_db_host", "c", "127.0.0.1", "slave database host")
	RootCmd.PersistentFlags().StringVarP(&slave_db_port, "slave_db_port", "d", "3307", "slave database port")
}
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
