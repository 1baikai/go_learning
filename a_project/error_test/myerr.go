package main

import (
	"fmt"
	"net/http"
	"strings"

	le "baikai/leetcode/a_project/8/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// HTTPRequestFailedV4 HTTP Request Failed V4  //todo 修改名字
func HTTPRequestFailedV4(c *gin.Context, err error, code int, data ...interface{}) {
	//err = le.NewFromGrpcError(err)

	var errCode int
	var errData []interface{}
	errCode, errData = le.GetCodeAndData(err)
	logrus.Infof("errCode is %v,errData is %v", errCode, errData)
	if errCode == 0 {
		errCode = code
		errData = data
	}
	if errCode == 0 {
		errCode = 500
	}
	if errData == nil {
		errData = data
	}

	msg := le.GetErrMsg(errCode, le.LangZh)
	if len(errData) > 0 && (errCode != 500 && errCode != 400) && strings.Contains(msg, "%") {
		msg = fmt.Sprintf(msg, errData...)
		msg = strings.TrimSpace(msg)
	}
	c.JSON(http.StatusOK, httpResponse(errCode, msg, ""))
}

type HTTPResponse struct {
	Code    int
	Message string
	Data    interface{}
}

// httpResponse return a http response
func httpResponse(code int, msg string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func HTTPRequestSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, httpResponse( code, "ssss", data))
}
