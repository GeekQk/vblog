package impl_test

import (
	"context"
	"testing"

	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/apps/token/impl"
	ui "github.com/GeekQk/vblog/apps/user/impl"
)

var (
	i   *impl.TokenServiceImpl
	ctx = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("test", "123456")
	req.RemindMe = true
	tk, err := i.IssueToken(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(tk)
}

func TestRevokeToken(t *testing.T) {

}

func TestValidateToken(t *testing.T) {
}

func init() {
	//加载测试对象 i就是UserServiceImpl
	i = impl.NewTokenServiceImpl(ui.NewUserServiceImpl())
}
