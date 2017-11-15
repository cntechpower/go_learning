package main
import(
	"fmt"
)
func Abs(x int) int{
	if x<0{
		return -x
	}
	return x
}
func printabx(){
	x1:=-5
	x2:=Abs(x1)
	fmt.Println(x2)
}


