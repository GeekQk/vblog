package impl

import (
	"github.com/GeekQk/vblog/apps/blog"
	"github.com/GeekQk/vblog/conf"
	"github.com/GeekQk/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registery(blog.AppName, &blogServiceImpl{})
}

// blog.service接口，直接注册给ioc,不需要对外暴露
type blogServiceImpl struct {
	// 注入db 依赖一个数据库连接对象
	db *gorm.DB
}

func (b *blogServiceImpl) Init() error {
	b.db = conf.C().DB()
	return nil
}

func (b *blogServiceImpl) Destory() error {
	return nil
}
