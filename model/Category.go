package model

import (
	"errors"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"gorm.io/gorm"
	"sync"
)

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

type CategoryDao struct {
}

var categoryDao *CategoryDao
var categoryDaoOnce sync.Once

func GetCategoryDao() *CategoryDao {
	categoryDaoOnce.Do(func() {
		categoryDao = &CategoryDao{}
	})
	return categoryDao
}

func (*CategoryDao) CheckCategoryExist(name string) (bool, codes.Code) {
	var count int64
	err := db.Model(&Category{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, codes.Error
	}
	return count > 0, codes.Success
}

func (*CategoryDao) AddCategory(category *Category) codes.Code {
	err := db.Create(category).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (d *CategoryDao) QueryCategories(params *results.PageParams) (*results.PageResult, codes.Code) {
	var categories []*Category
	var count int64
	err := db.Model(&Category{}).Count(&count).Limit(params.PageSize).Offset(params.Offset).Find(&categories).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, codes.Error
	}
	return &results.PageResult{
		Total:   count,
		Records: categories,
	}, codes.Success
}

func (d *CategoryDao) DeleteCategory(id int) codes.Code {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (d *CategoryDao) EditCategory(category *Category) codes.Code {
	err := db.Model(&Category{}).Where("id = ?", category.ID).Updates(category).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (*CategoryDao) CheckCategoryExistById(id int) (bool, codes.Code) {
	var count int64
	err := db.Model(&Category{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, codes.Error
	}
	return count > 0, codes.Success
}
