package token

import "github.com/GeekQk/vblog/exception"

// 定义token模块异常
// 约定俗称:ErrXXXX 进行错误定义
var (
	ErrTokenInvalid       = exception.NewApiException(5000, "access token invalid")
	ErrRefreshTokenExpire = exception.NewApiException(5001, "refresh token invalid")
	ErrTokenNotExist      = exception.NewApiException(5002, "token not exist")
)
