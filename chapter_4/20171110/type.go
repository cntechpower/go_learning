package main
import "fmt"

type d int
type s string
func main(){
	var a,b d=3,4
	var name s="dujinyang"
	c:=a+b
	fmt.Printf("c has the values: %d\n",c)
	fmt.Printf("My name is : %s\n",name)
	fmt.Printf(`test \n`)
}
