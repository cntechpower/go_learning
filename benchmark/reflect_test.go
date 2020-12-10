package benchmark

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

type Nums struct {
	N1 int
	N2 int
}

var t1 *Nums
var t2 *atomic.Value

func init() {
	n := rand.Int()
	t1 = &Nums{
		N1: n,
		N2: n,
	}
	t2 = &atomic.Value{}
	t3 := &Nums{
		N1: n,
		N2: n,
	}
	t2.Store(t3)
}

func updateT1InBackground() {
	for {
		n := rand.Int()
		t1 = &Nums{
			N1: n,
			N2: n,
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func updateT2InBackground() {
	for {
		n := rand.Int()
		t := &Nums{
			N1: n,
			N2: n,
		}
		t2.Store(t)
		time.Sleep(time.Millisecond)
	}
}

func BenchmarkDirectRead(b *testing.B) {
	go updateT1InBackground()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n1 := t1.N1
		n2 := t1.N2
		if n1 != n2 {
			b.Fatalf("n1 = %v, n2 = %v\n", n1, n2)
		}
	}
}

func BenchmarkAtomicRead(b *testing.B) {
	go updateT2InBackground()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := t2.Load().(*Nums)
		n1 := t.N1
		n2 := t.N2
		if n1 != n2 {
			b.Fatalf("n1 = %v, n2 = %v\n", n1, n2)
		}
	}
}
