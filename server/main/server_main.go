package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"go-grpc-client-server/server/math"
	pb "go-grpc-client-server/shared"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedMathServerServer
}

func (s *server) MagicAdd(_ context.Context, in *pb.DoubleTerms) (*pb.DoubleResult, error) {
	sum := math.MagicAdd(in.TermOne, in.TermTwo)
	return &pb.DoubleResult{Result: sum}, nil
}

func (s *server) MagicSubtract(_ context.Context, in *pb.DoubleTerms) (*pb.DoubleResult, error) {
	difference := math.MagicSubtract(in.TermOne, in.TermTwo)
	return &pb.DoubleResult{Result: difference}, nil
}

func (s *server) MagicFindMin(_ context.Context, in *pb.IntTerms) (*pb.IntResult, error) {
	minimum := math.MagicFindMin(in.TermOne, in.TermTwo, in.TermThree)
	return &pb.IntResult{Result: minimum}, nil
}

func (s *server) MagicFindMax(_ context.Context, in *pb.IntTerms) (*pb.IntResult, error) {
	minimum := math.MagicFindMax(in.TermOne, in.TermTwo, in.TermThree)
	return &pb.IntResult{Result: minimum}, nil
}

func main() {
	fmt.Println("Hello world from the server!")

	result := math.MagicAdd(2, 2)
	fmt.Println("Two plus two is equal to: ", result)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMathServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
