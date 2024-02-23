package api

import (
	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/apps/user"
	"github.com/GeekQk/vblog/ioc"
	"github.com/GeekQk/vblog/middleware"
	"github.com/gin-gonic/gin"
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

// 让ioc帮我们完成接口的路由注册 ioc.GinApi
//
//	type GinApi interface {
//		// 基础约束
//		Object
//		// 额外约束
//		// ioc.Api().Get(token.AppName).(*api.TokenApiHandler).Registry(rr)
//		Registry(rr gin.IRouter)
//	}

func (i *blogApiHandler) Registery(rr gin.IRouter) {
	r := rr.Group(blog.AppName)

	// 普通接口, 允许访客使用, 不需要权限
	r.GET("/", i.QueryBlog)
	r.GET("/:id", i.DescribeBlog)

	//中间件使用的两种方式
	//方式一:影响单个路由, 直接把中间件函数注册到路由上
	//r.GET("/", middleware.Auth, i.QueryBlog)
	//方式二: 影响后面全部,在路由之前进行调用
	//r.Use(middleware.Auth)
	//r.GET("/", i.QueryBlog)

	// 整个模块后面的请求 都需求认证
	//Gin的中间件的实现是根据顺序来执行的, 只会影响后面的Action
	r.Use(middleware.Auth)

	// 后面管理的接口 需要权限
	r.POST("/", i.CreateBlog)
	r.PATCH("/:id", i.PatchBlog)
	r.PUT("/:id", i.UpdateBlog)
	r.DELETE("/:id", middleware.Required(user.ROLE_ADMIN, user.ROLE_MEMBER), i.DeleteBlog)
}
