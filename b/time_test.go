package b

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	y, m, d := time.Now().Date()
	//h := time.Now().Hour()
	fmt.Println(time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix())
	fmt.Println(time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location()).Unix())
}
