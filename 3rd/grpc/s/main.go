package main

import (
	pb "github.com/figoxu/goPraticse/3rd/grpc/m"
	"context"
	"fmt"
	"net"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	"log"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	chk(err)
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, HelloService)
	grpclog.Info("Listen on" + Address)
	log.Println("Listen on" + Address)
	s.Serve(listen)
}

type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	log.Println("被调用：",resp.Message)
	return resp, nil
}

func chk(err error) {
	if err != nil {
		grpclog.Fatalln(err)
	}
}
