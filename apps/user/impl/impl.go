package impl

import (
	"github.com/GeekQk/vblog/conf"
	"gorm.io/gorm"
)

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
