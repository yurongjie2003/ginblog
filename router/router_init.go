package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	v1 "github.com/yurongjie2003/ginblog/api/v1"
	"github.com/yurongjie2003/ginblog/middleware"
	"github.com/yurongjie2003/ginblog/utils/Config"
	"github.com/yurongjie2003/ginblog/utils/Log"
	"time"
)

func Init() error {
	gin.SetMode(Config.AppMode)
	r := gin.New()

	// 使用Zap替换gin默认Logger
	r.Use(ginzap.Ginzap(Log.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(Log.Logger, true))

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = Config.MaxFileSize

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtAuth())
	{
		// User模块路由接口
		auth.POST("/user/", v1.AddUser)
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)

		// Category模块路由接口
		auth.POST("/category/", v1.AddCategory)
		auth.PUT("/category/", v1.EditCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)

		// Article模块路由接口
		auth.POST("/article/", v1.AddArticle)
		auth.PUT("/article/", v1.EditArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)
		auth.POST("/article/cover/", v1.UploadCover)
	}

	noAuth := r.Group("api/v1")

	{
		noAuth.POST("/login", v1.Login)

		// User模块路由接口
		noAuth.GET("/user/:id", v1.GetUserDetail)
		noAuth.GET("/users", v1.GetUsers)
		noAuth.GET("/user/exist", v1.CheckUserExist)

		// Category模块路由接口
		noAuth.GET("/category/:id/articles", v1.GetCategoryArticles)
		noAuth.GET("/categories", v1.GetCategories)
		noAuth.GET("/category/exist", v1.CheckCategoryExist)

		// Article模块路由接口
		noAuth.GET("/article/:id", v1.GetArticleDetail)
	}
	err := r.Run(Config.HttpPort)
	return err
}
