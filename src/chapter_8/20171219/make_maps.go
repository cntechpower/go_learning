package main
import "fmt"

func main(){
	var mapLit map[string] int
	var mapAssigned map[string]int
	mapLit=map[string]int{"one":1,"two":2}
	mapCreated:=make(map[string]float32)
	mapAssigned=mapLit
	mapCreated["key1"]=4.5
	mapCreated["key2"]=3.2425926
	mapAssigned["two"]=222
	fmt.Printf("mapLit one is: %d\n",mapLit["one"])
	fmt.Printf("mapCreated key2is: %f\n",mapCreated["key2"])
	fmt.Printf("mapAssigned two is: %d\n",mapAssigned["two"])
	fmt.Printf("mapLit ten is: %d\n",mapLit["ten"])
}
