package module

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string`gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=10" label:"用户名"`
	Password string`gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role int`gorm:"type:int; DEFAULT:2" json:"role" validate:"lt=3" label:"角色码"`
}

//验证用户是否存在
func CheckUser(name string) (code int) {
	var data User
	db.Select("id").Where("username = ?",name).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//更新查询
func CheckUpUser(id int, name string) (code int) {
	var data User
	db.Select("id").Where("username = ?",name).First(&data)
	if data.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if data.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//密码加密
func encryPw(pw string) string {
	const KeyLen = 10
	salt := make([]byte,5)
	salt = []byte{12,66,78,98,122}
	hashPw, err := scrypt.Key([]byte(pw), salt, 1<<5, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	toString := base64.StdEncoding.EncodeToString(hashPw)
	return toString
}

//登录接口
func CheckLogin(username, password string) (code int) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if encryPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_ROLE
	}
	return errmsg.SUCCESS
}

//添加用户
func CreateUser(data *User) (code int) {
	data.Password = encryPw(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个用户
func GetUser(id int) (User, int)  {
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

//查询用户列表
func GetUsers(username string, pageSize, pageNum int)(code int, users []User, total int) {
	if username == "" {
		db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)
		code = errmsg.SUCCESS
		return code, users, total
	}
	db.Where("username Like ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)
	if err == gorm.ErrRecordNotFound {
		code = errmsg.ERROR
		return code, nil, 0
	}
	code = errmsg.SUCCESS
	return code, users, total
}

//修改用户信息
func EditUser(id int, data *User) (code int) {
	var (
		maps = make(map[string]interface{})
		user User
	)
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DelUser(id int) (code int) {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil || id == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
