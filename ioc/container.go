package ioc

//ioc实现

func NewContainer() *Container {
	return &Container{storage: make(map[string]Object)}
}

type Container struct {
	storage map[string]Object
}

// 初始化容器注册
func (c *Container) Registery(name string, obj Object) {
	c.storage[name] = obj
}

// 获取的值不需要约束，因为可以获取到任何类型，但是需要使用者自己处理断言
// ioc.Get("module name").(*TestService)
func (c *Container) Get(name string) any {
	if _, ok := c.storage[name]; !ok {
		return nil
	}
	return c.storage[name]
}

// 提供一个对象遍历的方法
// 用户传递一个回调函数，回调函数中可以获取到对象名和对象值处理
func (c *Container) ForEatch(cb func(objectName string, objectValue Object)) {
	for key := range c.storage {
		cb(key, c.storage[key])
	}
}
