package impl_test

import (
	//1.把对象加载到ioc
	"context"

	_ "github.com/GeekQk/vblog/apps"
	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/ioc"
)

// blog service 的实现的具体对象是在ioc中,
// 需要在ioc中获取具体的svc 用来测试

var (
	impl blog.Service
	ctx  = context.Background()
)

func init() {
	//2.从ioc中获取对象
	impl = ioc.Controller().Get(blog.AppName).(blog.Service)

	//3.ioc初始化 才会填充db数据
	if err := ioc.Init(); err != nil {
		panic(err)
	}

}
