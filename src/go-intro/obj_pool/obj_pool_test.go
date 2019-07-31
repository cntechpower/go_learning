package objpool

import (
	"fmt"
	"testing"
)

func TestObjPool(t *testing.T) {
	p := newEmployeePool(10)
	//e1 := employee{1, "em1"}
	for i := 0; i < 15; i++ {
		ename := fmt.Sprintf("em%d", i)
		e1 := employee{i, ename}
		err := p.push(e1)
		if err != nil {
			fmt.Printf("Push No %d To Pool Error: %v\n", i, err)
		}
	}
	for i := 0; i < 15; i++ {
		pe1, err := p.get()
		if err != nil {
			fmt.Printf("Get No %d Employee From Pool Error: %v\n", i, err)
		} else {
			fmt.Println(pe1.id, pe1.name)
		}
	}
}
