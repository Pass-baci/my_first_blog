package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,                                                        //配置允许域名的请求
		AllowMethods:     []string{"*"},                                               //允许请求的方法
		AllowHeaders:     []string{"*"},                                               //是允许客户端与跨域请求一起使用的非简单标头的列表
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"}, //哪些标头可以安全地显示给CORS的API
		AllowCredentials: true,                                                        //是否允许携带cookie的请求
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	})
}
