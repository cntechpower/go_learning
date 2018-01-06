package main
import (
	"fmt"
	"bufio"
	"io"
	"os"
)
func main(){
	inputFile,inputError:=os.Open("input.txt")
	if inputError!=nil{
		fmt.Printf("Error Occurred\n")
		return
	}
	defer inputFile.Close()
	inputReader:=bufio.NewReader(inputFile)
	for{
		inputString,readerError:=inputReader.ReadString('\n')
		if readerError==io.EOF{
			return
		}
		fmt.Printf("The input was: %s\n",inputString)
	}
}
