package main

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"time"

	pb "github.com/karldmenzel/go-grpc-client-server/magicMath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var waitGroup sync.WaitGroup

func main() {
	// Set up a connection to the server.
	conn, server := connectToServer("localhost:50051")
	// This is run after the end of the main function, and forcefully terminates the HTTP connection.
	defer conn.Close()

	// Create the context / metadata object passed into each gRPC request.
	// This context object just tells gRPC to cancel the call if the server hasn't responded in five seconds.
	requestContext, cancel := createRequestContext()
	// This is run after the end of the main function, it cancels any dangling requests.
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
func connectToServer(serverAddress string) (*grpc.ClientConn, pb.MagicMathClient) {

	// Create a new gRPC connection with no authentication.
	connection, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("failed to connect to server: %v", err))
	}

	// This is the object which actually has the remote functions on it.
	server := pb.NewMagicMathClient(connection)

	return connection, server
}

// This function creates a context object which is passed in to all RPC requests.
// For us, that context object just says that the request should time out after five seconds.
func createRequestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// This function makes 1000 requests to the server for a random function.
// Each call is made in a go routine, which all run concurrently.
// These go routines are grouped in a 'wait group', which allows us to wait for them all to finish.
func make1000Requests(server pb.MagicMathClient, requestContext context.Context) {
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

// This function makes the gRPC addition call to the server using two random doubles.
func magicAdd(server pb.MagicMathClient, requestContext context.Context) {
	_, err := server.MagicAdd(requestContext, &pb.DoubleTerms{TermOne: randomDouble(), TermTwo: randomDouble()})
	if err != nil {
		panic("Error on MagicAdd.")
	}
}

// This function makes the gRPC subtraction call to the server using two random doubles.
func magicSubtract(server pb.MagicMathClient, requestContext context.Context) {
	_, err := server.MagicSubtract(requestContext, &pb.DoubleTerms{TermOne: randomDouble(), TermTwo: randomDouble()})
	if err != nil {
		panic("Error on MagicSubtract.")
	}
}

// This function makes the gRPC min call to the server using three random integers.
func magicFindMin(server pb.MagicMathClient, requestContext context.Context) {
	_, err := server.MagicFindMin(requestContext, &pb.IntTerms{TermOne: randomInt(), TermTwo: randomInt(), TermThree: randomInt()})
	if err != nil {
		panic("Error on MagicFindMin.")
	}
}

// This function makes the gRPC max call to the server using three random integers.
func magicFindMax(server pb.MagicMathClient, requestContext context.Context) {
	_, err := server.MagicFindMax(requestContext, &pb.IntTerms{TermOne: randomInt(), TermTwo: randomInt(), TermThree: randomInt()})
	if err != nil {
		panic("Error on MagicFindMin.")
	}
}

// This function generates a random double between 0 and half the maximum float value.
func randomDouble() float64 {
	// Here we divide by 2 so that if two large numbers are added they don't overflow
	return rand.Float64() * (math.MaxFloat64 / 2)
}

// This function generates a random integer between zero and 64 bit int max.
func randomInt() int64 {
	return rand.Int64() * math.MaxInt64
}

// This function calls four gRPC methods to get the counter for each method, print them all, and then print the total.
func getCounters(server pb.MagicMathClient, requestContext context.Context) {
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
