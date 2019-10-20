package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 2, 4}

	t.Log(a == b)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable &^ Writeable &^ Executable
	t.Log(a)
}
