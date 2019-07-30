package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seq := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	l := len(seq)
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(l * l * l)
	fmt.Printf("Length is %d, rand is %d\n", l, a)
	s := fmt.Sprintf("%c%c%c", seq[a%l], seq[(a/l)%l], seq[(a/l/l)%l])
	fmt.Println(s)
}
