package impl

import (
	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/apps/user"
	"github.com/GeekQk/vblog/conf"
	"gorm.io/gorm"
)

// 可以不写 但是为了规范 还是写上,规范实现的方法
var (
	_ token.Service = (*TokenServiceImpl)(nil)
)

func NewTokenServiceImpl(userServiceImpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		// 获取全局的DB对象
		db: conf.C().DB(),
		// 依赖user.Service接口
		user: userServiceImpl,
	}
}

// 怎么实现user.Service接口?
// 定义UserServiceImpl来实现接口
type TokenServiceImpl struct {
	// 注入db 依赖一个数据库连接对象
	db *gorm.DB

	//依赖user.Service接口
	//依赖接口 不依赖接口的具体实现
	user user.Service
}
