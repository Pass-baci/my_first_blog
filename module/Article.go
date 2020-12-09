package module

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int not null" json:"cid"`
	Desc string `gorm:"type:varchar(100)" json:"desc"`
	Content string `gorm:"type:longtext not null" json:"content"`
	Img string `gorm:"type:varchar(100) not null" json:"img"`
}

//CreateArticle 添加文章
func CreateArticle(data *Article) (code int) {
	err := db.Preload("Category").Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//GetArticles 查询文章列表
func GetArticles(artname string, pageSize, pageNum int)(code int, article []Article, total int) {
	if artname == "" {
		db.Preload("Category").Order("Updated_At desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Count(&total)
		code = errmsg.SUCCESS
	} else {
		err := db.Preload("Category").Where("title Like ?", "%"+artname+"%").Order("Updated_At desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			code = errmsg.ERROR
			return code, nil, 0
		}
	}
	return code, article, total
}

//GetCateArticle 查询分类下的所有文章
func GetCateArticle(id, pageSize, pageNum int) (articles []Article, code int, total int) {
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Where("cid = ?", id).Find(&articles).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return articles, errmsg.SUCCESS, total
}

//GetArt 查询单个文章
func GetArt(id int) (data Article, code int) {
	err = db.Preload("Category").Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return data, errmsg.SUCCESS
}

//EditArticle 修改文章
func EditArticle(id int, data *Article) (code int) {
	var (
		maps = make(map[string]interface{})
		art Article
	)
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Preload("Category").Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DelArticle 删除文章
func DelArticle(id int) (code int) {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil || id == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
