package main
import "fmt"

func main(){
	var arr1[5] int
	for i:=0;i<len(arr1);i++{
		arr1[i]=i*2
	}
	for i:=0;i<len(arr1);i++{
		value:=arr1[i]
		fmt.Printf("The No %d value is:%d\n",i,value)
	}
}
