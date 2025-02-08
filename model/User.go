package model

import (
	"errors"
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

type UserVo struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Role     int    `gorm:"type:int" json:"role"`
}

type UserDao struct {
}

var userDao *UserDao
var userDaoOnce sync.Once

func GetUserDao() *UserDao {
	userDaoOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// CheckUsernameExist 检查用户名是否存在
func (*UserDao) CheckUsernameExist(username string) (exist bool, code codes.Code) {
	var count int64
	if err := db.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, codes.Error
	}
	return count > 0, codes.Success
}

// CreateUser 新增用户
func (*UserDao) CreateUser(user *User) codes.Code {
	err := db.Create(user).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

// QueryUsers 查询用户列表
func (*UserDao) QueryUsers(pageParams *results.PageParams) (*results.PageResult, codes.Code) {
	var users []*UserVo
	var total int64
	err := db.Model(&User{}).Count(&total).Limit(pageParams.PageSize).Offset(pageParams.Offset).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, codes.Error
	}
	return &results.PageResult{
		Total:   total,
		Records: users,
	}, codes.Success
}

// DeleteUserById 删除用户
func (*UserDao) DeleteUserById(id int) codes.Code {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

// EditUser 编辑用户
func (*UserDao) EditUser(id int, user *User) codes.Code {
	user.ID = uint(id)
	err := db.Model(&user).Select("role").Updates(user).Error
	if err != nil {
		return codes.Error
	}
	return codes.Success
}

func (*UserDao) GetUserDetail(id int) (UserVo, codes.Code) {
	var user UserVo
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, codes.Error
	}
	return user, codes.Success
}
