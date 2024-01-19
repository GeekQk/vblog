package user_test

import (
	"testing"

	"github.com/GeekQk/vblog/apps/user"
)

// $2a$10$Vwfajh51ApbWGmhLapYgLuw.jITxJfI2Bjdm3eZ2D1Xe6Bxw8WjAS
func TestHashWord(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Password = "123456"
	req.HashPassword()
	t.Log(req.Password)
	t.Log(req.CheckPassword("123456"))
	//1:33:51 PM
}
