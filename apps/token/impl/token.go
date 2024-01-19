package impl

import (
	"context"
	"fmt"

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
	return nil, nil
}

// 验证: 校验令牌
// 依赖用户模块的来进行校验
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, in *token.ValidateToken) (*token.Token, error) {

	return nil, nil
}
