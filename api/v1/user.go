package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/model"
	"github.com/yurongjie2003/ginblog/service"
	"net/http"
)

// CheckUserExist 查询用户名是否存在
func CheckUserExist(c *gin.Context) {
	exist, code := service.GetUserService().CheckUserExist(c.Query("username"))
	c.JSON(http.StatusOK, results.NewResult(&gin.H{
		"exist": exist,
	}, code))
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, results.Error(codes.ErrorArgs))
		return
	}
	code := service.GetUserService().AddUser(&user)
	user.Password = ""
	c.JSON(http.StatusOK, results.NewResult(user, code))
}

// GetUserDetail 查询用户详情
func GetUserDetail(c *gin.Context) {

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageParams := results.GetPageParams(c)
	pageResult, code := service.GetUserService().GetUsers(pageParams)
	c.JSON(http.StatusOK, results.NewResult(pageResult, code))
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}
