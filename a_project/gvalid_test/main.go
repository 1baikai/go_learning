package main

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/util/gvalid"
	"github.com/sirupsen/logrus"
)

type DatabaseOperatingForm struct {
	HostPassword string `json:"hostPassword" v:"required#hostPassword不能为空"`
	Action       string `json:"action" v:"required#不满足规则"`
	SysUser      string `json:"sysUser" v:"required#sysUser不能为空"`
}

func (param DatabaseOperatingForm) Validate() error {
	if err := gvalid.CheckStruct(param, nil); err != nil {
		logrus.Errorf("PatchUsername err: %v", err.Error())
		return errors.New(err.Error())
	}
	return nil
}
type RestoreDBForm struct {
	BackupID     int64  `json:"backupID,string" v:"required#备份ID不能为空"`
	DbName       string `json:"dbName" v:"required#恢复数据库不能为空"`
	HostUser     string `json:"hostUser" v:"required#主机用户不能为空"`
	HostPassword string `json:"hostPassword" v:"required#主机密码不能为空"`
}

func (v RestoreDBForm) Validate() error {
	return gvalid.CheckStruct(v, nil)

	//if err := gvalid.CheckStruct(v, nil); err != nil {
	//	logrus.Errorf("RestoreDBForm err: %v", err.Error())
	//	return errors.New(err.Error())
	//}
	//return nil
}


func main() {
	Form:=DatabaseOperatingForm{
		HostPassword: "123",
		Action:       "stop",
		SysUser:      "123",
	}
	err := Form.Validate()
	fmt.Println(err == nil)

	c:=RestoreDBForm{
		BackupID:     1111111111111111,
		DbName:       "11",
		HostUser:     "dasdas",
		HostPassword: "ww",
	}
	err = c.Validate()
	fmt.Println(err==nil)
}