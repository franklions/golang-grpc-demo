package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"google.golang.org/grpc"
	pb "../protocol"
	"time"
	"golang.org/x/net/context"
	"flag"
)

const (
	address = "127.0.0.1:50052"
)

func main() {
	certFilePath := flag.String("certfile","keys/server1.pem","server public cert")
	creds, err :=credentials.NewClientTLSFromFile(*certFilePath,
		"server name")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(address,grpc.WithTransportCredentials(creds))
	if err !=nil {
		log.Fatalf("failed to connect %v",err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "jerry"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting : %s", r.Message)
}
