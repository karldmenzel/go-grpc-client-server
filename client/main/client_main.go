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
	conn, server := connectToServer("localhost:50051")
	defer conn.Close()

	requestContext, cancel := createRequestContext()
	defer cancel()

	make1000Requests(server, requestContext)

	getCounters(server, requestContext)
}

func connectToServer(serverAddress string) (*grpc.ClientConn, pb.MathServerClient) {

	connection, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("failed to connect to server: %v", err))
	}

	server := pb.NewMathServerClient(connection)

	return connection, server
}

func createRequestContext() (context.Context, context.CancelFunc) {
	// All requests should time out after one second.
	return context.WithTimeout(context.Background(), time.Second)

}

func make1000Requests(server pb.MathServerClient, requestContext context.Context) {
	for range 1000 {
		methodId := rand.IntN(4)
		switch methodId {
		case 0:
			magicAdd(server, requestContext)
		case 1:
			magicSubtract(server, requestContext)
		case 2:
			magicFindMin(server, requestContext)
		case 3:
			magicFindMax(server, requestContext)
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

func getCounters(server pb.MathServerClient, requestContext context.Context) {
	addCount, err := server.GetAddCount(requestContext, &pb.Empty{})
	subCount, err := server.GetSubCount(requestContext, &pb.Empty{})
	minCount, err := server.GetMinCount(requestContext, &pb.Empty{})
	maxCount, err := server.GetMaxCount(requestContext, &pb.Empty{})
	if err != nil {
		fmt.Printf("Error getting counts: %v\n", err)
	}

	fmt.Printf("Add count: %d\n", addCount.Count)
	fmt.Printf("Sub count: %d\n", subCount.Count)
	fmt.Printf("Min count: %d\n", minCount.Count)
	fmt.Printf("Max count: %d\n", maxCount.Count)
	fmt.Printf("Total request count: %d\n", addCount.Count+subCount.Count+minCount.Count+maxCount.Count)
}
