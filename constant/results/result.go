package results

import (
	"github.com/yurongjie2003/ginblog/constant/codes"
)

type Result struct {
	Code    codes.Code  `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult(data interface{}, code codes.Code) *Result {
	msg, _ := codes.GetMsgOfCode(code)
	return &Result{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func Error(code codes.Code) *Result {
	return NewResult(nil, code)
}

func Success(data interface{}) *Result {
	return NewResult(data, codes.Success)
}
