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
func main(){
	user:=flag.String("db_user","admin","Database User Name")
	passwd:=flag.String("db_passwd","1","Database User Password")
	master_host:=flag.String("master_db_host","127.0.0.1","Now Master Database Host")
	master_port:=flag.String("master_db_port","3306","Now Master Database Port")
	slave_host:=flag.String("salve_db_host","127.0.0.1","Now Slave Database Host")
	slave_port:=flag.String("slave_db_port","3307","Now Slave Database Port")
	flag.Parse()
	old_master_dsn:=get_dsn(*user,*passwd,*master_host,*master_port)
	old_slave_dsn:=get_dsn(*user,*passwd,*slave_host,*slave_port)
	old_master_status:=ping_db(old_master_dsn)
	old_slave_status:=ping_db(old_slave_dsn)
	if old_master_status==false||old_slave_status==false{
		fmt.Printf("Database Alive Check Failed\n")
	}
}
