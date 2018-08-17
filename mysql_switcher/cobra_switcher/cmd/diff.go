package cmd

import (
	"../common"
	"fmt"
	"github.com/spf13/cobra"
)

var DiffCMD = &cobra.Command{
	Use:   "diff",
	Short: "diff GTID",
	Long:  "diff given two database's GTID and print whether they are same",
	Run: func(cmd *cobra.Command, args []string) {
		db_pass, db_user, master_db_host, master_db_port, slave_db_host, slave_db_port = get_flags(cmd)
		print_flags(db_pass, db_user, master_db_host, master_db_port, slave_db_host, slave_db_port)
		old_master_dsn := common.Get_dsn(db_user, db_pass, master_db_host, master_db_port)
		old_slave_dsn := common.Get_dsn(db_user, db_pass, slave_db_host, slave_db_port)
		if common.Ping_db(old_master_dsn) == false || common.Ping_db(old_slave_dsn) == false {
			fmt.Printf("Database Chech Failed!\n")
			return
		}
		old_master_gtid := common.Get_gtid(old_master_dsn)
		old_slave_gtid := common.Get_gtid(old_slave_dsn)
		if common.Compare_gtid(old_master_gtid, old_slave_gtid) == false {
			fmt.Printf("Data Not Consist!\n")
			return
		}
		return
	},
}

func init() {
	RootCmd.AddCommand(DiffCMD)
}
