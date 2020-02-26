package main
import(
	"fmt"
	"io"
	"os"
)
func CopyFile(src,dst string)(written int64,err error){
	srcFile,err:=os.Open(src)
	if err!=nil{
		fmt.Printf("Error Occured\n")
	}
	defer srcFile.Close()
	dstFile,err:=os.OpenFile(dst,os.O_WRONLY|os.O_CREATE, 0644)
	if err!=nil{
		fmt.Printf("lifu disappered!!!!")
	}
	return io.Copy(dstFile,srcFile)
}
func main(){
	CopyFile("input.txt","input.txt.bak")
	fmt.Println("Copy Done!")
}
