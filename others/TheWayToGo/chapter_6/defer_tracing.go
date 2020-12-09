package main
import "fmt"
func trace(s string){fmt.Println("entring:",s)}
func untrace(s string){fmt.Println("leaving:",s)}

func a(){
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}
func main(){
	a()
}
