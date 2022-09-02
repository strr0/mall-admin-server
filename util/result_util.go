package util

import (
	"net/http"
)

type CommonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CommonPage struct {
	PageNum   int         `json:"pageNum"`
	PageSize  int         `json:"pageSize"`
	TotalPage int         `json:"totalPage"`
	Total     int64       `json:"total"`
	List      interface{} `json:"list"`
}

func Data(data interface{}) CommonResult {
	return CommonResult{
		Code: http.StatusOK,
		Data: data,
	}
}

func Page(list interface{}, total int64) CommonResult {
	return Data(CommonPage{
		List:  list,
		Total: total,
	})
}

func Success(message string) CommonResult {
	return CommonResult{
		Code:    http.StatusOK,
		Message: message,
	}
}

func Failed(message string) CommonResult {
	return CommonResult{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
