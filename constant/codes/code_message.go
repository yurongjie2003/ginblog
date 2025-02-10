package codes

import "errors"

type Code int

const (
	Success   Code = 200
	Error     Code = 500
	ErrorArgs Code = 301
)

// codes 1000... 用户模块
const (
	ErrorUsernameUsed      Code = 1001
	ErrorUserPasswordWrong Code = 1002
	ErrorUserNotExist      Code = 1003
	ErrorTokenNotExist     Code = 1004
	ErrorTokenExpired      Code = 1005
	ErrorTokenWrong        Code = 1006
	ErrorTokenFormatWrong  Code = 1007
	ErrorTokenNeed         Code = 1008
)

// codes 2000... 分类模块
const (
	ErrorCategoryNotExist Code = 2001
	ErrorCategoryExist    Code = 2002
)

// codes 3000... 文章模块

var codeToMsg = map[Code]string{
	Success:                "操作成功",
	Error:                  "服务器内部错误",
	ErrorArgs:              "参数错误",
	ErrorUsernameUsed:      "用户名已被使用",
	ErrorUserPasswordWrong: "密码错误",
	ErrorUserNotExist:      "用户不存在",
	ErrorTokenNotExist:     "Token不存在",
	ErrorTokenExpired:      "Token已过期",
	ErrorTokenWrong:        "Token无效",
	ErrorTokenFormatWrong:  "Token格式错误",

	ErrorCategoryNotExist: "分类不存在",
	ErrorCategoryExist:    "分类已存在",

	ErrorTokenNeed: "请求需要Token",
}

func GetMsgOfCode(code Code) (string, error) {
	s, exist := codeToMsg[code]
	if !exist {
		return "无效错误码，请反馈开发者", errors.New("无效错误码，请反馈开发者")
	}
	return s, nil
}
