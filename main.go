package main

import (
	// 通过import方法 完成注册
	_ "github.com/GeekQk/vblog/apps"
	"github.com/GeekQk/vblog/cmd"
)

func main() {
	/*
		// 1. 初始化程序配置, 这里没有配置，使用默认值
		if err := conf.LoadFromEnv(); err != nil {
			panic(err)
		}

		// 2.2 程序对象管理
		if err := ioc.Init(); err != nil {
			panic(err)
		}
		// Protocol
		engine := gin.Default()

		rr := engine.Group("/vblog/api/v1")

		// ioc 能不能帮忙把 模块API的注册也一起管理
		// ioc.Init() , 对象初始化完成后, 如果对象是 api对象，就帮忙完成下注册
		// ioc.RegistryGin(rr)?
		//ioc.Api().Get(token.AppName).(*api.TokenApiHandle).Registery(rr)

		ioc.RegisteryGinApi(rr)
		// 把Http协议服务器启动起来
		if err := engine.Run(":8010"); err != nil {
			panic(err)
		}
	*/

	//使用cobra接管启动命令
	cmd.Execute()

}
