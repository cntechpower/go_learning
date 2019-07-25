package main

import "fmt"

func main() {
	str := "172.10.1.1"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
	str1 := "Go 爱好者 "
	fmt.Printf("The string: %q\n", str1)
	fmt.Printf("  => runes(char): %q\n", []rune(str1))
	fmt.Printf("  => runes(hex): %x\n", []rune(str1))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str1))
	str2 := "0"
	fmt.Printf("The string: %q\n", str2)
	fmt.Printf("  => runes(char): %q\n", []rune(str2))
	fmt.Printf("  => runes(hex): %x\n", []rune(str2))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str2))
	var a int
	fmt.Println(a)
}
