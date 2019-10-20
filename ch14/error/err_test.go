package error

import (
	"errors"
	"fmt"
	"testing"
)

/*
panic是使程序崩溃的错误，在程序退出时会打出堆栈信息
对于panic错误程序退出前如果有defer，将会执行panic
recover相当于catch，会让程序继续执行下去
*/
/*
但是在实际编码时要注意两点:
当心recover成为恶魔
1. 形成僵尸服务程序，导致health check失效(health check 程序会检查程序是否正常
运行，而检查的方法时判断程序是否在执行，如果程序已经不能正常工作了，但是还在执行，
那health check就检查不到有问题的程序，不如直接让程序crash掉，然后health check会
重启程序，而重启程序可能会让程序恢复正常)
2. "Let it Crash!"往往是我们恢复不确定性错误的最好方法
*/
func TestPanicVxExit(t *testing.T) {
	defer func(i int) {
		t.Log(i)
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}(110)
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}
