package user

import (
	"github.com/cloud-arrow/task-server/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserVo struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func PasswordLogin(c *gin.Context) {
	var uv UserVo
	if err := c.ShouldBindJSON(&uv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDto := user.UserDto{uv.Phone, uv.Password}
	user.Login(&userDto)
}
