package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/model"
	"github.com/yurongjie2003/ginblog/service"
	"net/http"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var article model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetArticleService().AddArticle(&article)
	c.JSON(http.StatusOK, results.NewResult(&article, code))
}

// GetArticleDetail 获取文章详情
func GetArticleDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	article, code := service.GetArticleService().GetArticleDetail(id)
	c.JSON(http.StatusOK, results.NewResult(&article, code))
}

// SearchArticles 搜索文章列表
func SearchArticles(c *gin.Context) {
	pageParams := results.GetPageParams(c)
	pageResult, code := service.GetArticleService().SearchArticles(pageParams)
	c.JSON(http.StatusOK, results.NewResult(pageResult, code))
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var article *model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetArticleService().EditArticle(article)
	c.JSON(http.StatusOK, results.NewResult(nil, code))
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetArticleService().DeleteArticle(id)
	c.JSON(http.StatusOK, results.NewResult(nil, code))
}
