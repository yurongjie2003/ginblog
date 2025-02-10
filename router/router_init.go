package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yurongjie2003/ginblog/api/v1"
	"github.com/yurongjie2003/ginblog/config"
	"github.com/yurongjie2003/ginblog/middleware"
)

func Init() error {
	gin.SetMode(config.AppMode)
	r := gin.Default()

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
	err := r.Run(config.HttpPort)
	return err
}
