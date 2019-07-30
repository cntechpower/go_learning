package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestExitPanic(t *testing.T) {
	defer func() {
		//fmt.Println("Defer Called")
		if err := recover(); err != nil {
			fmt.Println("recoverd from ", err)
		}
	}()
	fmt.Println("Start")
	// os.Exit(-1)
	panic(errors.New("Test Panic"))
}
