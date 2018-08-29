package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "github.com/figoxu/goPraticse/3rd/grpc/m"
	"context"
	"log"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	chk(err)
	defer conn.Close()
	client := pb.NewHelloClient(conn)
	requestParam := &pb.HelloRequest{Name: "gRPC"}
	res, err := client.SayHello(context.Background(), requestParam)
	chk(err)
	log.Println(res.Message)
	grpclog.Info(res.Message)
}

func chk(err error) {
	if err != nil {
		grpclog.Fatalln(err)
	}
}
