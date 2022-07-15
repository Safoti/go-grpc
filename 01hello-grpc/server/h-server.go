package main

import (
	"context"
	"fmt"
	"go-grpc-demo/01hello-grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/14
 * @Description  grpc 服务端编写
 **/

// 定义结构体，在调用注册api的时候作为入参，
// 该结构体会带上SayHello方法，里面是业务代码
// 这样远程调用时就执行了业务代码了
type server struct {
	// pb.go中自动生成的，是个空结构体
	pb.UnimplementedGreeterServer
}
//服务端调用方法 改写 ,原方法如下，现在进行改写
//func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {

func ( gt *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 打印请求参数
	fmt.Println("获取到的数据为：",req.GetName())
	// 实例化结构体HelloResponse，作为返回值
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err :=net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化gRPC server结构体
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	log.Println("开始监听，等待远程调用...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


