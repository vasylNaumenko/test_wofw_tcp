package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/bwesterb/go-pow"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"test_wofw_client/proto/cite"
)

var (
	addr         = flag.String("addr", "server", "The address to dial")
	port         = flag.String("port", "9090", "The port to connect to")
	emptyRequest = flag.Int("empty-request", 1, "Number of times to send an empty request")
	powSalt      = flag.String("pow-salt", "pow", "The salt to use for the proof of work.")
	delay        = flag.Duration("delay", time.Millisecond, "Delay between requests. ex: 5s, Default is 1ms.")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr+":"+*port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := cite.NewCiteServiceClient(conn)
	stream, err := client.GetCite(context.Background())
	if err != nil {
		log.Fatalf("Error on get cite: %v", err)
	}

	var res *cite.CiteResponse
	for i := 0; i < *emptyRequest; i++ {
		// ask for pow riddle
		if err = stream.Send(&cite.POWRequest{}); err != nil {
			log.Fatalf("Error on sending empty request %d: %v", i, err)
		}

		// get pow riddle
		res, err = stream.Recv()
		if err != nil {
			log.Fatalf("Error when receiving server response: %v", err)
		}

		log.Printf("pow: %s \n", res.GetPowRiddle())
	}

	<-time.After(*delay)

	// calculate pow
	resp, _ := pow.Fulfil(res.GetPowRiddle(), []byte(*powSalt))

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
