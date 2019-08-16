package main

import "fmt"

type testStruct struct {
	retMessage string
}

const MAX_RPC_CONNECT_RETRY = 5

func loop(ts *testStruct) error {
	for retry := 0; retry < MAX_RPC_CONNECT_RETRY; retry++ {
		err := fmt.Errorf("error")
		if true {
			if retry == MAX_RPC_CONNECT_RETRY-1 {
				return err
			}
			fmt.Println("called")
			continue //retry
		}
		return err
	}
	return nil
}

func main() {
	var ts *testStruct
	if err := loop(ts); err != nil {
		fmt.Printf("call loop error, %v\n", err)
	} else {
		fmt.Println(ts.retMessage)
	}
}
