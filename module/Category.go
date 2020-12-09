package module

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//Category 定义一个分类
type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20) not null" json:"name"`
}

//CheckCategory 验证分类是否存在
func CheckCategory(name string) (code int) {
	var data Category
	db.Select("id").Where("name = ?",name).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_CATAGORY_USED
	}
	return errmsg.SUCCESS
}

//CreateCategory 添加分类
func CreateCategory(data *Category) (code int) {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//查询单个分类
func GetCate(id int) (category Category, code int) {
	err = db.Where("id = ?", id).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		code = errmsg.ERROR
		return category, code
	}
	code = errmsg.SUCCESS
	return category, code
}

//GetCategories 查询分类列表
func GetCategories(pageSize, pageNum int, catename string)(code int, category []Category, total int) {
	err := db.Where("name Like ?", "%"+catename+"%").Order("ID").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&category).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		code = errmsg.ERROR
		return code, nil, 0
	}
	code = errmsg.SUCCESS
	return code, category, total
}

//EditCategory 修改分类名称
func EditCategory(id int, data *Category) (code int) {
	var (
		maps = make(map[string]interface{})
		cate Category
	)
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DelCategory 删除分类
func DelCategory(id int) (code int) {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil || id == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
