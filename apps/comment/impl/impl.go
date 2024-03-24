package impl

import (
	"github.com/GeekQk/vblog/apps/comment"
	"github.com/GeekQk/vblog/conf"
	"github.com/GeekQk/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registery(comment.AppName, &commentServiceImpl{})
}

// comment.Service接口, 是直接注册给Ioc, 不需要对我暴露
type commentServiceImpl struct {
	// 要注册给grpc 必须继承后覆盖
	comment.UnimplementedServiceServer

	// 依赖了一个数据库操作的链接池对象
	db *gorm.DB
}

func (i *commentServiceImpl) Init() error {
	i.db = conf.C().DB()

	comment.RegisterServiceServer(conf.C().GrpcServer.GetServer(), i)
	return nil
}

func (i *commentServiceImpl) Destory() error {
	return nil
}
