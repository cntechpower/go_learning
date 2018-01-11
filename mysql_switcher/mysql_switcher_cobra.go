package main
import (
        "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/spf13/cobra"
)
var db_user string
var db_pass string
var master_db_host string
var master_db_port string
var slave_db_host string
var slave_db_port string

func main(){
	var cmdDiff=&cobra.Command{
		Use: "diff --flags",
		Short: "diff GTID",
		Long: "diff given two database's GTID and print whether they are same",
		Run: func(cmd *cobra.Command, args []string){
			old_master_dsn:=get_dsn(db_user,db_pass,master_db_host,master_db_port)
			old_slave_dsn:=get_dsn(db_user,db_pass,slave_db_host,slave_db_port)
			//old_master_status:=ping_db(old_master_dsn)
			//old_slave_status:=ping_db(old_slave_dsn)
			old_master_gtid:=get_gtid(old_master_dsn)
			old_slave_gtid:=get_gtid(old_slave_dsn)
			if compare_gtid(old_master_gtid,old_slave_gtid)==false{
				fmt.Printf("Exiting!\n")
				return
			}
			return
		},
	}
	cmdDiff.Flags().StringVarP(&db_user, "db_user", "u", "admin", "database user")
	cmdDiff.Flags().StringVarP(&db_pass, "db_pass", "p", "1", "database password")
	cmdDiff.Flags().StringVarP(&master_db_host, "master_db_host", "a", "127.0.0.1", "master database host")
	cmdDiff.Flags().StringVarP(&master_db_port, "master_db_port", "b", "3306", "master database port")
	cmdDiff.Flags().StringVarP(&slave_db_host, "slave_db_host", "c", "127.0.0.1", "slave database host")
	cmdDiff.Flags().StringVarP(&slave_db_port, "slave_db_port", "d", "3306", "slave database port")
	var rootCmd = &cobra.Command{Use: "mswitcher"}
	rootCmd.AddCommand(cmdDiff)
	rootCmd.Execute()
}
func get_dsn(user,passwd,host,port string)(string){
	return(fmt.Sprintf("%s:%s@(%s:%s)/", user,passwd,host,port))
	}
func ping_db(dsn string)(alive bool){
	alive=true
	db,err:=sql.Open("mysql",dsn)
	if err!=nil{
		alive=false
		fmt.Printf("Database %s is dead or Access Denied!\n",dsn)
		return alive
	}
	defer db.Close()
	err=db.Ping()
	if err!=nil{
		alive=false
		fmt.Printf("Database %s is dead or Access Denied!\n",dsn)
		return alive
	}
	fmt.Printf("Database %s is alive\n",dsn)
	return alive
}
func get_gtid(dsn string)(string){
	db,_:=sql.Open("mysql",dsn)
	defer db.Close()
	result,_:=db.Query("show master status")
	var s1,s2,s3,s4,Executed_Gtid_Set string
	result.Next()
	result.Scan(&s1,&s2,&s3,&s4,&Executed_Gtid_Set)
	fmt.Printf("Database %s has GTID: %s\n",dsn,Executed_Gtid_Set)
	return fmt.Sprintf("%s",Executed_Gtid_Set)
}
func compare_gtid(master_gtid,slave_gtid string)(bool){
	if strings.EqualFold(master_gtid,slave_gtid){
		fmt.Printf("Gtid Consistent\n")
		return true
	}else{
		fmt.Printf("Gtid Not Consistent\n")
		return false
	}
	return false
}

