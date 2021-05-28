package errors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCodeAndData(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []interface{}
	}{
		{
			name:  "取出code和data",
			args:  args{NewFmt(ErrorInvalidParams1, getMsg(ErrorInvalidParams1, "", SUCCESS), "邮箱格式不合法")},
			want:  10005,
			want1: []interface{}{"邮箱格式不合法"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCodeAndData(tt.args.err)
			fmt.Println(got, got1)
			if got != tt.want {
				t.Errorf("GetCodeAndData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetCodeAndData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewFromGrpcError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "测试",
			args:    args{New(ErrorInvalidParams1).GRPCStatus().Err()},
			wantErr: true,
		},
		{
			name:    "测试111",
			args:    args{NewFmt(ErrorInvalidParams1, GetErrMsg(ErrorInvalidParams1, ""), "asd")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewFromGrpcError(tt.args.err)
			errCode, errData := GetCodeAndData(err)
			fmt.Printf("err%v ,,%v", errCode, errData)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromGrpcError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
