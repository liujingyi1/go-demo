package Polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	ShowMessage(msg string) Code
	DoSomthing() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) ShowMessage(msg string) Code {
	return Code(fmt.Sprintf("go message:%s", msg))
}

func (p *GoProgrammer) DoSomthing() string {
	return ""
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) ShowMessage(msg string) Code {
	return Code(fmt.Sprintf("java message:%s", msg))
}

func (p *JavaProgrammer) DoSomthing() string {
	return ""
}

func WriteShowMessage(p Programmer) {
	fmt.Printf("%T %s\n", p, p.ShowMessage("hello"))
}

//当接口里有两个方法时，实现类必须实现所有的方法
func TestPolymorphism(t *testing.T) {
	goProgrammer := &GoProgrammer{}
	javaProgrammer := new(JavaProgrammer)
	WriteShowMessage(goProgrammer)
	WriteShowMessage(javaProgrammer)
}
