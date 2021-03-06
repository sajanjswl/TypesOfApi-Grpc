package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/sajanjswl/server_streaming/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {

		result := "hello " + firstName + "" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
