package main

import (
	"context"
	"fmt"
	"log"

	"github.com/awsome_projects/greeting_greets/greetpb"
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
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("starting to do Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Veer",
			LastName:  "Jaiswal",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet Rpc %v", err)
	}
	log.Printf("Response from greet : %v", res.Result)
}
