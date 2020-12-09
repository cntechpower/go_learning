package main

import (
	"fmt"
	"math"
)

func div(a, b int64) int64 {
	if a < b {
		return 0
	}
	count := int64(1)
	tb := b
	for (tb + tb) <= a {
		count = count + count
		tb = tb + tb
	}
	return count + div(a-tb, b)
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		if dividend >= math.MaxInt32 {
			return math.MaxInt32
		} else if dividend <= math.MinInt32 {
			return math.MinInt32
		} else {
			return dividend
		}
	}
	if divisor == -1 {
		if dividend >= math.MaxInt32 {
			return math.MinInt32
		} else if dividend <= math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}
	a := int64(dividend)
	b := int64(divisor)
	signed := false
	if a < 0 {
		signed = !signed
		a = -a
	}
	if b < 0 {
		signed = !signed
		b = -b
	}
	res := div(a, b)
	if signed {
		return int(-res)
	}
	return int(res)
}

func divide2(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		if dividend >= math.MaxInt32 {
			return math.MaxInt32
		} else if dividend <= math.MinInt32 {
			return math.MinInt32
		} else {
			return dividend
		}
	}
	if divisor == -1 {
		if dividend >= math.MaxInt32 {
			return math.MinInt32
		} else if dividend <= math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}
	var count int32
	signed := false
	if dividend < 0 {
		signed = !signed
		dividend = -dividend
	}
	if divisor < 0 {
		signed = !signed
		divisor = -divisor
	}
	if (dividend >= math.MaxInt32) &&
		((divisor == 1) || (divisor == -1)) {
		if signed {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	for dividend >= divisor {
		if count >= math.MaxInt32-1 {
			return math.MaxInt32
		}
		count++
		dividend -= divisor
	}
	if signed {
		return int(-count)
	}
	return int(count)
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
}
