package main
import "fmt"

func main(){
	slice1:=make([]int,0,10)
	for i:=0;i<cap(slice1);i++{
		slice1=slice1[0:i+1]
		slice1[i]=i
		fmt.Printf("The length of slice1 is: %d\n",len(slice1))
	}
	for i:=0;i<len(slice1);i++{
		fmt.Printf("The NO-%d values is:%d\n",i,slice1[i])
	}
}