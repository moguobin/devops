package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct{}

func (b *BaseRouter) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "devops is ok")
}
