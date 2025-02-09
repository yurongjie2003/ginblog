package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yurongjie2003/ginblog/api/v1"
	"github.com/yurongjie2003/ginblog/config"
)

func Init() error {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// User模块路由接口
		routerV1.POST("/user/", v1.AddUser)
		routerV1.GET("/user/:id", v1.GetUserDetail)
		routerV1.GET("/users", v1.GetUsers)
		routerV1.PUT("/user/:id", v1.EditUser)
		routerV1.DELETE("/user/:id", v1.DeleteUser)
		routerV1.GET("/user/exist", v1.CheckUserExist)

		// Category模块路由接口
		routerV1.POST("/category/", v1.AddCategory)
		routerV1.GET("/category/:id/articles", v1.GetCategoryArticles)
		routerV1.GET("/categories", v1.GetCategories)
		routerV1.PUT("/category/", v1.EditCategory)
		routerV1.DELETE("/category/:id", v1.DeleteCategory)
		routerV1.GET("/category/exist", v1.CheckCategoryExist)

		// Article模块路由接口
		routerV1.POST("/article/", v1.AddArticle)
		routerV1.GET("/article/:id", v1.GetArticleDetail)
		routerV1.GET("/articles", v1.SearchArticles)
		routerV1.PUT("/article/", v1.EditArticle)
		routerV1.DELETE("/article/:id", v1.DeleteArticle)
	}
	err := r.Run(config.HttpPort)
	return err
}
