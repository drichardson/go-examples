package main

import (
	"log"
	"os"
	"time"

	pb "github.com/drichardson/go-examples/grpc/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	name := "WORLD"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln("SayHello request failed:", err)
	}

	log.Println("Greeting:", r.Message)
}
