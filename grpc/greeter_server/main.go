package main

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

import (
	"log"
	"net"

	pb "github.com/drichardson/go-examples/grpc/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conn, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}
	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{})
	reflection.Register(srv)
	if err := srv.Serve(conn); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
