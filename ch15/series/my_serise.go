package series

import "fmt"

/*
每个包里面都可以写init方法并且可以写多个init方法
在包被引入的时候就会执行init方法
go会根据包的依赖关系依次执行init方法
*/
func init() {
	fmt.Println("series init")
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

//这个是不能被外部调用的
func square(n int) int {
	return n * n
}
