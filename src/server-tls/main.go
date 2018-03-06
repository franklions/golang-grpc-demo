package main

import (
	"golang.org/x/net/context"
	pb "../protocol"
	"net"
	"log"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"flag"
)

const (
	address="127.0.0.1:50052"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply,error){
	resp := new(pb.HelloReply)
	resp.Message = "hello "+req.Name
	return resp,nil
}

func main() {
	certFilePath := flag.String("certfile","keys/server1.pem","server pem cert")
	keyFilePath  := flag.String("keyfile","keys/server1.key","server private key")

	listen, err :=net.Listen("tcp",address)
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}

	creds ,err := credentials.NewServerTLSFromFile(*certFilePath,
		*keyFilePath)
	if err != nil {
		log.Fatalf("failed to generate credentials: %v",err)
	}

	// 实例化grpc Server, 并开启TLS认证
	s :=grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)

	if err :=s.Serve(listen);err !=nil {
		log.Fatalf("failed to server %v",err)
	}
}
