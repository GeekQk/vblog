package api

import (
	"github.com/GeekQk/vblog/apps/token"
	"github.com/GeekQk/vblog/conf"
	"github.com/GeekQk/vblog/response"
	"github.com/gin-gonic/gin"
)

// 来实现对外的Restful接口
type TokenApiHandle struct {
	svc token.Service
}

// 如何处理路由 把路由注册给http Server
func (h *TokenApiHandle) Registery(rr gin.IRouter) {
	//每个业务模块 都需要gin Engine对象添加注册路由 单独放到外面
	// r := gin.Default()
	// rr := r.Group("vblog/api/v1")

	//模块路径 /vblog/api/v1/tokens
	//restful风格Api 本身就是操作资源 所以只需要一级tokens就可以了
	mr := rr.Group(token.AppName)
	mr.POST("/", h.Login)
	mr.DELETE("/", h.Logout)
}

// 登录
func (h *TokenApiHandle) Login(c *gin.Context) {
	//1.解析用户请求  http请求参数在body里面
	//原始写法：
	// io.ReadAll(c.Request.Body)
	// defer c.Request.Body.Close()
	//body必须是json,把json解析成结构体
	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}
	//2.业务逻辑处理
	tk, err := h.svc.IssueToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
	}
	//3.set Cookie
	c.SetCookie(token.TOKEN_COOKIE_KEY, tk.AccessToken, tk.AccessTokenExpiredAt, "/", conf.C().Application.Domain, false, true)

	//4.返回处理结果
	response.Success(c, tk)

}

// 退出
func (h *TokenApiHandle) Logout(c *gin.Context) {
	//1.解析用户请求  http请求参数在body里面
	//为了安全 token存放在Header的Cookie里面
	acessToken := token.GetTokenFromHttpHeader(c.Request)
	req := token.NewRevokeTokenRequest(acessToken, c.Query("refresh_token"))
	//2.业务逻辑处理
	_, err := h.svc.RevokeToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
	}
	//3.删除Cookie
	c.SetCookie(token.TOKEN_COOKIE_KEY, "", -1, "/", conf.C().Application.Domain, false, true)

	//4.返回处理结果
	response.Success(c, "退出成功")

}
