package impl

import (
	"context"

	"github.com/GeekQk/vblog/apps/token"
)

// 需要复用的数据操作
func (i *TokenServiceImpl) getToken(ctx context.Context, acessToken string) (*token.Token, error) {
	tk := token.NewToken(false)
	err := i.db.WithContext(ctx).Model(&token.Token{}).Where("access_token = ?", acessToken).First(&tk).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}
