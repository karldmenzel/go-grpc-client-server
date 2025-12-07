package main

import (
	"context"
	"fmt"
	"go-grpc-client-server/server/math"
	pb "go-grpc-client-server/shared"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMathServerServer
}

var (
	addFuncCounter int64 = 0
	subFuncCounter int64 = 0
	minFuncCounter int64 = 0
	maxFuncCounter int64 = 0
)

var (
	addCounterMutex sync.Mutex
	subCounterMutex sync.Mutex
	minCounterMutex sync.Mutex
	maxCounterMutex sync.Mutex
)

func (s *server) MagicAdd(_ context.Context, in *pb.DoubleTerms) (*pb.DoubleResult, error) {
	addCounterMutex.Lock()
	addFuncCounter++
	addCounterMutex.Unlock()

	sum := math.MagicAdd(in.TermOne, in.TermTwo)
	return &pb.DoubleResult{Result: sum}, nil
}

func (s *server) MagicSubtract(_ context.Context, in *pb.DoubleTerms) (*pb.DoubleResult, error) {
	subCounterMutex.Lock()
	subFuncCounter++
	subCounterMutex.Unlock()

	difference := math.MagicSubtract(in.TermOne, in.TermTwo)
	return &pb.DoubleResult{Result: difference}, nil
}

func (s *server) MagicFindMin(_ context.Context, in *pb.IntTerms) (*pb.IntResult, error) {
	minCounterMutex.Lock()
	minFuncCounter++
	minCounterMutex.Unlock()

	minimum := math.MagicFindMin(in.TermOne, in.TermTwo, in.TermThree)
	return &pb.IntResult{Result: minimum}, nil
}

func (s *server) MagicFindMax(_ context.Context, in *pb.IntTerms) (*pb.IntResult, error) {
	maxCounterMutex.Lock()
	maxFuncCounter++
	maxCounterMutex.Unlock()

	maximum := math.MagicFindMax(in.TermOne, in.TermTwo, in.TermThree)
	return &pb.IntResult{Result: maximum}, nil
}

func (s *server) GetAddCount(_ context.Context, _ *pb.Empty) (*pb.Count, error) {
	addCounterMutex.Lock()
	count := addFuncCounter
	addCounterMutex.Unlock()

	return &pb.Count{Count: count}, nil
}

func (s *server) GetSubCount(_ context.Context, _ *pb.Empty) (*pb.Count, error) {
	subCounterMutex.Lock()
	count := subFuncCounter
	subCounterMutex.Unlock()

	return &pb.Count{Count: count}, nil
}

func (s *server) GetMinCount(_ context.Context, _ *pb.Empty) (*pb.Count, error) {
	minCounterMutex.Lock()
	count := minFuncCounter
	minCounterMutex.Unlock()

	return &pb.Count{Count: count}, nil
}

func (s *server) GetMaxCount(_ context.Context, _ *pb.Empty) (*pb.Count, error) {
	maxCounterMutex.Lock()
	count := maxFuncCounter
	maxCounterMutex.Unlock()

	return &pb.Count{Count: count}, nil
}

func main() {
	fmt.Println("The magic math server is running!")

	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("failed to listen to port %d: %v", port, err)
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterMathServerServer(s, &server{})

	fmt.Printf("The server is listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		panic(err)
	}
}
