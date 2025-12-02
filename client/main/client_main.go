package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "go-grpc-client-server/shared"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMathServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sum, err := c.MagicAdd(ctx, &pb.DoubleTerms{TermOne: 5.0, TermTwo: 2.0})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Got a sum of: %v", sum.GetResult())

	difference, err := c.MagicSubtract(ctx, &pb.DoubleTerms{TermOne: 5.0, TermTwo: 2.0})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Got a difference of: %v", difference.GetResult())

	minimum, err := c.MagicFindMin(ctx, &pb.IntTerms{TermOne: 5.0, TermTwo: 2.0, TermThree: 3.0})
	if err != nil {
		log.Fatalf("could not find min: %v", err)
	}
	log.Printf("Got a min of: %v", minimum.GetResult())

	maximum, err := c.MagicFindMax(ctx, &pb.IntTerms{TermOne: 5.0, TermTwo: 2.0, TermThree: 3.0})
	if err != nil {
		log.Fatalf("could not find max: %v", err)
	}
	log.Printf("Got a max of: %v", maximum.GetResult())

}
