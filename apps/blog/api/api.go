package api

import (
	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/ioc"
)

func init() {
	ioc.Api().Registery(blog.AppName, &blogApiHandler{})
}

// blog.Service接口, 是直接注册给Ioc, 不需要对我暴露
type blogApiHandler struct {
	svc blog.Service
}

func (i *blogApiHandler) Init() error {
	i.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}

func (i *blogApiHandler) Destory() error {
	return nil
}
