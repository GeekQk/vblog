package ioc_test

import (
	"fmt"
	"testing"

	"github.com/GeekQk/vblog/ioc"
)

func TestContainerGetReg(t *testing.T) {
	c := ioc.NewContainer()
	c.Registery("test", &TestStruct{})
	t.Log(c.Get("test"))
	//断言使用
	c.Get("test").(*TestStruct).XXX()
}

type TestStruct struct {
}

func (*TestStruct) Init() error {
	return nil
}

func (*TestStruct) Destory() error {
	return nil
}

func (*TestStruct) XXX() error {
	fmt.Println("xxx log")
	return nil
}
