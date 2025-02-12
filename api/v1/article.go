package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/model"
	"github.com/yurongjie2003/ginblog/service"
	"github.com/yurongjie2003/ginblog/utils/Config"
	"github.com/yurongjie2003/ginblog/utils/Minio"
	"mime/multipart"
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

// UploadCover 上传封面
func UploadCover(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrNoFileUploaded))
		return
	}

	// 验证文件类型
	if file.Header.Get("Content-Type") != "image/jpeg" {
		c.JSON(http.StatusOK, results.Error(codes.ErrFileTypeNotSupported))
		return
	}

	// 验证文件大小
	if file.Size > Config.MaxFileSize {
		c.JSON(http.StatusOK, results.Error(codes.ErrFileSizeExceedsLimit))
		return
	}

	// 处理文件名，确保安全性
	filename := fmt.Sprintf("%s.jpg", uuid.New().String())

	// 打开文件以获取 io.Reader
	fileReader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.Error))
		return
	}
	defer func(fileReader multipart.File) {
		err := fileReader.Close()
		if err != nil {
			c.JSON(http.StatusOK, results.Error(codes.ErrFileUpload))
		}
	}(fileReader)

	// 上传文件到 Minio
	err = Minio.Client.UploadPublicFile(fileReader, filename, file.Size, "image/jpeg")
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrFileUpload))
		return
	}

	url := Minio.Client.GetPublicFileURL(filename)

	// 返回成功响应
	c.JSON(http.StatusOK, results.NewResult(url, codes.Success))
}
