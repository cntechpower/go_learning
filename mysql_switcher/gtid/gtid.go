package gtid
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)
//func get_gtid()
//func compare_gtid()
//func read_only_on()
//func read_only_off()
//func stop_slave()
//func change_master()
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

