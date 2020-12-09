package recursive_test

import (
	"fmt"
	"testing"
)

func printN1(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func printN2(n int) {
	if n > 1 {
		printN2(n - 1)
	}
	fmt.Println(n)
}

func BenchmarkPrintN1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printN1(1000)
	}
}

func BenchmarkPrintN2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printN2(1000)
	}
}

func TestPrint(t *testing.T) {
	printN2(1000000)
}
