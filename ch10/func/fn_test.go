package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(100), rand.Intn(20)
}

//测试函数的多返回值
func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second)
	return op
}

//测试function pragmas
func TestSlowFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	ret := tsSF(100)
	t.Log(ret)

	retss := timeSpent(slowFun)
	retss(123)
}

func MySum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

//测试可变参数的函数
func TestVarParam(t *testing.T) {
	t.Log(MySum(1, 2, 3))
	t.Log(MySum(4, 5, 6))
}

func ClearResource() {
	fmt.Println("clear....")
}

func TestDefer(t *testing.T) {
	/*
		相当于finally， 在函数结束后return之前会调用defer
		这里用了一个非命名函数
	*/
	defer ClearResource()

	//还有一种命名函数的写法如下
	defer func() {
		fmt.Println("defer....")
	}()
	fmt.Println("start")

	//程序错误
	panic("err")
}
