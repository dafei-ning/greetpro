package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is a client.")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

}
