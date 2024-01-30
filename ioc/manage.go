package ioc

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
