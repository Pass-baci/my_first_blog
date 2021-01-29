package v1

import (
	"ginblog/module"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data module.Category
	_ = c.ShouldBindJSON(&data)
	//if err != nil {
	//	code = errmsg.ERROR
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": errmsg.GetErrMsg(code),
	//	})
	//}
	code = module.CheckCategory(data.Name)
	if data.Name == "" {
		code = errmsg.ERROR
	}
	if code == errmsg.SUCCESS {
		module.CreateCategory(&data)
	}
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
//查询单个分类
func GetCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := module.GetCate(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//SelectCategory 查询分类列表
func SelectCategory(c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	catename := c.Query("catename")
	switch {
	case pageSize >= 100:
			pageSize = 100
	case pageSize <= 0:
			pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	code, data, total := module.GetCategories(pageSize,pageNum, catename)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
		"total":total,
	})
}

//UpdateCategory 修改分类名称
func UpdateCategory(c *gin.Context) {
	var data module.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = module.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		code = module.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}

//DeleteCategory 删除分类信息
func DeleteCategory(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	code = module.DelCategory(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
