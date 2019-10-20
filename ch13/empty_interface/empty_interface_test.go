package empty_interface

import (
	"fmt"
	"testing"
)

func AsserType(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float64", v)
	default:
		fmt.Println("Unknown type")
	}
}

/*
用这个例子测试一个空接口可以传入任何类型，像是Java里面的Object一样
*/
func TestEmptyInterface(t *testing.T) {
	AsserType(10)
	AsserType("sdf")
	AsserType(1.22)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

type ReaderWriterImpl struct {
}

func (r *ReaderWriterImpl) Read(p []byte) (n int, err error) {
	fmt.Println(p)
	return n, nil
}

func (r *ReaderWriterImpl) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	return n + 1, nil
}

type Body struct {
	ReaderWriterImpl
}

//这一部分主要演示接口和多态
/*
这个例子里的意思相当于是：
ReaderWriter继承了Reader和Writer两个接口
ReaderWriterImpl实现了ReaderWriter接口
Body继承了ReaderWriterImpl
*/
func TestInterface(t *testing.T) {
	b := []byte{100, 222}
	body := new(Body)
	read, err := body.Read(b)
	write, err := body.Write(b)
	fmt.Println(read, err)
	fmt.Println(write, err)
}
