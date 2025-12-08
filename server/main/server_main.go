package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"go-grpc-client-server/server/math"
	magicInterface "go-grpc-client-server/shared"

	"google.golang.org/grpc"
)

// This is the server object that we will bind to in order to expose the remote methods.
type server struct {
	magicInterface.UnimplementedMathServerServer
}

// These counters store how many times each function has been called.
var (
	addFuncCounter int64 = 0
	subFuncCounter int64 = 0
	minFuncCounter int64 = 0
	maxFuncCounter int64 = 0
)

// These mutexes are used to protect read and write access to the function counters above.
var (
	addCounterMutex sync.Mutex
	subCounterMutex sync.Mutex
	minCounterMutex sync.Mutex
	maxCounterMutex sync.Mutex
)

func main() {
	fmt.Println("The magic math server is running!")

	// Listen to incoming TCP connections on port 50051 (common gRPC development port).
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("failed to listen to port %d: %v", port, err)
		panic(err)
	}
	// Create a new unbound gRPC server.
	s := grpc.NewServer()
	// Bind the magic interface to the gRPC server.
	magicInterface.RegisterMathServerServer(s, &server{})

	fmt.Printf("The server is listening at %v\n", lis.Addr())

	// Begin serving requests at the listening port, s.Server will run in a loop until the server terminates.
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		panic(err)
	}
}

// ========================================== Math Functions ==========================================

// MagicAdd takes a request context (which is ignored) and two doubles, and returns their sum.
func (s *server) MagicAdd(_ context.Context, in *magicInterface.DoubleTerms) (*magicInterface.DoubleResult, error) {
	addCounterMutex.Lock()
	addFuncCounter++
	addCounterMutex.Unlock()

	sum := math.LocalAdd(in.TermOne, in.TermTwo)
	responseObject := &magicInterface.DoubleResult{Result: sum}

	return responseObject, nil
}

// MagicSubtract takes a request context (which is ignored) and two doubles, and returns their difference.
func (s *server) MagicSubtract(_ context.Context, in *magicInterface.DoubleTerms) (*magicInterface.DoubleResult, error) {
	subCounterMutex.Lock()
	subFuncCounter++
	subCounterMutex.Unlock()

	difference := math.LocalSubtract(in.TermOne, in.TermTwo)
	responseObject := &magicInterface.DoubleResult{Result: difference}

	return responseObject, nil
}

// MagicFindMin takes a request context (which is ignored) and three integers, and returns the lowest value.
// If all three values are equal it returns the first value.
func (s *server) MagicFindMin(_ context.Context, in *magicInterface.IntTerms) (*magicInterface.IntResult, error) {
	minCounterMutex.Lock()
	minFuncCounter++
	minCounterMutex.Unlock()

	minimum := math.LocalFindMin(in.TermOne, in.TermTwo, in.TermThree)
	responseObject := &magicInterface.IntResult{Result: minimum}

	return responseObject, nil
}

// MagicFindMax takes a request context (which is ignored) and three integers, and returns the highest value.
// If all three values are equal it returns the first value.
func (s *server) MagicFindMax(_ context.Context, in *magicInterface.IntTerms) (*magicInterface.IntResult, error) {
	maxCounterMutex.Lock()
	maxFuncCounter++
	maxCounterMutex.Unlock()

	maximum := math.LocalFindMax(in.TermOne, in.TermTwo, in.TermThree)
	responseObject := &magicInterface.IntResult{Result: maximum}

	return responseObject, nil
}

// ========================================== Counter Functions ==========================================

// GetAddCount returns the total number of times MagicAdd has been called.
func (s *server) GetAddCount(_ context.Context, _ *magicInterface.Empty) (*magicInterface.Count, error) {
	addCounterMutex.Lock()
	count := addFuncCounter
	addCounterMutex.Unlock()

	responseObject := &magicInterface.Count{Count: count}

	return responseObject, nil
}

// GetSubCount returns the total number of times MagicSubtract has been called.
func (s *server) GetSubCount(_ context.Context, _ *magicInterface.Empty) (*magicInterface.Count, error) {
	subCounterMutex.Lock()
	count := subFuncCounter
	subCounterMutex.Unlock()

	responseObject := &magicInterface.Count{Count: count}

	return responseObject, nil
}

// GetMinCount returns the total number of times MagicFindMin has been called.
func (s *server) GetMinCount(_ context.Context, _ *magicInterface.Empty) (*magicInterface.Count, error) {
	minCounterMutex.Lock()
	count := minFuncCounter
	minCounterMutex.Unlock()

	responseObject := &magicInterface.Count{Count: count}

	return responseObject, nil
}

// GetMaxCount returns the total number of times MagicFindMax has been called.
func (s *server) GetMaxCount(_ context.Context, _ *magicInterface.Empty) (*magicInterface.Count, error) {
	maxCounterMutex.Lock()
	count := maxFuncCounter
	maxCounterMutex.Unlock()

	responseObject := &magicInterface.Count{Count: count}

	return responseObject, nil
}
