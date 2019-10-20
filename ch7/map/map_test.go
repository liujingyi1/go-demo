package my_map

import "testing"

func TestInitMap(t *testing.T) {
	map1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(map1)

	map2 := map[int]int{}
	t.Log(map2)
	map2[5] = 6
	t.Log(map2)

	map3 := make(map[int]int, 10)
	t.Log(map3)
}

func TestMapKeyExist(t *testing.T) {
	map1 := map[int]bool{}

	map1[1] = true
	if v, exist := map1[1]; exist {
		t.Log(v)
		t.Log(map1[1])
	} else {
		t.Log("key 1 is not exist")
	}

	if map1[1] {
		t.Log("map3 is exist")
	} else {
		t.Log("mp3 is not exist")
	}
}

func TestTraveArray(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	//for i := 0; i < len(arr); i++ {
	//	t.Log(arr[i])
	//}
	//
	//for idx, v := range arr {
	//	t.Logf("%d %d", idx, v)
	//}

	i := 0
	for i < len(arr) {
		t.Log(arr[i])
		i++
	}

	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arr1 {
		t.Log(i, v)
	}
}

func TestTravlMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	for k, v := range m1 {
		t.Log(k, v)
	}

}

//测试map的值为一个函数
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	t.Log(m[1](9))
}

func factory(int) int {
	return 0
}
