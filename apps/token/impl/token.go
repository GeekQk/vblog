package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/apps/user"
)

func (i *TokenServiceImpl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	//1.确认用户密码是否正确
	req := user.NewQueryUserRequest()
	req.Username = in.UserName
	us, err := i.user.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(us.Items)
	if len(us.Items) == 0 {
		return nil, fmt.Errorf("用户或者密码不存在1")
	}
	u := us.Items[0]
	if u.CheckPassword(in.Password); err != nil {
		return nil, fmt.Errorf("用户或者密码不存在2")
	}
	//2.账号密码正确 => 生成令牌
	tk := token.NewToken(in.RemindMe)
	tk.UserId = fmt.Sprintf("%d", u.Id)
	tk.UserName = u.Username
	tk.Role = u.Role

	//3.保存用户token
	err = i.db.WithContext(ctx).Model(&token.Token{}).Create(tk).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 退出: 销毁令牌
func (i *TokenServiceImpl) RevokeToken(ctx context.Context, in *token.RevokeTokenRequest) (*token.Token, error) {
	//1. 查询token 在进行删除
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("token 不不存在")
	}
	//2.refresh token 确认
	if !tk.CheckRefreshToken(in.RefreshToken) {
		return nil, fmt.Errorf("refresh token 不正确")
	}
	//3. 删除token
	err = i.db.WithContext(ctx).
		Where("access_token = ?", tk.AccessToken).
		Where("refresh_token = ?", tk.RefreshToken).
		Delete(&token.Token{}).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 验证: 校验令牌
// 依赖用户模块的来进行校验
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	//1. 查询token 在进行删除
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, token.ErrTokenNotExist.WithMessagef("token not exist,token: %v", in.AccessToken)
	}
	//2. 判断token 是否过期
	if tk.IsExpired() {
		return nil, token.ErrRefreshTokenExpire.WithMessagef("token expire %v", time.Now().Unix())
	}
	return tk, nil
}

// 验证: 刷新令牌
// 依赖用户模块的来进行校验
func (i *TokenServiceImpl) RefreshToken(ctx context.Context, in *token.RefreshTokenRequest) (*token.Token, error) {
	//1. 查询token
	return nil, nil
}
