package v1

import (
	"ginblog/module"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data module.Article
	_ = c.ShouldBindJSON(&data)
	//if err != nil {
	//	code = errmsg.ERROR
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": errmsg.GetErrMsg(code),
	//	})
	//}
	if data.Title == "" {
		code = errmsg.ERROR
	}
	code = module.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//SelectArticle 查询文章列表
func SelectArticle(c *gin.Context) {
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	artname := c.Query("title")
	switch {
	case pageSize >= 100:
			pageSize = 100
	case pageSize <= 0:
			pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := module.GetArticles(artname,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
		"total":total,
	})
}

//SelectCateArt 查询分类下的所有文章
func SelectCateArt(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	switch {
	case pageSize >= 100:
			pageSize = 100
	case pageSize <= 0:
			pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := module.GetCateArticle(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
		"total":total,
	})
}

//GetArtInfo 查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := module.GetArt(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//UpdateArticle 修改文章
func UpdateArticle(c *gin.Context) {
	var data module.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = module.EditArticle(id, &data)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}

//DeleteArticle 删除文章
func DeleteArticle(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	code = module.DelArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
