package unittest

import "testing"

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3, 4}
	expected := [...]int{1, 4, 9, 15}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("Input is %d, Expected is %d, Actual is %d", inputs[i], expected[i], ret)
			t.Fatal()
		}
	}
	t.Logf("Test Complate\n")

}
