package type_test

import "testing"

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(b)

}

func TestPtr(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T  %T", a, aPtr)
}
