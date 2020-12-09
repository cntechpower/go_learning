package main
import "fmt"
func Multiply(a,b int,reply *int){
	*reply=a*b
}
func main(){
	n:=0
	reply:=&n
	fmt.Printf("reply is now: %d\n",*reply)
	Multiply(10,5,reply)
	fmt.Printf("reply is now: %d\n",*reply)
}
