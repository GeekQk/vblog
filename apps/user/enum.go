package user

// 什么是枚举
type Role int

const (
	//当不传递值时，默认值从0开始
	ROLE_MEMBER Role = iota //成员
	ROLE_ADMIN              //管理员
)
