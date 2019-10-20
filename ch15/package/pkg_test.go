package _package

import (
	"errors"
	cm "github.com/easierway/concurrent_map"
	"go-demo/ch15/series"
	"testing"
)

/*
关于package的几个要点：
1. 基本复用模块
	以首字母大写来表明可被包外代码访问
2. 代码的package可以和所在的目录名字不一样
3. 同一目录里的Go代码的package要保持一致
*/

func TestPkg(t *testing.T) {
	ret := series.GetFibonacciSeries(20)
	t.Log(ret)

	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("key"), 123)
	if ret, ok := m.Get(cm.StrKey("key13")); ok {
		t.Log(ret)
	} else {
		e := errors.New("not find")
		t.Log(e)
	}

}
