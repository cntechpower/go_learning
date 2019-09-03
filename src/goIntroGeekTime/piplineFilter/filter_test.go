package piplineFilter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	pipeline := make([]Filter, 0)
	pipeline = append(pipeline, NewSplitFilter(","))
	pipeline = append(pipeline, NewToIntFilter())
	pipeline = append(pipeline, NewSumFilter())
	var err error
	var res Response
	res = "1,2,3"
	for _, filter := range pipeline {
		res, err = filter.Process(res)
		if err != nil {
			t.Errorf("process error : %v", err)
		}
	}
	fmt.Println(res)
}
