package main
import (
        "fmt"
	"github.com/spf13/cobra"
	 "./gtid"
)
var db_user string
var db_pass string
var master_db_host string
var master_db_port int
var slave_db_host string
var slave_db_port int

func main(){
	var cmdDiff=&cobra.Command{
		Use: "diff --flags",
		Short: "diff GTID",
		Long: "diff given two database's GTID and print whether they are same",
		Run: func(cmd *cobra.Command, args []string){
			old_master_dsn:=gtid.get_dsn(db_user,db_pass,master_db_host,master_db_port)
			old_slave_dsn:=gtid.get_dsn(db_user,db_pass,master_db_host,master_db_port)
			old_master_status:=gtid.ping_db(old_master_dsn)
			old_slave_status:=gtid.ping_db(old_slave_dsn)
			old_master_gtid:=gtid.get_gtid(old_master_dsn)
			old_slave_gtid:=gtid.get_gtid(old_slave_dsn)
			if gtid.compare_gtid(old_master_gtid,old_slave_gtid)==false{
				fmt.Printf("Exiting!\n")
				return
			}
			return
		},
	}
	cmdDiff.Flags().StringVarP(&db_user, "db_user", "u", "admin", "database user")
	cmdDiff.Flags().StringVarP(&db_pass, "db_pass", "p", "1", "database password")
	cmdDiff.Flags().StringVarP(&master_db_host, "master_db_host", "h1", "127.0.0.1", "master database host")
	cmdDiff.Flags().IntVarP(&master_db_port, "master_db_port", "P1", 3306, "master database port")
	cmdDiff.Flags().StringVarP(&slave_db_host, "slave_db_host", "h2", "127.0.0.1", "slave database host")
	cmdDiff.Flags().IntVarP(&slave_db_port, "slave_db_port", "P2", 3306, "slave database port")
	var rootCmd = &cobra.Command{Use: "mswitcher"}
	rootCmd.AddCommand(cmdDiff)
	rootCmd.Execute()
}
