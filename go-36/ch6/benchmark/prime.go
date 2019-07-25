package primes

import (
	"math"
)

func getPrimes(max int) []int {
	marks := make([]bool, max)
	if max <= 1 {
		return []int{}
	}
	squareRoot := int(math.Sqrt(float64(max)))
	var count int
	for i := 2; i <= squareRoot; i++ {
		for j := i * i; j < max; j += i {
			if marks[j] == false {
				marks[j] = true
				count++
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes

}
