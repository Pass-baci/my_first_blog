package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var (
	code int
	JwtKey = []byte(utils.JwtKey)
)

type Myclaims struct {
	UserName string
	jwt.StandardClaims
}


//生成token
func SetToken(username string) (string, int) {
	expirTime := time.Now().Add(10 * time.Hour) //定义一个Token的超时时间
	setClaims := Myclaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			Issuer: "ginblog",
		},
	}
	//使用自定义的方式创建Token
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)	//所使用的的算法
	token, err := reqClaim.SignedString(JwtKey)	//创建Token
	if err != nil {
		return "", errmsg.ERROR
	}
	return token,errmsg.SUCCESS
}

//验证token
func CheckToken(token string) (*Myclaims, int) {
	var claims Myclaims
	settoken, _  := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := settoken.Claims.(*Myclaims); settoken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checktoken := strings.SplitN(tokenHeader, " ", 2)
		if len(checktoken) != 2 && checktoken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, Tcode := CheckToken(checktoken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.UserName)
		c.Next()
	}
}