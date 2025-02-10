package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/service"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}

	token, code := service.GetUserService().Login(username, password)
	c.JSON(http.StatusOK, results.NewResult(token, code))
}
