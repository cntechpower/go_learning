package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"flag"
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
		return alive
	}
	defer db.Close()
	err=db.Ping()
	if err!=nil{
		alive=false
		return alive
	}
	return alive
}
func main(){
	user:=flag.String("db_user","admin","Database User Name")
	passwd:=flag.String("db_passwd","1","Database User Password")
	host:=flag.String("db_host","127.0.0.1","Database Host")
	port:=flag.String("db_port","3306","Database Port")
	flag.Parse()
	dsn:=get_dsn(*user,*passwd,*host,*port)
	status:=ping_db(dsn)
	if status==true{
		fmt.Printf("Database %s:%s is alive\n",*host,*port)
	}else{
		fmt.Printf("Database %s:%s is dead or access denied\n",*host,*port)
	}
}
