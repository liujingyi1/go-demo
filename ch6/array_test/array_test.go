package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	//for i := 0; i < len(arr); i++ {
	//	t.Log(arr[i])
	//}

	for idx, v := range arr {
		t.Logf("%d %d", idx, v)
	}

	i := 0
	for i < len(arr) {
		t.Log(arr[i])
		i++
	}

}
