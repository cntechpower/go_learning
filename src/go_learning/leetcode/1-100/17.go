package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
)

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) < 1 {
		return res
	}
	mapping := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	first, _ := strconv.Atoi(string(digits[0]))
	if first < 2 || first > 9 {
		return res
	}
	res = append(res, mapping[first]...)

	for i := 1; i < len(digits); i++ {
		num, _ := strconv.Atoi(string(digits[i]))
		if first < 2 || first > 9 {
			return res
		}
		var tmpRes []string
		for _, left := range res {
			for _, curr := range mapping[num] {
				tmpRes = append(tmpRes, left+curr)
			}
		}
		res = tmpRes
		tmpRes = nil

		fmt.Printf("----------------Round %v-------------------------------------------------------\n", i)
		fmt.Println("Before gc and free")
		printMemStats()
		runtime.GC()
		debug.FreeOSMemory()
		fmt.Println("After gc and free")
		printMemStats()

	}
	return res
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024, m.HeapReleased/1024)
}

func main() {
	fmt.Println(letterCombinations("23324734326487324637482374892375"))
	fmt.Println(letterCombinations("2"))
}
