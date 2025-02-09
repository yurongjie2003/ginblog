package model

import (
	"errors"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"gorm.io/gorm"
	"sync"
)

type Article struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int;not null" json:"cid"`
	Category    Category `gorm:"foreignKey:Cid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Img         string   `gorm:"type:varchar(100)" json:"img"`
}

type ArticleDao struct {
}

func (*ArticleDao) AddArticle(article *Article) codes.Code {
	err := db.Create(article).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (*ArticleDao) EditArticle(article *Article) codes.Code {
	err := db.Model(article).Where("id = ?", article.ID).Updates(article).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (*ArticleDao) DeleteArticle(id int) codes.Code {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (*ArticleDao) GetArticleDetail(id int) (Article, codes.Code) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, codes.Error
	}
	return article, codes.Success
}

func (d *ArticleDao) SearchArticles(params *results.PageParams) (*results.PageResult, codes.Code) {
	var articles []*Article
	var count int64
	err := db.Preload("Category").Model(&Article{}).Count(&count).Limit(params.PageSize).Offset(params.Offset).Find(&articles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, codes.Error
	}
	return &results.PageResult{
		Total:   count,
		Records: articles,
	}, codes.Success
}

func (d *ArticleDao) GetCategoryArticles(cid int, params *results.PageParams) (*results.PageResult, codes.Code) {
	var articles []*Article
	var count int64
	err := db.Preload("Category").Model(&Article{}).Where("cid = ?", cid).Count(&count).Limit(params.PageSize).Offset(params.Offset).Find(&articles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, codes.Error
	}
	return &results.PageResult{
		Total:   count,
		Records: articles,
	}, codes.Success
}

var articleDao *ArticleDao
var articleDaoOnce sync.Once

func GetArticleDao() *ArticleDao {
	articleDaoOnce.Do(func() {
		articleDao = &ArticleDao{}
	})
	return articleDao
}
