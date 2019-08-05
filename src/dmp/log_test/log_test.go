package log_test

import (
	"fmt"
	"strings"
	"testing"
)

var sample = []string{
	"password=", "Password=", "PASSWORD=", "password\"=", "Password\"=", "PASSWORD\"=",
	"password:", "Password:", "PASSWORD:", "password\":", "Password\":", "PASSWORD\":",
	"password(", "Password(", "PASSWORD(", "password\"(", "Password\"(", "PASSWORD\"(",
	"identified by ", "IDENTIFIED BY ", "-P ", "password\">", "password\" ",
}

func TestStringFilter(t *testing.T) {
	line := `"password":123 abcd`
	for _, key := range sample {
		fmt.Printf("Index of %v\n", key)
		idx := strings.Index(line, key)
		if idx == -1 {
			continue
		} else {
			fmt.Printf("Index is %d\n", idx)
			idxNextPart := idx + len(key)
			fmt.Printf("idxNextPart is %c\n", line[idxNextPart])
			fmt.Printf("idxNextPart -1 is %c\n", line[idxNextPart-1])
			fmt.Printf("nextPart is %v\n", line[idxNextPart:])
			break
		}
	}
}
