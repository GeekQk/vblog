package token

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	AppName = "tokens"
)

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
	UserName string `json:"username"`
	Password string `json:"password"`
	RemindMe bool   `json:"remind_me"` //有效期为1周
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

// 先从header中获取token,如果header中没有则从cookie中获取
func GetTokenFromHttpHeader(req *http.Request) string {
	//自定义头, 获取token,标准头:Authorization
	at := req.Header.Get(TOKEN_HEADER_KEY)
	if at != "" {
		hv := strings.Split(at, " ")
		if len(hv) > 1 {
			return hv[1]
		}
		return at
	}
	//直接读Cookie
	ck, err := req.Cookie(TOKEN_COOKIE_KEY)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	val, _ := url.QueryUnescape(ck.Value)
	return val

}
