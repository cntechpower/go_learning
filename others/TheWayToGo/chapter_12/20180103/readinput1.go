package main
import "fmt"
var(
	firstname,lastname,s string
	i int
	j float32
	input="56.12/5212/Go"
	format="%f/%d/%s"
)
func main(){
	fmt.Println("Please input your name: ")
	fmt.Scanln(&firstname,&lastname)
	fmt.Printf("Hi,%s %s!\n",firstname,lastname)
	fmt.Sscanf(input,format,&j,&i,&s)
	fmt.Println("From the string we read: ",i,j,s)
}
