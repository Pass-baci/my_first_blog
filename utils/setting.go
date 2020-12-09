package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	JwtKey string

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string

	AccessKey string
	SecretKey string
	Bucket string
	QiniuServer string
)

func init() {
	file, err := ini.Load("../ginblog/config/config.ini")
	if err != nil {
		fmt.Println("读取配置错误，请重新配置路径：",err)
	}
	ServerLoad(file)
	DatabaseLoad(file)
	QiniuLoad(file)
}

func ServerLoad(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3030")
	JwtKey = file.Section("server").Key("JwtKey").MustString(":89sajs56")
}

func DatabaseLoad(file *ini.File){
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}

func QiniuLoad(file *ini.File) {
	AccessKey =file.Section("qiniu").Key("AccessKey").String()
	SecretKey =file.Section("qiniu").Key("SecretKey").String()
	Bucket =file.Section("qiniu").Key("Bucket").String()
	QiniuServer =file.Section("qiniu").Key("QiniuServer").String()
}