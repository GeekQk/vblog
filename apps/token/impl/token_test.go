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
	req := token.NewIssueTokenRequest("test1", "123456")
	req.RemindMe = true
	tk, err := i.IssueToken(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(tk)
}

func TestRevokeToken(t *testing.T) {
	req := token.NewRevokeTokenRequest("cmm79cdiika5l5iol6bg", "cmm79cdiika5l5iol6c0")
	tk, err := i.RevokeToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)

}

func TestValidateToken(t *testing.T) {
}

func init() {
	//加载测试对象 i就是UserServiceImpl
	i = impl.NewTokenServiceImpl(ui.NewUserServiceImpl())
}
