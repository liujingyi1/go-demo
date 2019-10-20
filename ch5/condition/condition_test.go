package condition

import (
	"fmt"
	"runtime"
	"testing"
)

func swap(x, y string) (string, string) {
	return y, x
}

func TestIfMultiSec(t *testing.T) {

	if a, b := swap("a", "b"); b == "a" {
		t.Logf("111 %s %s", a, b)
	} else {
		t.Logf("222 %s %s", a, b)
	}
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux.")
	default:
		t.Logf("%s", os)
	}

	showRange(5)
}

func showRange(num int) {
	switch {
	case 0 <= num && num <= 3:
		fmt.Println("0-3")
	case 4 <= num && num <= 6:
		fmt.Println("4-6")
	case 6 <= num:
		fmt.Println(">6")
	default:
		fmt.Println("not def")

	}
}
