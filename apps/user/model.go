package user

import (
	"encoding/json"
	"fmt"
	"time"
)

// 用户创建请求
func NewUser(req *CreateUserRequest) *User {
	// 密码加密
	req.HashPassword()

	return &User{
		CreatedAt:         time.Now().Unix(),
		UpdatedAt:         time.Now().Unix(),
		CreateUserRequest: req,
	}
}

// 用户创建成功后返回一个User对象
// CreatedAt 为啥没用time.Time, int64(TimeStamp), 统一标准化, 避免时区你的程序产生影响
// 在需要对时间进行展示的时候，由前端根据具体展示那个时区的时间
type User struct {
	// 用户Id
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 创建时间, 时间戳 10位, 秒
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新时间, 时间戳 10位, 秒
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// 用户参数
	*CreateUserRequest
}

// 定义存储表名称
func (u *User) TableName() string {
	return "users"
}

func (c *User) String() string {
	sr, err := json.MarshalIndent(c, " ", "  ")
	if err != nil {
		return fmt.Sprintf("%s", err)
	}
	return string(sr)
}
