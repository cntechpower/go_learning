package main
import (
	"fmt"
	"strings"
)

func main(){
	var str string="Hello, how is it going,Hugo?"
	fmt.Printf("Number os H in %s is:%d\n",str,strings.Count(str,"H"))
}
