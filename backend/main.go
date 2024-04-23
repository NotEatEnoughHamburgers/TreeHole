package main

import (
	"TreeHole/controller"
	"TreeHole/dao"
	"TreeHole/middle"
	"github.com/gin-gonic/gin"
)

func main() {
	// 连接
	dao.Client()
	// 初始化表
	dao.AutoTables()
	// 创建gin
	r := gin.Default()
	// 加载中间件
	r.Use(middle.CORS())
	r.Use(middle.JWTAuth())
	// 加载路由
	controller.Router.InitApiRouter(r)
	// 运行WebService
	r.Run(":8081")
}