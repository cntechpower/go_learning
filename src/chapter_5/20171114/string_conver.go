package main
import (
	"fmt"
	"strconv"
	"os"
)
func string2int(){
	var orig string="123"
	var newS string
	fmt.Printf("The Size Of ints is: %d\n",strconv.IntSize)
	an,err:=strconv.Atoi(orig)
	if err!= nil{
		fmt.Printf("orig %s is not an integer -- exiting with error\n",orig)
		return
	}
	fmt.Printf("The integer is %d\n",an)
	an=an+5
	newS=strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n",newS)
}

func main(){
	var test string="test"
	an,err:=strconv.Atoi(test)
	if err!=nil{
		fmt.Printf("value is %d,exiting with error %v",an,err)
		os.Exit(1)
	}
}
