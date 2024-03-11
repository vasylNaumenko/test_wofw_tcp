package main

import (
	"context"
	"log"

	"github.com/bwesterb/go-pow"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"test_wofw_client/proto/cite"
)

const (
	port       = "9090"
	powMessage = "pow"
	serverAddr = "server"
)

func main() {
	conn, err := grpc.Dial(serverAddr+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := cite.NewCiteServiceClient(conn)
	stream, err := client.GetCite(context.Background())
	if err != nil {
		log.Fatalf("Error on get cite: %v", err)
	}

	// get pow
	if err = stream.Send(&cite.POWRequest{}); err != nil {
		log.Fatalf("Error on getting pow riddle: %v", err)
	}
	// get pow riddle
	res, err := stream.Recv()
	if err != nil {
		log.Fatalf("Error when receiving server response: %v", err)
	}
	// calculate pow
	resp, _ := pow.Fulfil(res.GetPowRiddle(), []byte(powMessage))

	// get cite
	if err = stream.Send(&cite.POWRequest{Pow: resp}); err != nil {
		log.Fatalf("Error on send pow request: %v", err)
	}
	// send pow
	res, err = stream.Recv()
	if err != nil {
		log.Fatalf("Error when receiving server response for cite: %v", err)
	}

	log.Printf("Cite from server: %s", res.GetCite())
}
