package router

import (
	"github.com/cloud-arrow/task-server/controller/user"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.POST("/PasswordLogin", user.PasswordLogin)
}
