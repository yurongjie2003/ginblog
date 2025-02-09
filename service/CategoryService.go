package service

import (
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/model"
	"sync"
)

type CategoryService struct {
}

var categoryService *CategoryService
var categoryServiceOnce sync.Once

func GetCategoryService() *CategoryService {
	categoryServiceOnce.Do(func() {
		categoryService = &CategoryService{}
	})
	return categoryService
}

func (*CategoryService) CheckCategoryExist(name string) (bool, codes.Code) {
	return model.GetCategoryDao().CheckCategoryExist(name)
}

func (*CategoryService) AddCategory(m *model.Category) codes.Code {
	exist, code := model.GetCategoryDao().CheckCategoryExist(m.Name)
	if code != codes.Success {
		return code
	}
	if exist {
		return codes.ErrorCategoryExist
	}
	return model.GetCategoryDao().AddCategory(m)
}

func (s *CategoryService) GetCategories(pageParams *results.PageParams) (*results.PageResult, codes.Code) {
	return model.GetCategoryDao().QueryCategories(pageParams)
}

func (s *CategoryService) DeleteCategory(id int) codes.Code {
	return model.GetCategoryDao().DeleteCategory(id)
}

func (s *CategoryService) EditCategory(category *model.Category) codes.Code {
	return model.GetCategoryDao().EditCategory(category)
}
