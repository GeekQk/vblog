package main

import (
	"github.com/GeekQk/vblog/apps/token/api"
	token_impl "github.com/GeekQk/vblog/apps/token/impl"
	user_impl "github.com/GeekQk/vblog/apps/user/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	//user service impl
	usvc := user_impl.NewUserServiceImpl()

	//tk service impl
	tsvs := token_impl.NewTokenServiceImpl(usvc)

	//api
	tokenApiHandler := api.NewTokenApiHandle(tsvs)
	//程序启动
	engine := gin.Default()
	rr := engine.Group("/vblog/api/v1")
	tokenApiHandler.Registery(rr)
	if err := engine.Run(":8090"); err != nil {
		panic(err)
	}

}
