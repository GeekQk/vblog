package api

import (
	"github.com/GeekQk/vblog/apps/token"
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
	re := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(re); err != nil {
		return
	}

	//2.业务逻辑处理
	//3.返回处理结果

}

// 退出
func (h *TokenApiHandle) Logout(c *gin.Context) {

}
