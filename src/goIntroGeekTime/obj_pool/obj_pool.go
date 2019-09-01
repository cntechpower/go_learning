package objpool

import (
	"errors"
	"time"
)

type employee struct {
	id   int
	name string
}

type employeePool struct {
	poolChan chan employee
}

var errOperTimeout = errors.New("Operation Timeout")
var errPoolOverflow = errors.New("Pool Overflow")

func newEmployeePool(num int) (employeePoolInstance *employeePool) {
	poolChan := make(chan employee, num)
	employeePoolInstance = &employeePool{poolChan}
	return employeePoolInstance
}

func (p *employeePool) get() (e employee, err error) {
	select {
	case employeeInstance := <-p.poolChan:
		return employeeInstance, nil
	case <-time.After(1 * time.Microsecond):
		return employee{}, errOperTimeout
	}
}

func (p *employeePool) push(e employee) (err error) {
	select {
	case p.poolChan <- e:
		return nil
	default:
		return errPoolOverflow
	}

}
