package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/sajanjswl/client_streaming/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(stream greetpb.GreetService_GreetManyTimesServer) error {

	for {

		request, err := stream.Recv()

		if err != nil {

			if err == io.EOF {
				stream.SendAndClose(&greetpb.GreetManyTimesResponse{Result: "ending client stresaming"})
				return nil
			}
			return err

		}
		fmt.Println(request)
	}

}

func main() {

	fmt.Println("hello idotoa")

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
