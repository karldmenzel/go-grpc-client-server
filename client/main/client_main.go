package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	pb "go-grpc-client-server/shared"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	serverAddress := "localhost:50051"
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	server := pb.NewMathServerClient(conn)

	// All requests should time out after one second.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for range 1000 {
		methodId := rand.IntN(4)
		switch methodId {
		case 0:
			magicAdd(server, ctx)
		case 1:
			magicSubtract(server, ctx)
		case 2:
			magicFindMin(server, ctx)
		case 3:
			magicFindMax(server, ctx)
		default:
			panic("Random generation went out of range 0 - 3.")
		}
	}
}

func magicAdd(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicAdd(requestContext, &pb.DoubleTerms{TermOne: 5.0, TermTwo: 2.0})
	if err != nil {
		panic("Error on MagicAdd.")
	}
}

func magicSubtract(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicSubtract(requestContext, &pb.DoubleTerms{TermOne: 10.0, TermTwo: 5.0})
	if err != nil {
		panic("Error on MagicSubtract.")
	}
}

func magicFindMin(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicFindMin(requestContext, &pb.IntTerms{TermOne: 1.0, TermTwo: 2.0, TermThree: 5.0})
	if err != nil {
		panic("Error on MagicFindMin.")
	}
}

func magicFindMax(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicFindMax(requestContext, &pb.IntTerms{TermOne: 1.0, TermTwo: 2.0, TermThree: 5.0})
	if err != nil {
		panic("Error on MagicFindMin.")
	}
}
