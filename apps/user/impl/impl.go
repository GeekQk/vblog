package impl

import (
	"github.com/GeekQk/vblog/apps/user"
	"github.com/GeekQk/vblog/conf"
	"github.com/GeekQk/vblog/ioc"
	"gorm.io/gorm"
)

// 为什么在注册的时候没有初始化
// 这个逻辑是在import的时候操作，程序在import后才指定的配置文件加载,此时的conf.C为nil
// 所以这个db属性的初始化一定要在配置加载后才会执行:conf.Load()->ioc.Init()
// init的机制 就是在import的时候执行
func init() {
	ioc.Controller().Registery(user.AppName, &UserServiceImpl{})
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		// 获取全局的DB对象
		db: conf.C().DB(),
		// db: conf.C().MySql.GetDB(),
	}
}

// 怎么实现user.Service接口?
// 定义UserServiceImpl来实现接口
type UserServiceImpl struct {
	// 注入db 依赖一个数据库连接对象
	db *gorm.DB
}

// 先程序加载init->补充属性db
func (i *UserServiceImpl) Init() error {
	i.db = conf.C().DB()
	return nil
}

func (i *UserServiceImpl) Destory() error {
	return nil
}
