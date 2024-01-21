package token

import "context"

type Service interface {
	//登录：颁发令牌
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)

	//退出: 销毁令牌
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)

	//验证: 校验令牌
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)

	//刷新: 刷新令牌
	RefreshToken(context.Context, *RefreshTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(userName, password string) *IssueTokenRequest {
	return &IssueTokenRequest{UserName: userName, Password: password}
}

// 颁发令牌请求
type IssueTokenRequest struct {
	UserName string
	Password string
	RemindMe bool
}

func NewRevokeTokenRequest(accessToken, refreshToken string) *RevokeTokenRequest {
	return &RevokeTokenRequest{AccessToken: accessToken, RefreshToken: refreshToken}
}

// 撤销令牌请求
type RevokeTokenRequest struct {
	//颁发的token 撤销
	AccessToken string
	//撤销的时候需要验证token的合法性
	RefreshToken string
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{AccessToken: accessToken}
}

// 验证令牌请求
type ValidateTokenRequest struct {
	AccessToken string
}

func NewRefreshTokenRequest(accessToken, refreshToken string) *RefreshTokenRequest {
	return &RefreshTokenRequest{AccessToken: accessToken, RefreshToken: refreshToken}
}

// 刷新令牌请求
type RefreshTokenRequest struct {
	AccessToken  string
	RefreshToken string
}
