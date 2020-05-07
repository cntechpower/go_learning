package main
import(
	"fmt"
	"os"
	"bufio"
)
func main(){
	inputreader:=bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	input,err:=inputreader.ReadString('\n')
	if err!=nil{
		fmt.Printf("There is something wrong\n")
		return
	}
	fmt.Printf("Your name is : %s",input)
	switch input{
	case "Dujinyang\n": fmt.Println("Hello,my lord")
	case "dujinyang\n": fmt.Println("Hello,my lord")
	case "DuJinYang\n": fmt.Println("Hello,my lord")
	default: fmt.Println("Get Lost,You are not my lord")
	}
}
