package main
import "fmt"
func switch1(){
	var num1 int=99
	switch num1{
	case 98:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}

func main(){
	var num1 int=20
	switch{
	case num1<0:
		fmt.Println("Negaitve")
	case num1>0&&num1 <10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}
