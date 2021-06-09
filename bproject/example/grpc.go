package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"

	"bproject/pkg/grpc/pb"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var ceshiClient *SvcClient
var ctx = context.Background()

type SvcClient struct {
	test pb.BProjectClient
}

var run = flag.String("run", "", "执行函数名称")

// NewBasicUsersService 返回一个简单的、无状态的用户服务
func NewCaGroupClient() *SvcClient {
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		logrus.Printf("unable to connect to notificator: %s", err.Error())
		return new(SvcClient)
	}

	logrus.Printf("connected to test service")
	return &SvcClient{
		test: pb.NewBProjectClient(conn),
	}
}

func main() {
	flag.Parse()
	ceshiClient = NewCaGroupClient()
	//shanghai, _ = time.LoadLocation("Asia/Shanghai")
	if *run == "" {
		fmt.Println("使用-run输出需要执行的函数名")
		listFunc()
	} else {
		execFunc(*run)
		//fmt.Println("接口", *run, "不存在")
	}
}

func execFunc(name string) {
	fun := reflect.ValueOf(ceshiClient).MethodByName(name)
	fun.Call(nil)
}

func listFunc() {
	fmt.Println("Function List:")
	v := reflect.ValueOf(ceshiClient)
	v.NumMethod()
	t := v.Type()
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("  ", i, ">", m.Name)
	}
}

func (s SvcClient) TestAdd() {
	reply, err := s.test.AddUser(context.Background(), &pb.AddUserRequest{
		Username: "baikai",
		Fraction: "100",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	bt, _ := json.MarshalIndent(reply, "  ", "  ")
	fmt.Println(string(bt))
}
func (s SvcClient) GetTest() {
	reply, err := s.test.GetUser(ctx, &pb.GetUserRequest{
		UserID: 1402112376626483200,
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	bt, _ := json.MarshalIndent(reply, "  ", "  ")
	fmt.Println(string(bt))
}
