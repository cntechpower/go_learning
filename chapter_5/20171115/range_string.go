package main
import "fmt"
func main(){
	str:="Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n",len(str))
	for pos,char:=range str{
		fmt.Printf("Character on position %d is: %c \n",pos,char)
	}
	fmt.Println()
	str2:="萎靡不振"
	fmt.Printf("The length of str2 is: %d\n",len(str))
	for pos,char:=range str2{
	fmt.Printf("character %c starts at byte position %d\n",char,pos)
	}
	fmt.Println()
}
