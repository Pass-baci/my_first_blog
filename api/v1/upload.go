package v1

import (
	"ginblog/module"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpLoadFile(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := module.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"url": url,
		"message": errmsg.GetErrMsg(code),
	})
}
