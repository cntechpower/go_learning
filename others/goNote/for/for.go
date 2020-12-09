package main

import "fmt"

func main() {
	data := []int{10, 20, 30}
	for i, x := range data {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("No%v range value: %v, actual value is %v\n", i, x, data[i])
	}
	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("No%v range value: %v, actual value is %v\n", i, x, data[i])
	}

}
