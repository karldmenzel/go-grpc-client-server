package main

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"time"

	pb "go-grpc-client-server/shared"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var waitGroup sync.WaitGroup

func main() {
	// Set up a connection to the server.
	conn, server := connectToServer("localhost:50051")
	defer conn.Close()

	// Create the context object passed into each request.
	requestContext, cancel := createRequestContext()
	defer cancel()

	// Make all 1000 requests concurrently.
	make1000Requests(server, requestContext)

	// Wait for all 1000 requests to finish.
	waitGroup.Wait()

	// Print the stats on how many times each function was called.
	getCounters(server, requestContext)
}

// This function takes the URL of the server, and returns the raw connection object, and our server's object.
// The connection object is returned strictly so that it can be closed by the main function.
// The server object is what actually has the remote procedures exposed on it, and is what we will call.
func connectToServer(serverAddress string) (*grpc.ClientConn, pb.MathServerClient) {

	connection, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("failed to connect to server: %v", err))
	}

	server := pb.NewMathServerClient(connection)

	return connection, server
}

// This function creates a context object which is passed in to all RPC requests.
// For us, that context object just says that the request should time out after one second.
func createRequestContext() (context.Context, context.CancelFunc) {
	// All requests should time out after one second.
	return context.WithTimeout(context.Background(), 1*time.Second)
}

// This function makes 1000 requests to the server for a random function.
// Each call is made in a go routine, which all run concurrently.
// These go routines are grouped in a 'wait group', which allows us to wait for them all to finish.
func make1000Requests(server pb.MathServerClient, requestContext context.Context) {
	for range 1000 {
		methodId := rand.IntN(4)
		switch methodId {
		case 0:
			waitGroup.Go(func() { magicAdd(server, requestContext) })
		case 1:
			waitGroup.Go(func() { magicSubtract(server, requestContext) })
		case 2:
			waitGroup.Go(func() { magicFindMin(server, requestContext) })
		case 3:
			waitGroup.Go(func() { magicFindMax(server, requestContext) })
		default:
			panic("Random generation went out of range 0 - 3.")
		}
	}
}

func magicAdd(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicAdd(requestContext, &pb.DoubleTerms{TermOne: randomDouble(), TermTwo: randomDouble()})
	if err != nil {
		panic("Error on MagicAdd.")
	}
}

func magicSubtract(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicSubtract(requestContext, &pb.DoubleTerms{TermOne: randomDouble(), TermTwo: randomDouble()})
	if err != nil {
		panic("Error on MagicSubtract.")
	}
}

func magicFindMin(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicFindMin(requestContext, &pb.IntTerms{TermOne: randomInt(), TermTwo: randomInt(), TermThree: randomInt()})
	if err != nil {
		panic("Error on MagicFindMin.")
	}
}

func magicFindMax(server pb.MathServerClient, requestContext context.Context) {
	_, err := server.MagicFindMax(requestContext, &pb.IntTerms{TermOne: randomInt(), TermTwo: randomInt(), TermThree: randomInt()})
	if err != nil {
		panic("Error on MagicFindMin.")
	}
}

func randomDouble() float64 {
	// Here we divide by 2 so that if two large numbers are added they don't overflow
	return rand.Float64() * (math.MaxFloat64 / 2)
}

func randomInt() int64 {
	return rand.Int64() * math.MaxInt64
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
