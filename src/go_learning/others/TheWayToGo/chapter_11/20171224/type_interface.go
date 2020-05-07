package main
import (
	"fmt"
	"math"
)
type Square struct{
	side float32
}
type Circle struct{
	radius float32
}
type Shaper interface{
	Area() float32
}
func (sq *Square)Area()float32{
	return sq.side*sq.side
}
func (ci *Circle)Area()float32{
	return ci.radius*ci.radius*math.Pi
}
func main(){
	var areaIntf Shaper
	sq1:=new(Square)
	sq1.side=5
	areaIntf=sq1
	if t,ok:=areaIntf.(*Square);ok{
		fmt.Printf("The type of areaIntf is: %T\n",t)
		fmt.Printf("The area of areaIntf is: %v\n",areaIntf.Area())
		return 
	}
	if t,ok:=areaIntf.(*Circle);ok{
		fmt.Printf("The type of areaIntf is: %T\n",t)
	}else{
		fmt.Printf("Type doesnot match circle\n")
	}
}
