package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sajanjswl/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	fmt.Println("hello dsfk")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v ", err)
	}

}
