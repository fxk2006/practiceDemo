package main

import (
	"fmt"
	"net"
	pb "practiceDemo/grcp_gateway_work/simple"
	"practiceDemo/grcp_gateway_work/sserver"

	"google.golang.org/grpc"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	addr := "127.0.0.1:32111"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("server start failed ：", err)
		return
	}
	s := sserver.NewServerSimple()
	grpcS := grpc.NewServer()
	fmt.Println("server run ", addr)
	pb.RegisterSimpleServerServer(grpcS, s)
	grpcS.Serve(lis)
}
