package ioc_test

import (
	"testing"

	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/ioc"
)

func TestManageGetReg(t *testing.T) {

	//测试Controller
	ioc.Controller().Registery("TestStruct", &TestStruct{})
	t.Log(ioc.Controller().Get("TestStruct"))
	// t.Log(ioc.Controller().Get("TestStruct").(*TestStruct).XXX())

	//测试Api
	ioc.Api().Registery("TestStruct", &TestStruct{})
	t.Log(ioc.Api().Get("TestStruct"))
	// t.Log(ioc.Api().Get("TestStruct").(*TestStruct).XXX())

	ioc.Controller().Registery(token.AppName, &TestStruct{})
	t.Log(ioc.Controller().Get(token.AppName))

}
