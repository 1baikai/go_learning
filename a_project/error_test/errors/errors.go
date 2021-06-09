package errors

import (
	"fmt"
	"strings"

	"baikai/leetcode/a_project/error_test/i18n"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// LangZh 中文
	LangZh = "zh"
	// LangEn 英文
	LangEn = "en"
	// LangJp 日文
	LangJp = "jp"
)

type LocalError struct {
	code    int
	message string
	data    []interface{}
}

type errorsInterface interface {
	Code() int
	Error() string
	Data() []interface{}
}

func (le *LocalError) Code() int {
	return le.code
}

func (le *LocalError) Error() string {
	return le.message
}

// GRPCStatus 转换 LocalError 类型为 grpc 的 Status 类型
func (le *LocalError) GRPCStatus() *status.Status {
	return status.New(codes.Code(le.code), le.message)
}

// NewFromGrpcError 从将微服务的错误本地化
func NewFromGrpcError(err error) error {
	if err == nil {
		return nil
	}
	if rpcErr, ok := status.FromError(err); ok {
		code := int(rpcErr.Code())
		fmt.Println("code", code)
		if rpcErr.Code() == codes.DeadlineExceeded {
			return NewS(ErrorConnectTimeout, GetSuccessMsg(ErrorConnectTimeout, ""))
		}
		if code > 100 {
			return NewS(code, rpcErr.Message())
		}
		if strings.Contains(err.Error(), "pq") {
			return NewFmt(ErrorDBInfo, GetErrMsg(ErrorDBInfo, ""), err.Error())
		}
		if strings.Contains(err.Error(), "ssh:") {
			return NewFmt(ErrorSsh, GetErrMsg(ErrorSsh, ""), err.Error())
		}
		if strings.Contains(err.Error(), "gs_om: command not found") {
			return NewFmt(ErrorGsOm, GetErrMsg(ErrorGsOm, ""), err.Error())
		}
		if strings.Contains(err.Error(), "i/o timeout") {
			return NewFmt(ErrorTestContent, GetErrMsg(ErrorTestContent, ""), err.Error())
		}
		return NewS(code, rpcErr.Message())
	}
	return err
}

func New(code int) *LocalError {
	lang := i18n.LangZh
	return &LocalError{
		code:    code,
		message: codeToStr[code][lang],
	}
}

func NewS(code int, message string) *LocalError {
	return &LocalError{
		code:    code,
		message: message,
	}
}

func NewF(code int, format string, a ...interface{}) *LocalError {
	return NewS(code, fmt.Sprintf(format, a...))
}
func NewFmt(code int, format string, a ...interface{}) *LocalError {
	return &LocalError{
		code:    code,
		message: format,
		data:    a,
	}
}

//
//// Covert 将一个未知的错误类型转为 LocalError 类型
//func Covert(err error) *LocalError {
//	if e, ok := err.(interface {
//		Code() ErrCode
//		Error() string
//	}); ok {
//		return &LocalError{
//			code:    e.Code(),
//			message: e.Error(),
//		}
//	}
//	return NewS(ErrorOther, err.Error())
//}

// GetCodeAndData Get Code And Data
func GetCodeAndData(err error) (int, []interface{}) {
	if c, ok := err.(errorsInterface); ok {
		return c.Code(), c.Data()
	}
	return 0, nil
}

// Data error data
func (le *LocalError) Data() []interface{} {
	return le.data
}

// GetErrMsg Get Err Msg
func GetErrMsg(code int, lang i18n.Language) string {
	return getMsg(code, lang, ERROR)
}

// GetSuccessMsg Get Success Msg
func GetSuccessMsg(code int, lang i18n.Language) string {
	return getMsg(code, lang, SUCCESS)
}

// getMsg get error information based on Code
func getMsg(code int, lang i18n.Language, status int) string {
	var msg string

	if lang == "" {
		lang = i18n.LangZh
	}
	errorMsg, ok := codeToStr[code][lang]
	if ok {
		msg = errorMsg
	} else {
		msg = codeToStr[status][lang]
	}

	return msg
}

//func ToErrCode(i int) ErrCode {
//	return ErrCode(i)
//}
