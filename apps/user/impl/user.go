package impl

import (
	"context"

	"github.com/GeekQk/vblog/apps/user"
)

// 实现 user.Service
// 怎么判断这个服务有没有实现这个接口喃？
// &UserServiceImpl{} 是会分配内存, 怎么才能不分配内存
// nil 如何生命 *UserServiceImpl 的 nil
// (*UserServiceImpl)(nil) ---> int8 1  int32(1)  (int32)(1)
// nil 就是一个*UserServiceImpl的空指针
var _ user.Service = (*UserServiceImpl)(nil)

// 用户创建
func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	in *user.CreateUserRequest) (
	*user.User, error) {
	//1.校验用户请求参数
	if err := in.Validate(); err != nil {
		return nil, err
	}

	//2.创建用户实例
	u := user.NewUser(in)
	//3.把对象持久化  存储表、映射关系
	//比如createUser 4s的时候,请求没有返回,用户取消了请求,后端会因为请求中断而结束？
	//怎么解决这个问题？ 使用WithContext 上下文
	if err := i.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}
	//4.返回用户实例
	return u, nil
}

// 查询用户列表, 对象列表 [{}]
func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	//构建一下mysql查询条件 select * from user where id > 0 offset 0 limit 10
	query := i.db.WithContext(ctx)
	if in.Username != "" {
		query = query.Where("username = ?", in.Username)
	} else {
		query = query.Limit(in.Limit()).Offset(in.OffSet()).Order("id asc")
	}
	//获取总页数
	userList := user.NewUserSet()
	err := query.Model(&user.User{}).Count(&userList.Total).Error
	if err != nil {
		return nil, err
	}
	//分页查询 limit 10 offset 0
	//limit 0,20
	//limit 20,20
	err = query.Find(&userList.Items).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

// 查询用户详情, 通过Id查询,
func (i *UserServiceImpl) DescribeUser(
	ctx context.Context,
	in *user.DescribeUserRequest) (
	*user.User, error) {

	u := user.NewUser(user.NewCreateUserRequest())
	err := i.db.WithContext(ctx).Model(&user.User{}).Where(in.UserId).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (i *UserServiceImpl) UpdateUser(
	ctx context.Context,
	in *user.AlterUserRequest) (*user.User, error) {
	return nil, nil
}

func (i *UserServiceImpl) DeleteUser(
	ctx context.Context,
	in *user.DescribeUserRequest) (bool, error) {
	return true, nil
}
