package ip

import (
	"fmt"
	"strings"
	"testing"
)

func TestDistinct(t *testing.T) {
	res := strings.Builder{}
	for i := 0; i < 100; i++ {
		res.WriteString(fmt.Sprintf("select mid from rewards_award_record_%02d UNION ", i))
	}
	fmt.Println(res.String())
}
