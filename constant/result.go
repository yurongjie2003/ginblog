package constant

import (
	"github.com/yurongjie2003/ginblog/constant/code"
)

type Result struct {
	Code    code.Code   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult(data interface{}, responseCode code.Code) *Result {
	msg, _ := code.GetMsgOfCode(responseCode)
	return &Result{
		Code:    responseCode,
		Message: msg,
		Data:    data,
	}
}
