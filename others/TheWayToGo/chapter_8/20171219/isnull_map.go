package main
import "fmt"
func main(){
	var value int
	var isNull bool
	map1:=make(map[string]int)
	map1["Beijing"]=1
	map1["Shanghai"]=2
	value,isNull=map1["Beijing"]
	if isNull{
		fmt.Printf("The values of Beijing is %d\n",value)
	}else{
		fmt.Printf("The values of Beijing not exists")
	}

	value,isNull=map1["Shanghai"]
	if isNull{
		fmt.Printf("The values of Shanghai is %d\n",value)
	}else{
		fmt.Printf("The values of Shanghai not exists")
	}
	value,isNull=map1["Anyang"]
	if isNull{
		fmt.Printf("The values of Anyang is %d\n",value)
	}else{
		fmt.Printf("The values of Anyang not exists\n")
	}
}
