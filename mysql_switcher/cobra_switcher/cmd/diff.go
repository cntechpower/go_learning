package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	//"os"
	//"flag"
)

var DiffCMD = &cobra.Command{
	Use:   "diff",
	Short: "diff GTID",
	Long:  "diff given two database's GTID and print whether they are same",
	Run: func(cmd *cobra.Command, args []string) {
		/*old_master_dsn := Get_dsn(db_user, db_pass, master_db_host, master_db_port)
		  old_slave_dsn := Get_dsn(db_user, db_pass, slave_db_host, slave_db_port)
		  if Ping_db(old_master_dsn) == false || Ping_db(old_slave_dsn) == false {
		          fmt.Printf("Database Chech Failed!\n")
		          return
		  }
		  old_master_gtid := Get_gtid(old_master_dsn)
		  old_slave_gtid := Get_gtid(old_slave_dsn)
		  if Compare_gtid(old_master_gtid, old_slave_gtid) == false {
		          fmt.Printf("Exiting!\n")
		          return
		  }*/
		/*if 0 == len(args) {
		        os.Exit(1)
		}*/
		fmt.Println("Print: " + strings.Join(args, " "))
		return
	},
}

func init() {
	RootCmd.AddCommand(DiffCMD)
}

