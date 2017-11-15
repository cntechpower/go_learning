package main
import "fmt"

func for1(){
	for i:=0;i<5;i++{
		fmt.Printf("This is the %d iteration\n",i)
	}
}

func main(){
	str:="Go is a beautiful language"
	fmt.Printf("The length of str is: %d\n",len(str))
	for ix:=0;ix<len(str);ix++{
		fmt.Printf("Character on position %d is: %c\n",ix,str[ix])
	}
	str1:="萎靡不振"
	for ix:=0;ix<len(str1);ix++ {
		fmt.Printf("Character on position %d is: %c\n",ix,str1[ix])
	}
}
