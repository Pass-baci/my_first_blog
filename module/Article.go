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
func GetArticles(artname string, pageSize, pageNum int)(article []Article, code int, total int) {
	var articleList []Article
	var err error

	if artname == "" {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&articleList).Error
		// 单独计数
		db.Model(&articleList).Count(&total)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		return articleList, errmsg.SUCCESS, total
	}
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?",
		"%"+artname+"%",
	).Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Where("title LIKE ?", "%"+artname+"%").Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

//GetCateArticle 查询分类下的所有文章
func GetCateArticle(id, pageSize, pageNum int) (articles []Article, code int, total int) {
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Where("cid = ?", id).Find(&articles).Error
	// 单独计数
	db.Model(&articles).Where("cid = ?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
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
