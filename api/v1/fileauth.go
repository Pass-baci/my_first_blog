package v1

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Fileauth(c *gin.Context) {
	file, err := os.Open("../fileauth.txt")
	if err != nil {
		return
	}
	defer file.Close()
}
