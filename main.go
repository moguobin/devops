package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moguobin/devops/router"
)

func main() {
	e := gin.Default()
	r := e.Group("v1")
	//初始化路由
	router.InitRouter(r)
	e.Run()
}
