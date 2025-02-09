package service

import (
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/model"
	"sync"
)

type ArticleService struct {
}

func (*ArticleService) AddArticle(article *model.Article) codes.Code {
	cid := article.Cid
	categoryExist, code := model.GetCategoryDao().CheckCategoryExistById(cid)
	if code != codes.Success {
		return code
	}
	if !categoryExist {
		return codes.ErrorCategoryNotExist
	}
	return model.GetArticleDao().AddArticle(article)
}

func (*ArticleService) EditArticle(article *model.Article) codes.Code {
	cid := article.Cid
	categoryExist, code := model.GetCategoryDao().CheckCategoryExistById(cid)
	if code != codes.Success {
		return code
	}
	if !categoryExist {
		return codes.ErrorCategoryNotExist
	}
	return model.GetArticleDao().EditArticle(article)
}

func (*ArticleService) DeleteArticle(id int) codes.Code {
	return model.GetArticleDao().DeleteArticle(id)
}

var articleService *ArticleService
var articleServiceOnce sync.Once

func GetArticleService() *ArticleService {
	articleServiceOnce.Do(func() {
		articleService = &ArticleService{}
	})
	return articleService
}
