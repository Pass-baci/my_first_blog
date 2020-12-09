package routes

import (
	"ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.LoadHTMLGlob("static/index.html")
	r.Static("/static","./static/static")
	r.StaticFile("favicon.ico","./static/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.DELETE("users/:id",v1.DeleteUser)	//删除用户
		auth.PUT("users/:id",v1.UpdateUser)	//修改用户
		auth.POST("category/add",v1.AddCategory)	//新增分类
		auth.DELETE("category/:id",v1.DeleteCategory)	//删除分类
		auth.PUT("category/:id",v1.UpdateCategory)	//修改分类
		auth.POST("article/add",v1.AddArticle)	//新增文章
		auth.DELETE("article/:id",v1.DeleteArticle)	//删除文章
		auth.PUT("article/:id",v1.UpdateArticle)	//修改文章
		auth.POST("upload",v1.UpLoadFile)	//上传图片
	}
	routev1 := r.Group("api/v1")
	{
		//用户路由接口
		routev1.POST("users/add",v1.AddUser) //新增用户
		routev1.POST("login",v1.Login)	//登录
		routev1.GET("users",v1.SelectUser)	//查询用户列表
		routev1.GET("users/:id",v1.GetUserInfo)//查询单个用户
		//分类路由接口
		routev1.GET("category",v1.SelectCategory)	//查询分类
		routev1.GET("category/:id",v1.GetCate)
		//文章路由接口
		routev1.GET("article",v1.SelectArticle)	//查询所有文章
		routev1.GET("article/List/:id",v1.SelectCateArt)	//查询分类下的所有文章
		routev1.GET("article/info/:id",v1.GetArtInfo)	//查询单个文章
	}
	r.Run(utils.HttpPort)
}
