package main
import (
	"fmt"
	"strings"
)

func main(){
	var origS string="Hi where!"
	var newS string
	newS =strings.Repeat(origS,3)
	fmt.Printf("The new is %s \n",newS)
}
