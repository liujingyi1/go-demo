package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 7 {
		t.Log(n)
		n++
	}
}
