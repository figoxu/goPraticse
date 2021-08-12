package main

import (
	"context"
	"log"
	"time"

	"figoxu.me/sample/proto/hello"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	srcgrpc "google.golang.org/grpc"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	r := registry.New(cli)

	connGrpc, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///helloworld"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connGrpc.Close()

	for {
		callGRPC(r, connGrpc)
		time.Sleep(time.Second)
	}
}

func callGRPC(r *registry.Registry, conn *srcgrpc.ClientConn) {
	client := hello.NewFooClient(conn)
	reply, err := client.Bar(context.Background(), &hello.HelloReq{Content: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}
