package middleware

import (
	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/apps/user"
	"github.com/GeekQk/vblog/exception"
	"github.com/GeekQk/vblog/response"
	"github.com/gin-gonic/gin"
)

func Required(roles ...user.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 校验用户的信息
		// 直接通过上下文取出Token, 通过Token获取用户信息,来定义访问这个接口的角色
		// 通过Token获取用户角色信息 来定义接口访问的权限
		if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
			// 遍历判断 用户是否在运行的角色列表中
			hasPerm := false
			for _, r := range roles {
				if r == v.(*token.Token).Role {
					hasPerm = true
				}
			}
			if !hasPerm {
				response.Failed(c, exception.ErrPermissionDeny.WithMessagef("允许访问的角色: %v", roles))
				return
			}

		} else {
			response.Failed(c, exception.ErrUnauthorized)
			return
		}

	}
}
