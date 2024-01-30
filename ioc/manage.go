package ioc

import "github.com/gin-gonic/gin"

// ioc管理

// ioc全局管理
var nc = &NameSpaceContainer{
	ns: map[string]*Container{
		"api":        NewContainer(),
		"controller": NewContainer(),
		"config":     NewContainer(),
		"default":    NewContainer(),
	},
}

// ioc全局管理-自动init
func Init() error {
	return nc.InIt()
}

// ioc全局管理-自动Destory
func Destory() error {
	return nc.Destory()
}

// 外部直接获取
// ioc.Controller.Registery()
// ioc.Controller.Get()
func Controller() *Container {
	return nc.ns["controller"]
}

// ioc.Api.Registery()
// ioc.Api.Get()
func Api() *Container {
	return nc.ns["api"]
}

func RegisteryGinApi(rr gin.IRouter) {
	nc.RegisteryGinApi(rr)
}

// 基于这个构建多空间的container
type NameSpaceContainer struct {
	ns map[string]*Container
}

func (n *NameSpaceContainer) InIt() error {
	//遍历namespace
	for key := range n.ns {
		c := n.ns[key]
		//遍历namespace container
		for ObjectName := range c.storage {
			if err := c.storage[ObjectName].Init(); err != nil {
				return err
			}
		}
	}
	return nil

}

func (n *NameSpaceContainer) Destory() error {
	//遍历namespace
	for key := range n.ns {
		c := n.ns[key]
		//遍历namespace container
		for ObjectName := range c.storage {
			if err := c.storage[ObjectName].Destory(); err != nil {
				return err
			}
		}
	}
	return nil
}

type GinApi interface {
	Object //基础约束
	//ioc.Api().Get(token.AppName).(*api.TokenApiHandle).Registery(rr)
	Registery(rr gin.IRouter)
}

// 遍历所有Api对象 完成对象注册
// ioc调用对象提供的Registery方法 来完成模块Api注册给gin.IRouter
func (n *NameSpaceContainer) RegisteryGinApi(rr gin.IRouter) {
	// 遍历namespace
	for key := range n.ns {
		c := n.ns[key]
		//遍历namespace container
		for ObjectName := range c.storage {
			obj := c.storage[ObjectName]
			//如何判断一个对象是否是GinApp对象
			//判断对象有咩有Registery方法
			//操作:断言该对象是否满足GinApi接口,实现了Registery函数
			if v, ok := obj.(GinApi); ok {
				v.Registery(rr)
			}
		}
	}
}
