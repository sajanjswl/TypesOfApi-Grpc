package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sajanjswl/server_streaming/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hello i am client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client %f ", c)

	doGreetManyTimes(c)
}

func doGreetManyTimes(c greetpb.GreetServiceClient) {

	fmt.Println("starting to do a server Streamin  RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "idiota",
			LastName:  "Jaiswal",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet Rpc %v", err)
	}
	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("Response from greetManyTimes : %v", msg.Result)
	}

}
