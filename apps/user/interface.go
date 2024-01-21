package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	AppName = "users"
)
var vali = validator.New()

// 面向对象
// []*User和*[]User的区别
// []*User 是一个指针数组，*[]User 是一个指针，指向一个数组
// []*User 可以单独使用，*[]User 必须和指针一起使用

// user.Service, 设计你这个模块提供的接口
// 接口定义, 一定要考虑兼容性, 接口的参数不能变
type Service interface {
	// 用户创建
	// CreateUser(username, password, role string, lable map[string]string)
	// 设计CreateUserRequest, 可以扩展对象, 而不影响接口的定义
	// 1. 这个接口支持取消吗? 要支持取消应该怎么办?
	// 2. 这个接口支持Trace, TraceId怎么传递？
	// 中间件参数，取消/Trace/... 怎么产生怎么传递
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表, 对象列表 [{}]
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// 查询用户详情, 通过Id查询,
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	// 用户修改
	UpdateUser(context.Context, *AlterUserRequest) (*User, error)
	// 用户删除
	DeleteUser(context.Context, *DescribeUserRequest) (bool, error)
}

// 为了避免对象出现空指针,指针对象为初始化,搞一个构造函数
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  ROLE_MEMBER, //默认值0 不写也可以
		Label: map[string]string{},
	}
}

// 用户创建的参数
type CreateUserRequest struct {
	Username string            `json:"username" validate:"required" gorm:"column:username"`
	Password string            `json:"password" validate:"required" gorm:"column:password"`
	Role     Role              `json:"role" gorm:"column:role"`
	Label    map[string]string `json:"lable" gorm:"column:label;serializer:json"` // 序列map化成json字符串
}

func (c *CreateUserRequest) HashPassword() {
	hp, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	c.Password = string(hp)
}

func (c *CreateUserRequest) CheckPassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(pass))
}

// 引入validator, 验证参数
// go get github.com/go-playground/validator/v10
func (req *CreateUserRequest) Validate() error {
	//生成校验器对象 就可以使用规则验证
	return vali.Struct(req)
}

// 默认值
func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

// 查询用户列表
type QueryUserRequest struct {
	// 分页大小, 一个多少个
	PageSize int
	// 当前页, 查询哪一页的数据
	PageNumber int
	// 更加用户name查找用户
	Username string
}

func (req *QueryUserRequest) Limit() int {
	return req.PageSize
}

// 0,20
// 20,40
func (req *QueryUserRequest) OffSet() int {
	return req.PageNumber * (req.PageSize - 1)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

type UserSet struct {
	// 总共有多少个
	Total int64 `json:"total"`
	// 当前查询的数据清单
	Items []*User `json:"items"`
}

func (c *UserSet) String() string {
	sr, err := json.MarshalIndent(c, " ", "  ")
	if err != nil {
		return fmt.Sprintf("%s", err)
	}
	return string(sr)
}

func NewDescribeUserRequest(uid int) *DescribeUserRequest {
	return &DescribeUserRequest{uid}
}

type DescribeUserRequest struct {
	UserId int
}

// 修改用户的字段
type AlterUserRequest struct {
	Username string
	Password string
	Role     string
	Label    map[string]string
}
