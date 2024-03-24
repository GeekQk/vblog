package cmd

import (
	"github.com/GeekQk/vblog/conf"
	"github.com/GeekQk/vblog/ioc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务器",
	Run: func(cmd *cobra.Command, args []string) {
		// 什么都不做的时候打印帮助信息
		// 1. 初始化程序配置, 这里没有配置，使用默认值
		cobra.CheckErr(conf.LoadFromEnv())

		// 2. 程序对象管理
		cobra.CheckErr(ioc.Init())

		// Protocol
		engine := gin.Default()
		engine.Use(cors.Default())

		rr := engine.Group("/vblog/api/v1")
		ioc.RegisteryGinApi(rr)

		//启动grpc服务
		go conf.C().GrpcServer.Start()

		// 把Http协议服务器启动起来
		cobra.CheckErr(engine.Run(":8080"))
	},
}
