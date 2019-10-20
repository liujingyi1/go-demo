package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 2)
	t.Log(len(s0), cap(s0))
	t.Log(s0[0])

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	t.Log(s1[0], s1[1], s1[2], s1[3])

	t.Logf("a的地址:%x", &s1[0])

	s1 = append(s1, 1, 2)
	t.Log(len(s1), cap(s1))
	t.Log(s1[0], s1[1], s1[2], s1[3], s1[4], s1[5])
	t.Logf("a的地址:%x", &s1[0])

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 10)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

	arr := make([]int, 8, 4)
	t.Log(arr)
}
