package main

import (
	"net/http"

	_err "baikai/leetcode/a_project/error_test/errors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.POST("/ping", TestErr)
	r.Run()
}

type UserInfo struct {
	Name string `json:"name" v:"required|in:a,b#不能为空|只能是a或者b"`
}

func (param UserInfo) Validate() error {
	if err := gvalid.CheckStruct(param, nil); err != nil {
		logrus.Errorf("PatchUsername err: %v", err.Error())
		return _err.NewFmt(_err.ErrorInvalidParams1, _err.GetErrMsg(_err.ErrorInvalidParams1, ""), err.Error())
	}

	return nil
}

func TestErr(c *gin.Context) {
	//_ := c.Request.Context()
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo)
	if err != nil {
		logrus.Errorf("ShouldBind %v", err)
		HTTPRequestFailedV4(c, _err.New(_err.ErrorInvalidParams), http.StatusBadRequest)
		return
	}
	if err = userInfo.Validate(); err != nil {
		HTTPRequestFailedV4(c, err, http.StatusBadRequest)
		return
	}
	HTTPRequestSuccess(c, http.StatusOK, userInfo.Name)
}
