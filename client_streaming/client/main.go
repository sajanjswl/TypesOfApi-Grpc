package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sajanjswl/client_streaming/greetpb"
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

	stream, err := c.GreetManyTimes(context.Background())

	if err != nil {
		log.Fatalf("error while calling greet Rpc %v", err)
	}
	fmt.Println("starting client streaming")

	for i := 0; i < 5; i++ {

		err := stream.Send(req)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Println(" %v", err)
	}

	fmt.Println("response from server", res)
}
