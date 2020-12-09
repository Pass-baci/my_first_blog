package errmsg


//定义一些状态码，用于处理各个环节的错误
const (
	SUCCESS = 200
	ERROR = 500

	//code = 1000...用户模块的错误状态码
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NOT_ROLE    = 1008

	//code = 2000... 分类模块的错误状态码
	ERROR_CATAGORY_USED      = 2001
	ERROR_CATEGORY_NOT_EXIST = 2002
	//code = 3000... 文章模块的错误状态码
	ERROR_ARTICLE_NOT_EXIST = 3001
)

var codemsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户名不存在",
	ERROR_USER_NOT_ROLE:      "该用户没有权限",
	ERROR_TOKEN_EXIST:        "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:      "TOKEN已过期",
	ERROR_TOKEN_WRONG:        "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误",
	ERROR_CATAGORY_USED:      "分类名已存在",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
	ERROR_ARTICLE_NOT_EXIST:  "文章不存在",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}