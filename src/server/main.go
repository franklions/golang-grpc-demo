package main

import (
	"fmt"
	"context"
	pb "../protocol"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port=":50051"
)

type server struct{}

func (s *server ) SayHello(ctx context.Context,in *pb.HelloRequest) (*pb.HelloReply,error){
	fmt.Println("Client Request Message:" + in.Name)
	return &pb.HelloReply{Message:"Hello "+in.Name},nil
}

func main(){

	lis,err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)
	fmt.Println("run server....."  )
	fmt.Println(lis.Addr())
	if err :=s.Serve(lis);err != nil{
		log.Fatalf("failed to serve: %v", err)
	}

}