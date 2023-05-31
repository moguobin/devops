package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moguobin/devops/controller/BaseController"
)

func InitRouter(r *gin.RouterGroup) {
	base := r.Group("/base")
	{
		base.GET("/health", BaseController.Health)
	}
}
