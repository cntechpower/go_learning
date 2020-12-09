package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(str string) int {
	init := true
	reverse := false
	res := 0
	for _, r := range str {
		if unicode.IsSpace(r) && init {
			continue
		}
		if (r == 45 || r == 43) && init {
			if r == 45 {
				reverse = true
			}
			init = false
		} else if 48 <= r && r <= 57 {
			init = false
			if int(r-48) > math.MaxInt32-res*10 {
				if reverse {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			res = res*10 + int(r-48)
		} else {
			break
		}
	}
	if reverse {
		return -res
	}
	return res
}

func main() {
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("-42"))
	//fmt.Println(strconv.Atoi("words and 987")) //strconv.Atoi: parsing "words and 987": invalid syntax
}
