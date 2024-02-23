package middleware

import (
	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/apps/user"
	"github.com/GeekQk/vblog/exception"
	"github.com/GeekQk/vblog/ioc"
	"github.com/GeekQk/vblog/response"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// 获取tk 模块
	svc := ioc.Controller().Get(token.AppName).(token.Service)

	ak := token.GetTokenFromHttpHeader(c.Request)
	req := token.NewValidateTokenRequest(ak)
	tk, err := svc.ValidateToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 请求通过, 用户的身份信息 携带在 请求的上下文当中，传递给后续请求
	// 对于Gin, 如果把请求的中间数据传递下，使用了一个map, 这个map在request对象上
	c.Set(token.TOKEN_MIDDLEWARE_KEY, tk)

	// 后面请求如何获取 中间信息
	// c.Get(token.TOKEN_MIDDLEWARE_KEY).(*token.Token).UserName
}

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
