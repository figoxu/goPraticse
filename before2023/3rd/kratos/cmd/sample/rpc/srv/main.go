package main

import (
	"context"
	"fmt"
	"log"

	"figoxu.me/sample/proto/hello"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	etcd "go.etcd.io/etcd/client/v3"
)

type server struct {
	hello.UnimplementedFooServer
}

func (s *server) Bar(ctx context.Context, in *hello.HelloReq) (*hello.WorldRsp, error) {
	return &hello.WorldRsp{Content: fmt.Sprintf("Welcome %+v!", in.Content)}, nil
}

func main() {
	client, err := etcd.New(etcd.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}

	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)

	s := &server{}
	hello.RegisterFooServer(grpcSrv, s)

	r := registry.New(client)
	app := kratos.New(
		kratos.Name("helloworld"),
		kratos.Server(
			grpcSrv,
		),
		kratos.Registrar(r),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
