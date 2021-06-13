package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	cidrs := strings.Split("", ",")
	fmt.Println(len(cidrs))
}
