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

// CheckCategoryExist 查询分类名是否存在
func CheckCategoryExist(c *gin.Context) {
	name := c.Query("name")
	exist, code := service.GetCategoryService().CheckCategoryExist(name)
	c.JSON(http.StatusOK, results.NewResult(&gin.H{
		"exist": exist,
	}, code))
}

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetCategoryService().AddCategory(&category)
	c.JSON(http.StatusOK, results.NewResult(&category, code))
}

// GetCategoryArticles 获取单个分类下的文章
func GetCategoryArticles(c *gin.Context) {

}

// GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
	pageParams := results.GetPageParams(c)
	pageResult, code := service.GetCategoryService().GetCategories(pageParams)
	c.JSON(http.StatusOK, results.NewResult(pageResult, code))
}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var category *model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetCategoryService().EditCategory(category)
	c.JSON(http.StatusOK, results.NewResult(nil, code))
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetCategoryService().DeleteCategory(id)
	c.JSON(http.StatusOK, results.NewResult(nil, code))
}
