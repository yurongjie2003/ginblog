package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yurongjie2003/ginblog/constant"
	"github.com/yurongjie2003/ginblog/constant/code"
	"net/http"
)

// CheckUserExist 查询用户名是否存在
func CheckUserExist(c *gin.Context) {
	c.JSON(http.StatusOK, constant.NewResult(false, code.Success))
}

// AddUser 添加用户
func AddUser(c *gin.Context) {

}

// GetUserDetail 查询用户详情
func GetUserDetail(c *gin.Context) {

}

// GetUsers 搜索用户列表
func GetUsers(c *gin.Context) {

}

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}
