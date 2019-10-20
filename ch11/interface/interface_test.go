package _interface

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

type Go2Programmer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "hello world"
}

func (g *Go2Programmer) WriteHelloWorld() string {
	return "hello home"
}

func TestInterface(t *testing.T) {
	var i int
	t.Log(i)

	p := new(Go2Programmer)
	p2 := GoProgrammer{}

	t.Logf("p type is %T", p)
	t.Logf("p2 type is %T", p2)

	ret := p2.WriteHelloWorld()
	t.Log(ret)
}
