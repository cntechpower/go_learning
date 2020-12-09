package primes

import "testing"

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPrimes(1000)
	}
}
