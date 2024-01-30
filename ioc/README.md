# ioc 简单实现


[ioc设计](../docs/ioc.drawio)

功能:
1.  对象托管:
    + 对象的注册
    + 对象获取
2. 对象统一管理:
    + 配置
    + 初始化
    + 销毁


## 基于Map来实现Ioc容器


### 单容器设计

基于Map来实现一个这样的 对象托管容器
```go
container = map[string]any
```

注册的对象必须进行接口约束, 要求他必须实现 对象统一管理的方法:
```go
// 对象统一管理
type IocObject interface {
    Init() error
    Destory() error
}


type TestStruct struct {

}

func (t *TestStruct) Init() error {

}

func (t *TestStruct) Destory() error {

}

func (t *TestStruct) XXX() error {

}

// container = map[string]IocObject
ioc.Registry("service name", &TestStruct{})

// 启动的时候, 完成对象的统一管理, 循环容器里面的所有对象, 调用的Init方法
ioc.Init()
```

一个map, 不允许重名的,  有一个模块叫token：
+ TokenServiceImpl: Controller
+ TOkenApiHandler: Api

```go
// 注册控制器
ioc.Controller.Registry("module name", &TokenServiceImpl{})
ioc.Api.Registry("module name", &TOkenApiHandler{})
```

根据程序设计，对这些对象的 职责约束, 将容器进行分区:
+ Api: 负责注册 Api实现类型的对象
+ Controller: 负责注册服务实现类的对象
+ Config: 配置对象, db,kafak,redis
+ Default: 预留区域

### 多容器

```go
api_contailer = map[string]IocObject
controller_container = map[string]IocObject
```

### 实现ioc

封装Container
```go
func TestContainerGetAndRegistry(t *testing.T) {
	c := ioc.NewContainer()
	c.Registry("TestStruct", &TestStruct{})
	t.Log(c.Get("TestStruct"))

	// 断言使用
	c.Get("TestStruct").(*TestStruct).XXX()
}
```

封装Manager
```go
func TestManageGetAndRegistry(t *testing.T) {
	ioc.Controller().Registry("TestStruct", &TestStruct{})
	t.Log(ioc.Controller().Get("TestStruct"))

	// 断言使用
	ioc.Controller().Get("TestStruct").(*TestStruct).XXX()
}
```