package service

import (
	"github.com/yurongjie2003/ginblog/constant/codes"
	"github.com/yurongjie2003/ginblog/constant/results"
	"github.com/yurongjie2003/ginblog/model"
	"sync"
)

type UserService struct {
}

var userService *UserService
var userServiceOnce sync.Once

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (*UserService) CheckUserExist(username string) (bool, codes.Code) {
	return model.GetUserDao().CheckUsernameExist(username)
}

func (*UserService) AddUser(user *model.User) codes.Code {
	exist, code := model.GetUserDao().CheckUsernameExist(user.Username)
	if code != codes.Success {
		return code
	}
	if exist {
		return codes.ErrorUsernameUsed
	}
	code = model.GetUserDao().CreateUser(user)
	return code
}

func (*UserService) GetUsers(pageParams *results.PageParams) (*results.PageResult, codes.Code) {
	return model.GetUserDao().QueryUsers(pageParams)
}
