package v1

import (
	"ginblog/middleware"
	"ginblog/module"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var (
		token string
		data module.User
	)
	_ = c.ShouldBindJSON(&data)
	code = module.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"token":token,
	})
}
