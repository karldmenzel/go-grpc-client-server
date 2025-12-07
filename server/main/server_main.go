package main

import (
	"context"
	"fmt"
	"go-grpc-client-server/server/math"
	pb "go-grpc-client-server/shared"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMathServerServer
}

var (
	magicAddCounter      = 0
	magicSubtractCounter = 0
	magicFindMinCounter  = 0
	magicFindMaxCounter  = 0
)

func (s *server) MagicAdd(_ context.Context, in *pb.DoubleTerms) (*pb.DoubleResult, error) {
	fmt.Println("MagicAdd")
	magicAddCounter++

	sum := math.MagicAdd(in.TermOne, in.TermTwo)
	return &pb.DoubleResult{Result: sum}, nil
}

func (s *server) MagicSubtract(_ context.Context, in *pb.DoubleTerms) (*pb.DoubleResult, error) {
	fmt.Println("MagicSubtract")
	magicSubtractCounter++

	difference := math.MagicSubtract(in.TermOne, in.TermTwo)
	return &pb.DoubleResult{Result: difference}, nil
}

func (s *server) MagicFindMin(_ context.Context, in *pb.IntTerms) (*pb.IntResult, error) {
	fmt.Println("MagicFindMin")
	magicFindMinCounter++

	minimum := math.MagicFindMin(in.TermOne, in.TermTwo, in.TermThree)
	return &pb.IntResult{Result: minimum}, nil
}

func (s *server) MagicFindMax(_ context.Context, in *pb.IntTerms) (*pb.IntResult, error) {
	fmt.Println("MagicFindMax")
	magicFindMaxCounter++

	minimum := math.MagicFindMax(in.TermOne, in.TermTwo, in.TermThree)
	return &pb.IntResult{Result: minimum}, nil
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
