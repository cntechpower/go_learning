package main
import "fmt"
type employee struct{
	name string
	salary int
}
func (em employee)AddSalary(percent int)(int){
//func (em *employee)AddSalary(percent int)(){
	return em.salary+em.salary*percent/100
	//em.salary=em.salary+em.salary*percent/100
}
func main(){
	em1:=employee{"dujinyang",7000}
	fmt.Printf("before add,em1's salary is: %v\n",em1.salary)
	em1.salary=em1.AddSalary(25)
	//em1.AddSalary(25)
	fmt.Printf("After add,em1's salary is: %v\n",em1.salary)
}
