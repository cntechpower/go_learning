package main
import "fmt"

type person struct{
	age int
	sex int
	name string
}
func main(){
	p1:=new(person)
	p1.age=18
	p1.sex=0
	p1.name="dujinyang"
	fmt.Printf("Person 1's age is: %v\n",p1.age)
	if p1.sex==0{
		fmt.Printf("Person 1's sex is: Male\n")
	}else{
		fmt.Printf("Person 1's sex is: Female\n")
	}
	fmt.Printf("Person 1's Names is: %v\n",p1.age)
}
