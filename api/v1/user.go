package v1

import (
	"ginblog/module"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var data module.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"message": msg,
		})
		return
	}
	//if err != nil {
	//	code = errmsg.ERROR
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": errmsg.GetErrMsg(code),
	//	})
	//}
	code = module.CheckUser(data.Username)
	if data.Username == "" || data.Password == "" {
		code = errmsg.ERROR
	}
	if code == errmsg.SUCCESS {
		module.CreateUser(&data)
	}
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
//查询单个用户
func GetUserInfo(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := module.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"total": 1,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func SelectUser(c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	if pageSize == 0 {
		pageSize =-1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	code, data, total := module.GetUsers(username, pageSize, pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
		"total":total,
	})
}

//修改用户信息
func UpdateUser(c *gin.Context) {
	var data module.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = module.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		code = module.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	code = module.DelUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
