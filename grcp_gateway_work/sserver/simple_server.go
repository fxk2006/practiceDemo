package sserver

import (
	"context"
	"fmt"
	pb "practiceDemo/grcp_gateway_work/simple"
)

type ServerSimple struct {
}

func NewServerSimple() *ServerSimple {
	ss := new(ServerSimple)
	return ss
}

func (s *ServerSimple) Start() {
}

// 接口继承 protobuf的grpc 中的 SimpleServerServer 接口。 请求数据入口是在这里
func (s *ServerSimple) Stream(ctx context.Context, r *pb.Package_Request) (w *pb.Package_Response, err error) {
	fmt.Println("请求命令：", r.Cmd)
	w = new(pb.Package_Response)
	w.Code = 200
	w.Message = "ok"
	return w, nil
}
