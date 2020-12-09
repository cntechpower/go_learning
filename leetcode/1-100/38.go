package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	p1 := 0
	cur := 1
	builder := strings.Builder{}
	str := countAndSay(n - 1)
	for ; cur < len(str); cur++ {
		if str[p1] != str[cur] {
			builder.WriteString(strconv.Itoa(cur - p1))
			builder.WriteByte(str[p1])
			p1 = cur
		}
	}
	if p1 != cur {
		builder.WriteString(strconv.Itoa(cur - p1))
		builder.WriteByte(str[p1])
	}
	return builder.String()

}

func main() {
	fmt.Println(countAndSay(4))
}
