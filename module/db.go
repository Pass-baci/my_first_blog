package module

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db *gorm.DB
	err error
)

//InitDataBase 初始化数据库
func InitDataBase() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName))
	if err != nil {
		fmt.Println("连接数据库失败，请重新配置：",err)
		return
	}
	// 禁用默认表名的复数形式
	db.SingularTable(true)

	//AutoMigrate 自动迁移 只会 创建表、缺失的列、缺失的索引
	db.AutoMigrate(&User{},&Category{},&Article{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//关闭数据库连接
	//db.Close()
}
