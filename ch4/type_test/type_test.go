package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
	t.Log(*aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	if s == "" {
		t.Log("1")
	} else {
		t.Log("2")
	}
}