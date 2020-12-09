package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	rotatelog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"math"
	"os"
	"time"
)


func Logger() gin.HandlerFunc {
	filePath := "log/log"	//日志路径
	linkName := "latest_log.log"	//软连接名称
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:",err)
	}
	logger := logrus.New()	//实例化logrus
	logger.SetLevel(logrus.DebugLevel)	//设置日志级别为Debug
	logger.Out = file //日志的输出方式
	logs, _ := rotatelog.New(
		filePath+"%Y%m%d.log",	//文件名称
		rotatelog.WithMaxAge(7*24*time.Hour),	//日志保留时长
		rotatelog.WithRotationTime(24*time.Hour),	//日志分割时间间隔
		rotatelog.WithLinkName(linkName),	//软连接到最新日志
	)
	//将日志级别映射到io.Writer
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logs,
		logrus.FatalLevel: logs,
		logrus.DebugLevel: logs,
		logrus.WarnLevel:  logs,
		logrus.ErrorLevel: logs,
		logrus.PanicLevel: logs,
	}
	//实例化Hook，在输出前进行时间格式化
	hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//将Hook添加到记录器中
	logger.AddHook(hook)
	return func(c *gin.Context) {
		startime := time.Now() //获取开始执行的时间
		c.Next()
		stoptime := time.Since(startime) //执行时间的时间戳
		spendtime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stoptime)) / 1000000.0)) //格式化输出
		hostName, err := os.Hostname() //获取主机用户名
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status() //获取当前状态码
		clientIp := c.ClientIP() //获取客户端Ip
		userAgent := c.Request.UserAgent()	//获取客户端访问的软件
		dataSize := c.Writer.Size()	//获取客户端请求的字节大小
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method	//获取客户端请求的方法
		path := c.Request.RequestURI	//获取客户端请求的Path
		//将需要打印的日志条目添加到当中
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Ip":        clientIp,
			"Spendtime": spendtime,
			"Status":    statusCode,
			"Agent":     userAgent,
			"DataSize":  dataSize,
			"Method":    method,
			"Path":      path,
		})
		//判断错误
		//1.gin内部错误
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		//2.自定义的状态码所显示的错误
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
