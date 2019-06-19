package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/hyperledger/awsome_projects/greeting/greet_greetpb/greetpb"
)

func main() {

	fmt.Println("hello i am client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client %f ", c)
}
