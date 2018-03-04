/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package main

import (
	"google.golang.org/grpc"
	"log"
	"os"
	pb "../protocol"
	"time"
	"golang.org/x/net/context"
)

const (
	address ="localhost:50051"
	defaultName ="world"
)
func main() {
	conn, err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}

	//main 方法退出的时候执行close
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting : %s", r.Message)
}