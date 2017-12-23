package main
import "fmt"
type TwoInts struct{
	a int
	b int
}
func (tn *TwoInts)AddThem()int{
	return tn.a+tn.b
}
func (tn *TwoInts)AddPra(pra int)int{
	return tn.a+tn.b+pra
}
func main(){
	two1:=new(TwoInts)
	two1.a=1
	two1.b=2
	fmt.Printf("The Sum of two1 is: %v\n",two1.AddThem())
	fmt.Printf("The Sum of two1 is: %v\n",two1.AddPra(30))
}

